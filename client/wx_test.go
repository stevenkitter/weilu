package client_test

import (
	"github.com/stevenkitter/weilu/client"
	pb "github.com/stevenkitter/weilu/proto"
	"github.com/stevenkitter/weilu/wxcrypter"
	"os"
	"testing"
)

func TestClient_DecryptMsg(t *testing.T) {
	encrypted := "7nztsYHzKhcm100LfI9Ch5Up5V9Ibe5bQ6SPZFMgW064Hff6mjNAnt6DYeQRdven5XLfbwGo/N21fdQ7Mzq3VZqUXAIZgz5FAVFpwvzG0DO/Ntj8wqQwxzRen0zwnsSFKfm0tqOgptxaMTBGf1G7XPA+e7axa+8l5j30+9gOdhRAEeIbFPNJQhYlyW/+avaE7jvjZUZoScGPVHHNsnCwUMU2V4aIJDxa22OiIjBCekF6W/xf8lsN4U/lcC7V5ImxpLWpUltHN0RV8SRV7BR0yHR2JbUd+MH6DTC9D/+zQCzCdb2S2PO2BjYJQkRaFCw7zt7aSW7bIzTEcRd3hGIMR2wE0ZJ06XGgIkGu/iC7dJ4D/6eCozZMr3QypciNczCZ20QK3D0fLN0cdSV0iXZZ4iQPIy+A/i3H6nOosfacXoNl/eGPVOab30Q4ecPhxsOsMiSAmxlY5pPrQIzSrHHzB7bCQ6etKCMEoTYGYAwrzDUi3IXRgmEhj4SmPg4CogO5"
	var decryptMsgReq = &pb.WXEncryptedMessage{
		Msg:          encrypted,
		MsgSignature: "f4fa85579e0d5d8a5e095909a4eddb86171b60b0",
		Timestamp:    "1544773947",
		Nonce:        "iYReSN2SyhGD3yRN",
	}
	const expected = "<xml>\n\t\t\t<AppId><![CDATA[wxdd9779d0ca45ea77]]></AppId>\n\t\t\t<CreateTime>1541214207</CreateTime>\n\t\t\t<InfoType><![CDATA[component_verify_ticket]]></InfoType>\n\t\t\t<ComponentVerifyTicket><![CDATA[ticket@@@GTzK1SPD-Ox_UJbmednfythV0KcGryo0XrMiQ2ob9-jShVOb2DwbrktcEfd6bIy0chk1xW_XIODBzTIJ9gvloA]]></ComponentVerifyTicket>\n\t\t</xml>"

	cl := client.Client{
		Address: os.Getenv("WX_SERVER_ADDRESS"),
	}
	res, err := cl.DecryptMsg(decryptMsgReq)
	if err != nil {
		t.Errorf("cl.DecryptMsg failed Err: %v", err)
	}
	if res.Data != expected {
		t.Errorf("expected %s, but get %s", expected, res.Data)
	}
}

func TestClient_TicketReceived(t *testing.T) {
	req := &pb.WXTicketReq{
		AppID:     "wxdd9779d0ca45ea77",
		InfoType:  "component_verify_ticket",
		Component: "ticket@@@OR-oiXvB5nbXy5VAf2DXqWyf3zABMWUI0BBd33f8Zgdrd_YdbuqqaU_tI3j_VbV063UCMqKc8TpTQeuKiI7Hig",
	}
	cl := client.Client{
		Address: os.Getenv("WX_SERVER_ADDRESS"),
	}
	res, err := cl.TicketReceived(req)
	if err != nil {
		t.Errorf("cl.TicketReceived err : %v", err)
	}
	if res.Code != 200 {
		t.Errorf("TicketReceived client failed")
	}
}

func TestClient_Ticket(t *testing.T) {
	req := &pb.GetTicketReq{
		AppID: "wxdd9779d0ca45ea77",
	}
	cl := client.Client{
		Address: os.Getenv("WX_SERVER_ADDRESS"),
	}
	res, err := cl.Ticket(req)
	if err != nil {
		t.Errorf("cl.Ticket err : %v", err)
	}
	if res.Code != 200 {
		t.Errorf("GetTicketReq client failed")
	}
	if res.Data == "" {
		t.Logf("ticket record not finded i will import sql later")
	}
}

func TestClient_AccessToken(t *testing.T) {
	req := &pb.GetAccessTokenReq{
		AppID: "wxdd9779d0ca45ea77",
	}
	cl := client.Client{
		Address: os.Getenv("WX_SERVER_ADDRESS"),
	}
	res, err := cl.AccessToken(req)
	if err != nil {
		t.Errorf("cl.AccessToken err : %v", err)
	}
	if res.Code != 200 {
		t.Errorf("AccessToken client failed")
	}
}
func TestClient_PreAuthCode(t *testing.T) {
	req := &pb.GetPreAuthCodeReq{
		AppID: wxcrypter.AppID,
	}
	cl := client.Client{
		Address: os.Getenv("WX_SERVER_ADDRESS"),
	}
	res, err := cl.PreAuthCode(req)
	if err != nil {
		t.Errorf("cl.PreAuthCode err : %v", err)
	}
	if res.Code != 200 {
		t.Errorf("PreAuthCode client failed")
	}
}
