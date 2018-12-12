package wxcrypter

//Encrypter encrypter
type Encrypter struct {
	prpcrypter     *Prpcrypt
	token          string
	encodingAesKey string
	appID          string
}

//NewEncrypter new encrypter
func NewEncrypter(token, encodingAesKey, appID string) (e *Encrypter, err error) {
	if len(encodingAesKey) != 43 {
		err = ErrorIllegalAesKey
		return
	}

	p, err := NewPrpcrypt(encodingAesKey)
	if err != nil {
		return
	}

	e = &Encrypter{
		prpcrypter:     p,
		token:          token,
		appID:          appID,
		encodingAesKey: encodingAesKey,
	}
	return
}

//Encrypt encrypt msg
func (e *Encrypter) Encrypt(replyMsg []byte, timestamp, nonce string) (b []byte, err error) {
	encrypt, err := e.prpcrypter.Encrypt(e.appID, replyMsg)
	if err != nil {
		return
	}

	signature := Sha1(e.token, timestamp, nonce, encrypt)

	b, err = GenerateResponseXML(encrypt, signature, timestamp, nonce)
	return
}

//Decrypt decrypt msg
func (e *Encrypter) Decrypt(msgSignature, timestamp, nonce string, data []byte) (b []byte, err error) {
	reqXML, err := ParseRequestXML(data)
	if err != nil {
		return
	}

	signature := Sha1(e.token, timestamp, nonce, reqXML.Encrypt)
	if signature != msgSignature {
		err = ErrorValidateSignature
		return
	}
	b, err = e.prpcrypter.Decrypt(e.appID, reqXML.Encrypt)
	return
}
