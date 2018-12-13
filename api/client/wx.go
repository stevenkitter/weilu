package client

import (
	"context"

	pb "github.com/stevenkitter/weilu/proto"

	"google.golang.org/grpc"
)

//Client used for request server
type Client struct {
}

const (
	//WXAddress wx server address
	WXAddress = "localhost:51001"
)

//DecryptMsg decryp msg
func (c *Client) DecryptMsg(req *pb.WXEncryptedMessage) (*pb.Resp, error) {
	conn, err := grpc.Dial(WXAddress, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	cl := pb.NewWXServiceClient(conn)
	return cl.DecryptMsg(context.Background(), req)
}

//TicketReceived ticket received request
func (c *Client) TicketReceived(req *pb.WXTicketReq) (*pb.Resp, error) {
	conn, err := grpc.Dial(WXAddress, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	cl := pb.NewWXServiceClient(conn)
	return cl.TicketReceived(context.Background(), req)
}
