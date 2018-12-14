#!/bin/sh
mockgen -package mock_services -source=proto/wx.pb.go > mock_services/wx_mock.go