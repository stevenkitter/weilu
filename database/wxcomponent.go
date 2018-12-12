package database

import (
	"github.com/jinzhu/gorm"
)

//ComponentType info type
type ComponentType string

//WXComponent component need to be saved
//ticket update every ten minites coming from tencent server which use https://api.julu666.com/wx api
//token [component_appid component_appsecret component_verify_ticket] => component_access_token
//code used for authorizer web or phone
//auth access token use [component_access_token] => authorizer_access_token
//auth refresh token [component_access_token] => authorizer_refresh_token
const (
	//ComponentVerifyTicket ticket
	ComponentVerifyTicket ComponentType = "component_verify_ticket"
	//ComponentAccessToken token
	ComponentAccessToken ComponentType = "component_access_token"
	//PreAuthCode auth code
	PreAuthCode ComponentType = "pre_auth_code"
	//AuthorizerAccessToken auth access token
	AuthorizerAccessToken ComponentType = "authorizer_access_token"
	//AuthorizerRefreshToken auth refresh token
	AuthorizerRefreshToken ComponentType = "authorizer_refresh_token"
)

//WXComponent save the component
type WXComponent struct {
	gorm.Model
	AppID     string `gorm:"unique_index"`
	InfoType  string
	Component string
}
