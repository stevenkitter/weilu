package client

import (
	"context"
	"os"

	pb "github.com/stevenkitter/weilu/proto"

	"google.golang.org/grpc"
)

//Client used for request server
type Client struct {
}

var (
	//WXAddress wx server address
	WXAddress = os.Getenv("WX_SERVER_ADDRESS")
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
