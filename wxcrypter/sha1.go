package wxcrypter

import (
	"crypto/sha1"
	"encoding/hex"
	"io"
	"sort"
)

//Sha1 sha1 the paramters
func Sha1(token, timestamp, nonce, msg string) string {
	sl := []string{token, timestamp, nonce, msg}
	sort.Strings(sl)

	h := sha1.New()
	for _, s := range sl {
		io.WriteString(h, s)
	}
	encode := h.Sum(nil)

	return hex.EncodeToString(encode)

}

//Signature signature
func Signature(token, timestamp, nonce string) string {
	sl := []string{token, timestamp, nonce}
	sort.Strings(sl)

	h := sha1.New()
	for _, s := range sl {
		io.WriteString(h, s)
	}
	encode := h.Sum(nil)

	return hex.EncodeToString(encode)
}
