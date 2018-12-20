package endpoint_test

import (
	"encoding/json"
	"github.com/stevenkitter/weilu/api/data"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"testing"
)

func TestAuthURL(t *testing.T) {
	var url = os.Getenv("WX_API_SERVER_ADDRESS") + "/v1/auth_url?"
	params := map[string]string{
		"device":      "web",
		"authType":    "1",
		"redirectUrl": "https://weilu.julu666.com/wx",
	}
	paramsArr := make([]string, 0)
	for key, value := range params {
		item := key + "=" + value
		paramsArr = append(paramsArr, item)
	}
	url += strings.Join(paramsArr, "&")
	resp, err := http.Get(url)
	if err != nil {
		t.Errorf("http.Get(url) err : %v", err)
	}
	defer func() {
		if err := resp.Body.Close(); err != nil {
			log.Printf("resp.Body.Close err %v", err)
		}
	}()
	da, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Errorf("ioutil.ReadAll(resp.Body) err : %v", err)
	}
	res := data.Resp{}
	err = json.Unmarshal(da, &res)
	if err != nil {
		t.Errorf("json.Unmarshal(da, &res) err : %v", err)
	}
}
