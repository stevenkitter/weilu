#!/bin/sh
protoc -Iproto proto/*.proto --go_out=plugins=grpc:${GOPATH}/src


# mock
mockgen -package mock_services -source=proto/wx.pb.go > mock_services/wx_mock.go