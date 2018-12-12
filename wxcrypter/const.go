package wxcrypter

import "os"

const (
	//Token token
	Token = "K81ec6fF37"
	//AppID app id
	AppID = "wxdd9779d0ca45ea77"
	//EncodingAESKey encoding aes key
	EncodingAESKey = "7WfXuJfsGHYqt5eSPH8Gg7B9Y115vU8dx4Z48rZbzH1"
)

var (
	//AppSecrect app secrect
	AppSecrect = os.Getenv("WXAppSecrect")
)
