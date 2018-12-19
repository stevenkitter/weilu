package server

import (
	"context"
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
	component := database.WXComponent{}
	err := s.DB.Where(database.WXComponent{
		AppID:    req.AppID,
		InfoType: req.InfoType,
	}).Assign(database.WXComponent{
		Component: req.Component,
	}).FirstOrCreate(&component).Error
	log.Printf("ticket received is %s", req.Component)
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
