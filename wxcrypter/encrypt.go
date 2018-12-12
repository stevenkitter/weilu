package wxcrypter

// this package is used for contact with tencent api
// this file is used for encrypt and decrypt the tencent incoming encrypted message
import (
	"crypto/rand"
	"log"
)

//RandBytes random bytes of size
func RandBytes(size int) ([]byte, error) {
	r := make([]byte, size)
	_, err := rand.Read(r)
	if err != nil {
		log.Printf("rand.Read err: %v", err)
		return r, err
	}
	return r, nil
}
