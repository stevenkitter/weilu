package wx

import (
	"context"
	"log"

	"github.com/jinzhu/gorm"
	bapb "github.com/stevenkitter/protorepo/base"
	wlpb "github.com/stevenkitter/protorepo/weilu"
	"github.com/stevenkitter/weilu/wxcrypter"
)

//Server wx server
type Server struct {
	DB *gorm.DB
}

//DecryptMsg decrypt tencent incoming message
func (s *Server) DecryptMsg(ctx context.Context, req *wlpb.WXEncryptedMessage) (*bapb.Resp, error) {
	e, err := wxcrypter.NewEncrypter(wxcrypter.Token, wxcrypter.EncodingAESKey, wxcrypter.AppID)
	if err != nil {
		log.Printf("NewEncrypter err : %v", err)
		return nil, err
	}
	b, err := e.Decrypt([]byte(req.Msg))
	if err != nil {
		log.Printf("e.Decrypt err : %v", err)
		return nil, err
	}
	resp := &bapb.Resp{
		Code: 200,
		Msg:  "",
		Data: string(b),
	}
	return resp, nil
}

//TicketReceived tencent ticket hander
func (s *Server) TicketReceived(ctx context.Context, req *wlpb.WXTicketReq) (*bapb.Resp, error) {

}
