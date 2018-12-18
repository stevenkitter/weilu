package client

import (
	"context"
	"log"
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

//DecryptMsg decrypt msg
func (c *Client) DecryptMsg(req *pb.WXEncryptedMessage) (*pb.Resp, error) {
	conn, cl, err := dialGrpc()
	defer func() {
		if err := conn.Close(); err != nil {
			log.Printf("conn.Close err : %v", err)
		}
	}()
	if err != nil {
		return nil, err
	}
	return cl.DecryptMsg(context.Background(), req)
}

//TicketReceived ticket received request
func (c *Client) TicketReceived(req *pb.WXTicketReq) (*pb.Resp, error) {
	conn, cl, err := dialGrpc()
	defer func() {
		if err := conn.Close(); err != nil {
			log.Printf("conn.Close err : %v", err)
		}
	}()
	if err != nil {
		return nil, err
	}
	return cl.TicketReceived(context.Background(), req)
}

func (c *Client) Ticket(req *pb.GetTicketReq) (*pb.Resp, error) {
	conn, cl, err := dialGrpc()
	defer func() {
		if err := conn.Close(); err != nil {
			log.Printf("conn.Close err : %v", err)
		}
	}()
	if err != nil {
		return nil, err
	}
	return cl.Ticket(context.Background(), req)
}

func dialGrpc() (*grpc.ClientConn, pb.WXServiceClient, error) {
	conn, err := grpc.Dial(WXAddress, grpc.WithInsecure())
	if err != nil {
		return nil, nil, err
	}
	cl := pb.NewWXServiceClient(conn)
	return conn, cl, nil
}
