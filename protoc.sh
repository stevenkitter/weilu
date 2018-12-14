#!/bin/sh
`protoc -Iproto proto/*.proto --go_out=plugins=grpc:$GOPATH/src`