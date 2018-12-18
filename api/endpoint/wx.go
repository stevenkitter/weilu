package endpoint

import (
	"encoding/xml"
	"errors"
	pb "github.com/stevenkitter/weilu/proto"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/stevenkitter/weilu/client"
	"github.com/stevenkitter/weilu/wxcrypter"
)

//SecurityQuery url query
type SecurityQuery struct {
	Signature    string `form:"signature"`
	Timestamp    string `form:"timestamp"`
	Nonce        string `form:"nonce"`
	EncryptType  string `form:"encrypt_type"`
	MsgSignature string `form:"msg_signature"`
}

//WXReceiveEndpoint wx received endpoint
func WXReceiveEndpoint(c *gin.Context) (interface{}, error) {
	query := SecurityQuery{}
	err := c.ShouldBindQuery(&query)
	if err != nil {
		return "success", err
	}

	signature := wxcrypter.Signature(wxcrypter.Token, query.Timestamp, query.Nonce)
	if signature != query.Signature {
		return nil, errors.New("post data is not coming from tencent")
	}
	if query.EncryptType != "aes" {
		return nil, errors.New("encrypt_type is not aes")
	}

	postData := wxcrypter.EncryptedRequestXML{}
	err = c.ShouldBindXML(&postData)
	if err != nil {
		return "success", err
	}
	log.Printf("postData %v\n", postData)
	cl := client.Client{}
	res, err := cl.DecryptMsg(&pb.WXEncryptedMessage{
		Msg:          postData.Encrypt,
		MsgSignature: query.MsgSignature,
		Timestamp:    query.Timestamp,
		Nonce:        query.Nonce,
	})
	if err != nil {
		return "success", err
	}

	wxMsg := wxcrypter.ReceivedMessage{}
	err = xml.Unmarshal([]byte(res.Data), &wxMsg)
	if err != nil {
		return "success", err
	}
	ticketReq := pb.WXTicketReq{
		AppID:     wxcrypter.AppID,
		InfoType:  wxMsg.InfoType,
		Component: wxMsg.ComponentVerifyTicket,
	}
	_, err = cl.TicketReceived(&ticketReq)
	if err != nil {
		return "success", err
	}
	return "success", nil

}
