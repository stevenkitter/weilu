package mock_services_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/golang/protobuf/proto"
	wxmock "github.com/stevenkitter/weilu/mock_services"
	pb "github.com/stevenkitter/weilu/proto"
)

type rpcMsg struct {
	msg proto.Message
}

func (r *rpcMsg) Matches(msg interface{}) bool {
	m, ok := msg.(proto.Message)
	if !ok {
		return false
	}
	return proto.Equal(m, r.msg)
}

func (r *rpcMsg) String() string {
	return fmt.Sprintf("is %s", r.msg)
}

const encrypted = "7nztsYHzKhcm100LfI9Ch5Up5V9Ibe5bQ6SPZFMgW064Hff6mjNAnt6DYeQRdven5XLfbwGo/N21fdQ7Mzq3VZqUXAIZgz5FAVFpwvzG0DO/Ntj8wqQwxzRen0zwnsSFKfm0tqOgptxaMTBGf1G7XPA+e7axa+8l5j30+9gOdhRAEeIbFPNJQhYlyW/+avaE7jvjZUZoScGPVHHNsnCwUMU2V4aIJDxa22OiIjBCekF6W/xf8lsN4U/lcC7V5ImxpLWpUltHN0RV8SRV7BR0yHR2JbUd+MH6DTC9D/+zQCzCdb2S2PO2BjYJQkRaFCw7zt7aSW7bIzTEcRd3hGIMR2wE0ZJ06XGgIkGu/iC7dJ4D/6eCozZMr3QypciNczCZ20QK3D0fLN0cdSV0iXZZ4iQPIy+A/i3H6nOosfacXoNl/eGPVOab30Q4ecPhxsOsMiSAmxlY5pPrQIzSrHHzB7bCQ6etKCMEoTYGYAwrzDUi3IXRgmEhj4SmPg4CogO5"
const decrypted = "<xml>\n\t\t\t<AppId><![CDATA[wxdd9779d0ca45ea77]]></AppId>\n\t\t\t<CreateTime>1541214207</CreateTime>\n\t\t\t<InfoType><![CDATA[component_verify_ticket]]></InfoType>\n\t\t\t<ComponentVerifyTicket><![CDATA[ticket@@@GTzK1SPD-Ox_UJbmednfythV0KcGryo0XrMiQ2ob9-jShVOb2DwbrktcEfd6bIy0chk1xW_XIODBzTIJ9gvloA]]></ComponentVerifyTicket>\n\t\t</xml>"

var decryptMsgReq = &pb.WXEncryptedMessage{
	Msg:          encrypted,
	MsgSignature: "f4fa85579e0d5d8a5e095909a4eddb86171b60b0",
	Timestamp:    "1544773947",
	Nonce:        "iYReSN2SyhGD3yRN",
}

var ticketReceivedReq = &pb.WXTicketReq{
	AppID:    "",
	InfoType: "",
	Componet: "",
}

func TestDecryptMsg(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockWXClient := wxmock.NewMockWXServiceClient(ctrl)
	mockWXClient.EXPECT().DecryptMsg(
		context.Background(),
		&rpcMsg{msg: decryptMsgReq},
	).Return(&pb.Resp{Code: 200, Data: decrypted}, nil)
	testDecryptMsg(t, mockWXClient)
}

func testDecryptMsg(t *testing.T, client pb.WXServiceClient) {
	r, err := client.DecryptMsg(context.Background(), decryptMsgReq)
	if err != nil || r.Code != 200 || r.Data != decrypted {
		t.Errorf("mocking failed")
	}
}

func TestTicketReceived(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockClient := wxmock.NewMockWXServiceClient(ctrl)
	mockClient.EXPECT().TicketReceived(
		context.Background(),
		&rpcMsg{msg: ticketReceivedReq},
	).Return(&pb.Resp{Code: 200}, nil)
	testTicketReceived(t, mockClient)
}
func testTicketReceived(t *testing.T, client pb.WXServiceClient) {
	r, err := client.TicketReceived(context.Background(), ticketReceivedReq)
	if err != nil || r.Code != 200 {
		t.Errorf("mocking failed")
	}
}
