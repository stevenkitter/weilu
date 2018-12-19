package server

import (
	"github.com/stevenkitter/weilu/database"
	"github.com/stevenkitter/weilu/wxcrypter"
	"time"
)

func (s *Server) SaveTicket(ticket string) error {
	return s.SaveToComponent(ticket, 0, string(database.ComponentVerifyTicket))
}

func (s *Server) SaveTokenExpires(token string, expiresIn int) error {
	return s.SaveToComponent(token, expiresIn, string(database.ComponentAccessToken))
}

func (s *Server) SaveAuthCodeExpires(authCode string, expiresIn int) error {
	return s.SaveToComponent(authCode, expiresIn, string(database.PreAuthCode))
}

func (s *Server) SaveToComponent(component string, expiresIn int, infoType string) error {
	comp := database.WXComponent{}
	return s.DB.Where(database.WXComponent{
		AppID:    wxcrypter.AppID,
		InfoType: infoType,
	}).Assign(database.WXComponent{
		Component: component,
		Expired:   int(time.Now().Unix()) + expiresIn,
	}).FirstOrCreate(&comp).Error
}
