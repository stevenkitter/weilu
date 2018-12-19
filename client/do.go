package client

import (
	pb "github.com/stevenkitter/weilu/proto"
	"google.golang.org/grpc"
)

//
func dialGrpc(address string) (*grpc.ClientConn, pb.WXServiceClient, error) {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		return nil, nil, err
	}
	cl := pb.NewWXServiceClient(conn)
	return conn, cl, nil
}
