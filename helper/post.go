package helper

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

//Post Json to bytes
// @params url dict
func PostJson(url string, dict map[string]interface{}) ([]byte, error) {
	bts, err := json.Marshal(dict)
	if err != nil {
		return nil, err
	}
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(bts))
	if err != nil {
		return nil, err
	}
	defer func() {
		if err := resp.Body.Close(); err != nil {
			log.Printf("resp.Body.Close err : %v", err)
		}
	}()
	result, err := ioutil.ReadAll(resp.Body)
	return result, err
}
