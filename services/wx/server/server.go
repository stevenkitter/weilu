package server

import (
	"context"
	"log"
	"net"

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
		"app_id = ? and info_type = component_verify_ticket",
		req.AppID).
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
