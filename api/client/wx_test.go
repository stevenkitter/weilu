package client_test

import (
	"testing"

	"github.com/stevenkitter/weilu/api/client"
	pb "github.com/stevenkitter/weilu/proto"
)

func TestDecryptMsg(t *testing.T) {
	encryped := "7nztsYHzKhcm100LfI9Ch5Up5V9Ibe5bQ6SPZFMgW064Hff6mjNAnt6DYeQRdven5XLfbwGo/N21fdQ7Mzq3VZqUXAIZgz5FAVFpwvzG0DO/Ntj8wqQwxzRen0zwnsSFKfm0tqOgptxaMTBGf1G7XPA+e7axa+8l5j30+9gOdhRAEeIbFPNJQhYlyW/+avaE7jvjZUZoScGPVHHNsnCwUMU2V4aIJDxa22OiIjBCekF6W/xf8lsN4U/lcC7V5ImxpLWpUltHN0RV8SRV7BR0yHR2JbUd+MH6DTC9D/+zQCzCdb2S2PO2BjYJQkRaFCw7zt7aSW7bIzTEcRd3hGIMR2wE0ZJ06XGgIkGu/iC7dJ4D/6eCozZMr3QypciNczCZ20QK3D0fLN0cdSV0iXZZ4iQPIy+A/i3H6nOosfacXoNl/eGPVOab30Q4ecPhxsOsMiSAmxlY5pPrQIzSrHHzB7bCQ6etKCMEoTYGYAwrzDUi3IXRgmEhj4SmPg4CogO5"
	var decryptMsgReq = &pb.WXEncryptedMessage{
		Msg:          encrypted,
		MsgSignature: "f4fa85579e0d5d8a5e095909a4eddb86171b60b0",
		Timestamp:    "1544773947",
		Nonce:        "iYReSN2SyhGD3yRN",
	}
	cl := client.Client{}
	res,err := cl.DecryptMsg(decryptMsgReq)
	if err != nil {
		t.Errorf("mocking failed")
	}
}
