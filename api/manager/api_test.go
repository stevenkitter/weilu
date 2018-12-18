package manager_test

import (
	"bytes"
	"encoding/xml"
	"github.com/stevenkitter/weilu/wxcrypter"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"testing"
)

func TestWXEndpoint(t *testing.T) {
	var url = os.Getenv("WX_API_SERVER_ADDRESS")
	url = url + "/wx" +
		"?signature=23b6cf3b709fd0cc0f211fe94881a1610e8f59b4&timestamp=1545107020&nonce=1818822260&encrypt_type=aes&msg_signature=231168c3d0ae6ef5ce7eb391971ff18ed3e2f847"
	postData := &wxcrypter.EncryptedRequestXML{
		Encrypt: "j2x5ykY7tarUa5wMRkFHhgAwJTZMB6ZkuXpo4qNHCO7wYbqLaRg6U26gvv2NWbGqq4w8cQlP6xE+OoMv5MAr79oSeq1L/FOOG+qHZZD2ZN46Lib6orkY3n14+/g/4SSa7C2Sh/giA7KS/C0WzazNChF1c9n7OdRfhhNccvLUK0ZOKfXluyTCSM2gf9eR4TYP9mqs0ogyBHi+oug/yRGE+h+WNT3UyipA4mibZkj8uyLNsxvy/K/zKwmWoFPk9OYQg1LSu7PQLbPEYN1gcOmXn/Q3QddlbqLhg5/fr3Qh3+TPn2WqHuMBZ+3gku/cINGvn5fRLzfg8QmkqzR9ocuYfrQIW01Q21wnjubZ39P4QHIw7kJy9oWasDy83V87tnHqkFvUhm52Bfm7LauwoRLM4wQSlxxoZYwKDrJbmiHT6x/fh6lxNysZlTZaAUvUpboeSsVcApbxGUNsGx6u8ToDqQ==",
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
