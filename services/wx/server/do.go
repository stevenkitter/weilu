package server

import (
	"github.com/stevenkitter/weilu/database"
	"github.com/stevenkitter/weilu/wxcrypter"
	"time"
)

func (s *Server) SaveTokenExpires(token string, expiresIn int) error {
	comp := database.WXComponent{}
	err := s.DB.Where(database.WXComponent{
		AppID:    wxcrypter.AppID,
		InfoType: string(database.ComponentAccessToken),
	}).Assign(database.WXComponent{
		Component: token,
		Expired:   int(time.Now().Unix()) + expiresIn,
	}).FirstOrCreate(&comp).Error
	return err
}
