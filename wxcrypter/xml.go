package wxcrypter

import (
	"encoding/xml"
)

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
