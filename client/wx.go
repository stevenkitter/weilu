package client

import (
	"context"
	pb "github.com/stevenkitter/weilu/proto"
	"log"
)

//Client used for request server
type Client struct {
	Address string //哪个微服务的客户端
}

//DecryptMsg decrypt msg
func (c *Client) DecryptMsg(req *pb.WXEncryptedMessage) (*pb.Resp, error) {
	conn, cl, err := dialGrpc(c.Address)
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
	conn, cl, err := dialGrpc(c.Address)
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
	conn, cl, err := dialGrpc(c.Address)
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

func (c *Client) AccessToken(req *pb.GetAccessTokenReq) (*pb.Resp, error) {
	conn, cl, err := dialGrpc(c.Address)
	defer func() {
		if err := conn.Close(); err != nil {
			log.Printf("conn.Close err : %v", err)
		}
	}()
	if err != nil {
		return nil, err
	}
	return cl.AccessToken(context.Background(), req)
}

func (c *Client) PreAuthCode(req *pb.GetPreAuthCodeReq) (*pb.Resp, error) {
	conn, cl, err := dialGrpc(c.Address)
	defer func() {
		if err := conn.Close(); err != nil {
			log.Printf("conn.Close err : %v", err)
		}
	}()
	if err != nil {
		return nil, err
	}
	return cl.PreAuthCode(context.Background(), req)
}

func (c *Client) AuthURL(req *pb.GetAuthURLReq) (*pb.Resp, error) {
	conn, cl, err := dialGrpc(c.Address)
	defer func() {
		if err := conn.Close(); err != nil {
			log.Printf("conn.Close err : %v", err)
		}
	}()
	if err != nil {
		return nil, err
	}
	return cl.AuthURL(context.Background(), req)
}
