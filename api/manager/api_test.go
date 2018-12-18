package manager_test

import (
	"bytes"
	"encoding/xml"
	pb "github.com/stevenkitter/weilu/proto"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"testing"
)

func TestWXEndpoint(t *testing.T) {
	var url = os.Getenv("WX_API_SERVER_ADDRESS")
	url = url + "?signature=3bb7c0a3ec15d3ffe2158e3995a3446e9ce91fa5&timestamp=1545102154&nonce=1212295081&encrypt_type=aes&msg_signature=6d20f174608bc74861a52893469d7280232f7e49"
	postData := &pb.WXEncryptedMessage{
		Msg:          "",
		MsgSignature: "",
		Timestamp:    "",
		Nonce:        "",
	}
	bts, err := xml.Marshal(postData)
	if err != nil {
		t.Errorf("xml.Marshal err : %v", err)
	}

	resp, err := http.Post(url, "application/xml", bytes.NewBuffer(bts))
	if err != nil {
		t.Errorf("Post url err : %v", err)
	}
	defer func() {
		err := resp.Body.Close()
		if err != nil {
			log.Printf("resp.Body.Close err : %v", err)
		}
	}()
	respBites, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Errorf("ioutil.ReadAll err : %v", err)
	}
	if string(respBites) != "success" {
		t.Errorf("expected response is %s, but get %s", "success", string(respBites))
	}

}
