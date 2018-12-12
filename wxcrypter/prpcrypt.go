package wxcrypter

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"encoding/binary"
	"math/rand"
	"time"
)

//Prpcrypt msg
type Prpcrypt struct {
	key     []byte
	Encoder PKCS7Encoder
}

//NewPrpcrypt aes
func NewPrpcrypt(key string) (p *Prpcrypt, err error) {
	b, err := base64.StdEncoding.DecodeString(key + "=")
	if err != nil {
		err = ErrorDecodeBase64
		return
	}
	p = &Prpcrypt{
		key:     b,
		Encoder: pkcs7Encoder{},
	}
	return
}

//Encrypt encrypt
func (p *Prpcrypt) Encrypt(appID string, src []byte) (ret string, err error) {
	b, err := aes.NewCipher(p.key)
	if err != nil {
		err = ErrorEncryptAES
		return
	}

	buf := &bytes.Buffer{}
	random := Random()

	_, err = buf.Write(random)
	if err != nil {
		err = ErrorEncryptAES
		return
	}

	err = binary.Write(buf, binary.BigEndian, int32(len(src)))
	if err != nil {
		err = ErrorEncryptAES
		return
	}

	_, err = buf.Write(src)

	if err != nil {
		err = ErrorEncryptAES
		return
	}
	_, err = buf.WriteString(appID)

	if err != nil {
		err = ErrorEncryptAES
		return
	}
	content := buf.Bytes()

	content = p.Encoder.Encode(content)

	c := cipher.NewCBCEncrypter(b, p.key[:16])
	c.CryptBlocks(content, content)
	ret = base64.StdEncoding.EncodeToString(content)
	return
}

//Decrypt src
func (p *Prpcrypt) Decrypt(appID string, src string) (ret []byte, err error) {
	b, err := aes.NewCipher(p.key)
	if err != nil {
		err = ErrorDecryptAES
		return
	}

	content, err := base64.StdEncoding.DecodeString(src)
	if err != nil {
		err = ErrorDecodeBase64
		return
	}

	c := cipher.NewCBCDecrypter(b, p.key[:16])
	c.CryptBlocks(content, content)

	content = p.Encoder.Decode(content)[16:]

	xmlLen := binary.BigEndian.Uint32(content[:4])

	ret = content[4 : xmlLen+4]

	fromAppID := string(content[xmlLen+4:])
	if appID != fromAppID {
		err = ErrorValidateAppID
		return
	}

	return
}

//Random 16 random nonce
func Random() []byte {
	src := []byte("ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789abcdefghijklmnopqrstuvwxyz")
	n := len(src)
	buf := &bytes.Buffer{}

	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	for i := 0; i < 16; i++ {
		index := r.Intn(n)
		buf.WriteByte(src[index])
	}
	return buf.Bytes()
}
