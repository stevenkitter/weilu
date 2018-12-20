package server

import (
	"context"
	"errors"
	"fmt"
	"github.com/stevenkitter/weilu/services/wx/httpClient"
	"log"
	"net"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/stevenkitter/weilu/database"

	"github.com/jinzhu/gorm"
	pb "github.com/stevenkitter/weilu/proto"
	"github.com/stevenkitter/weilu/wxcrypter"
)

//Server wx server
type Server struct {
	DB *gorm.DB
}

//Run run wx server
func (s *Server) Run(port string) error {
	tcp, err := net.Listen("tcp", port)
	if err != nil {
		log.Printf("net.Listen err : %v", err)
		return err
	}
	grpcServer := grpc.NewServer()
	pb.RegisterWXServiceServer(grpcServer, s)
	reflection.Register(grpcServer)
	if err := grpcServer.Serve(tcp); err != nil {
		log.Printf("grpcServer.Serve err : %v", err)
		return err
	}
	return nil
}

//DecryptMsg decrypt tencent incoming message
func (s *Server) DecryptMsg(ctx context.Context, req *pb.WXEncryptedMessage) (*pb.Resp, error) {
	e, err := wxcrypter.NewEncrypter(wxcrypter.Token, wxcrypter.EncodingAESKey, wxcrypter.AppID)
	if err != nil {
		log.Printf("NewEncrypter err : %v", err)
		return nil, err
	}
	b, err := e.Decrypt([]byte(req.Msg), req.MsgSignature, req.Timestamp, req.Nonce)
	if err != nil {
		log.Printf("e.Decrypt err : %v", err)
		return nil, err
	}
	resp := &pb.Resp{
		Code: 200,
		Msg:  "",
		Data: string(b),
	}
	return resp, nil
}

//TicketReceived tencent ticket handler
func (s *Server) TicketReceived(ctx context.Context, req *pb.WXTicketReq) (*pb.Resp, error) {
	err := s.SaveTicket(req.Component)
	if err != nil {
		log.Printf("TicketReceived Save sql err : %v", err)
		return nil, err
	}
	resp := &pb.Resp{
		Code: 200,
	}
	return resp, nil
}

//Ticket get the saved Ticket in mysql
func (s *Server) Ticket(ctx context.Context, req *pb.GetTicketReq) (*pb.Resp, error) {
	tick := database.WXComponent{}
	err := s.DB.Where(
		"app_id = ? AND info_type = ?",
		req.AppID, database.ComponentVerifyTicket).
		Find(&tick).Error
	if err != nil && !gorm.IsRecordNotFoundError(err) {
		log.Printf("find ticket from mysql err : %v", err)
		return nil, err
	}
	return &pb.Resp{
		Code: 200,
		Data: tick.Component,
	}, nil
}

func (s *Server) AccessToken(ctx context.Context, req *pb.GetAccessTokenReq) (*pb.Resp, error) {
	comp := database.WXComponent{}
	err := s.DB.Where(
		"app_id = ? AND info_type = ?",
		req.AppID, database.ComponentAccessToken).Find(&comp).Error
	if err != nil && !gorm.IsRecordNotFoundError(err) {
		log.Printf("s.DB.Where err : %v", err)
		return nil, err
	}
	if gorm.IsRecordNotFoundError(err) || time.Now().Unix() >= int64(comp.Expired) {
		//update or create
		hcl := httpClient.Client{}
		ticket, err := s.Ticket(ctx, &pb.GetTicketReq{
			AppID: wxcrypter.AppID,
		})
		if err != nil {
			log.Printf("s.Ticket err : %v", err)
			return nil, err
		}
		token, expires, err := hcl.RequestAccessToken(ticket.Data)
		if err != nil {
			log.Printf("hcl.AccessToken err : %v ticket is %s", err, ticket.Data)
			return nil, err
		}
		go func() {
			if err := s.SaveTokenExpires(token, expires); err != nil {
				log.Printf("SaveTokenExpires err : %v", err)
			}
		}()
		return &pb.Resp{
			Code: 200,
			Data: token,
		}, nil
	}
	return &pb.Resp{
		Code: 200,
		Data: comp.Component,
	}, nil
}

func (s *Server) PreAuthCode(ctx context.Context, req *pb.GetPreAuthCodeReq) (*pb.Resp, error) {
	comp := database.WXComponent{}
	err := s.DB.Where(
		"app_id = ? AND info_type = ?",
		req.AppID, database.PreAuthCode).Find(&comp).Error
	if err != nil && !gorm.IsRecordNotFoundError(err) {
		log.Printf("s.DB.Where err : %v", err)
		return nil, err
	}
	if gorm.IsRecordNotFoundError(err) || time.Now().Unix() >= int64(comp.Expired) {
		hcl := httpClient.Client{}
		res, err := s.AccessToken(ctx, &pb.GetAccessTokenReq{
			AppID: wxcrypter.AppID,
		})
		if err != nil {
			log.Printf("s.AccessToken err : %v", err)
			return nil, err
		}
		authCode, expires, err := hcl.RequestPreAuthCode(res.Data)
		if err != nil {
			log.Printf("hcl.RequestPreAuthCode err : %v", err)
			return nil, err
		}
		go func() {
			if err := s.SaveAuthCodeExpires(authCode, expires); err != nil {
				log.Printf("SaveAuthCodeExpires err : %v", err)
			}
		}()
		return &pb.Resp{
			Code: 200,
			Data: authCode,
		}, nil
	}
	return &pb.Resp{
		Code: 200,
		Data: comp.Component,
	}, nil
}

func (s *Server) AuthURL(ctx context.Context, req *pb.GetAuthURLReq) (*pb.Resp, error) {
	device := req.Device
	authType := req.AuthType
	if device == "" || (device != "phone" && device != "web") {
		log.Printf("query device is wrong device : %s", device)
		return nil, errors.New("query device is required and must be web or phone")
	}
	if authType == "" || (authType != "1" && authType != "2" && authType != "3") {
		log.Printf("query authType is wrong authType : %s", authType)
		return nil, errors.New("query authType is required and must be 1 2 3")
	}
	res, err := s.PreAuthCode(ctx, &pb.GetPreAuthCodeReq{
		AppID: wxcrypter.AppID,
	})
	if err != nil {
		log.Printf("s.PreAuthCode err : %v", err)
		return nil, err
	}
	authCode := res.Data
	var url = ""
	if device == "web" {
		url += fmt.Sprintf("https://mp.weixin.qq.com/cgi-bin/componentloginpage?"+
			"component_appid=%s&pre_auth_code=%s&redirect_uri=%s&auth_type=%s", wxcrypter.AppID, authCode, req.RedirectURL, authType)
	} else {
		url += fmt.Sprintf("https://mp.weixin.qq.com/safe/bindcomponent?action=bindcomponent&no_scan=1"+
			"&component_appid=%s&pre_auth_code=%s&redirect_uri=%s&auth_type=%s#wechat_redirect",
			wxcrypter.AppID, authCode, req.RedirectURL, authType)
	}
	return &pb.Resp{
		Code: 200,
		Data: url,
	}, nil
}
