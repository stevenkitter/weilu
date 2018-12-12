package wx

import (
	"context"

	"github.com/stevenkitter/weilu/wxcrypter"
	"github.com/jinzhu/gorm"
	bapb "github.com/stevenkitter/protorepo/base"
	wlpb "github.com/stevenkitter/protorepo/weilu"
)

//Server wx server
type Server struct {
	DB *gorm.DB
}

//DecryptMsg decrypt tencent incoming message
func (s *Server) DecryptMsg(ctx context.Context, req *wlpb.WXEncryptedMessage) (*bapb.Resp, error) {
	bts := []byte{req.msg}
	var resXML wxcrypter.EncryptedResponseXML
	err = xml.Unmarshal(bts, &resXML)
	if err != nil {
		log.Printf("xml.Unmarshal err : %v", err)
		return
	}
	encrypt := resXML.Encrypt
	msgSignature := resXML.MsgSignature
	format := "<xml><ToUserName><![CDATA[toUser]]></ToUserName><Encrypt><![CDATA[%s]]></Encrypt></xml>"
	fromXML := fmt.Sprintf(format, encrypt)
	b, err = e.Decrypt(msgSignature, timestamp, nonce, []byte(fromXML))
	if err != nil {
		fmt.Println(err)
		return
	}
	// fmt.Printf("e.Decrypt msg : %s \n", string(b))
	if string(b) != text {
		t.Errorf("expected text but get %s", string(b))
	}
}

//TicketReceived tencent ticket hander
func (s *Server) TicketReceived(ctx context.Context, req *wlpb.WXTicketReq) (*bapb.Resp, error) {

}
