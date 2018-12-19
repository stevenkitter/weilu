package httpClient

import (
	"encoding/json"
	"github.com/pkg/errors"
	"github.com/stevenkitter/weilu/helper"
	"github.com/stevenkitter/weilu/wxcrypter"
)

const (
	WXBaseURL      = "https://api.weixin.qq.com/"
	WXComponentURL = WXBaseURL + "cgi-bin/component/"
	TokenURL       = WXComponentURL + "api_component_token"
	PreAuthCodeURL = WXComponentURL + "api_create_preauthcode?component_access_token="
)

type Client struct {
}

//获取第三方平台component_access_token
func (cl *Client) RequestAccessToken(ticket string) (string, int, error) {
	postData := map[string]interface{}{
		"component_appid":         wxcrypter.AppID,
		"component_appsecret":     wxcrypter.AppSecret,
		"component_verify_ticket": ticket,
	}
	data, err := helper.PostJson(TokenURL, postData)
	if err != nil {
		return "", 0, err
	}
	respDict := map[string]interface{}{}
	err = json.Unmarshal(data, &respDict)
	if err != nil {
		return "", 0, err
	}
	acToken, ok := respDict["component_access_token"]
	if !ok {
		return "", 0, errors.New(respDict["errmsg"].(string))
	}
	expiresIn, ok := respDict["expires_in"]
	if !ok {
		return "", 0, errors.New(respDict["errmsg"].(string))
	}
	return acToken.(string), int(expiresIn.(float64)), nil
}

func (cl *Client) RequestPreAuthCode(token string) (string, int, error) {
	postData := map[string]interface{}{
		"component_appid": wxcrypter.AppID,
	}
	data, err := helper.PostJson(PreAuthCodeURL+token, postData)
	if err != nil {
		return "", 0, err
	}
	respDict := map[string]interface{}{}
	err = json.Unmarshal(data, &respDict)
	if err != nil {
		return "", 0, err
	}
	authCode, ok := respDict["pre_auth_code"]
	if !ok {
		return "", 0, errors.New(respDict["errmsg"].(string))
	}
	expiresIn, ok := respDict["expires_in"]
	if !ok {
		return "", 0, errors.New(respDict["errmsg"].(string))
	}
	return authCode.(string), int(expiresIn.(float64)), nil
}
