package wxcrypter

import (
	"encoding/xml"
	"fmt"
	"log"
	"testing"
	"time"
)

func TestCrypt(t *testing.T) {
	timestamp := fmt.Sprintf("%d", time.Now().Unix())
	nonce := string(Random())
	text :=
		`<xml>
			<AppId><![CDATA[wxdd9779d0ca45ea77]]></AppId>
			<CreateTime>1541214207</CreateTime>
			<InfoType><![CDATA[component_verify_ticket]]></InfoType>
			<ComponentVerifyTicket><![CDATA[ticket@@@GTzK1SPD-Ox_UJbmednfythV0KcGryo0XrMiQ2ob9-jShVOb2DwbrktcEfd6bIy0chk1xW_XIODBzTIJ9gvloA]]></ComponentVerifyTicket>
		</xml>`
	e, err := NewEncrypter(Token, EncodingAESKey, AppID)
	if err != nil {
		log.Printf("NewEncrypter err : %v", err)
		return
	}
	b, err := e.Encrypt([]byte(text), timestamp, nonce)
	if err != nil {
		log.Printf("e.Encrypt err : %v", err)
		return
	}
	// fmt.Printf("encrypt msg : %s \n", string(b))
	var resXML EncryptedResponseXML
	err = xml.Unmarshal(b, &resXML)
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
