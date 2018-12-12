package wxcrypter

import (
	"encoding/xml"
)

//EncMessage wx coming msg
type EncMessage struct {
	XMLName      xml.Name `xml:"xml"`
	ToUserName   string   // 开发者微信号
	Encrypt      string   // 加密的消息报文
	MsgSignature string   // 报文签名
	TimeStamp    string   // 时间戳
	Nonce        string   // 随机数
}

//ReceivedMessage encrypt msg
type ReceivedMessage struct {
	XMLName                      xml.Name `xml:"xml"`
	AppID                        string   `xml:"AppId"`                 //第三方平台appid
	CreateTime                   int64    `xml:"CreateTime"`            //时间戳
	InfoType                     string   `xml:"InfoType"`              //component_verify_ticket
	ComponentVerifyTicket        string   `xml:"ComponentVerifyTicket"` //Ticket内容
	AuthorizerAppid              string
	AuthorizationCode            string
	AuthorizationCodeExpiredTime string
	PreAuthCode                  string
}

//EncryptedResponseXML tencent response
type EncryptedResponseXML struct {
	XMLName      xml.Name `xml:"xml"`
	TimeStamp    string
	Encrypt      string
	MsgSignature string
	Nonce        string
}

//EncryptedRequestXML request
type EncryptedRequestXML struct {
	XMLName    xml.Name `xml:"xml"`
	ToUserName string
	Encrypt    string
}

//ParseEncRequestXML parse wx request
func ParseEncRequestXML(data []byte) (EncMessage, error) {
	e := EncMessage{}
	err := xml.Unmarshal(data, &e)
	if err != nil {
		return e, err
	}
	return e, nil
}

//ParseRequestXML parse request
func ParseRequestXML(data []byte) (e EncryptedRequestXML, err error) {
	err = xml.Unmarshal(data, &e)
	if err != nil {
		err = ErrorParseXML
	}
	return
}

//ParseResponseXML parse response
func ParseResponseXML(data []byte) (e EncryptedResponseXML, err error) {
	err = xml.Unmarshal(data, &e)
	if err != nil {
		err = ErrorParseXML
	}
	return
}

//GenerateResponseXML generate response
func GenerateResponseXML(encrypt, signature, timestamp, nonce string) ([]byte, error) {
	e := EncryptedResponseXML{
		Nonce:        nonce,
		Encrypt:      encrypt,
		TimeStamp:    timestamp,
		MsgSignature: signature,
	}

	b, err := xml.Marshal(e)
	if err != nil {
		err = ErrorGenReturnXML
	}
	return b, err
}
