package endpoint

import (
	"github.com/gin-gonic/gin"
	"github.com/stevenkitter/weilu/api/data"
	"github.com/stevenkitter/weilu/client"
	pb "github.com/stevenkitter/weilu/proto"
	"log"
	"os"
)

type AuthURLQuery struct {
	Device      string `form:"device" binding:"required"`
	AuthType    string `form:"authType" binding:"required"`
	RedirectURL string `form:"redirectUrl" binding:"required"`
}

// AuthURL 授权路径
func AuthURL(c *gin.Context) (interface{}, error) {
	queries := AuthURLQuery{}
	err := c.ShouldBindQuery(&queries)
	if err != nil {
		log.Printf("c.ShouldBindQuery err : %v", err)
		return data.NewErrorResp(err), err
	}
	cl := client.Client{Address: os.Getenv("WX_SERVER_ADDRESS")}
	res, err := cl.AuthURL(&pb.GetAuthURLReq{
		Device:      queries.Device,
		AuthType:    queries.AuthType,
		RedirectURL: queries.RedirectURL,
	})
	if err != nil {
		log.Printf("cl.AuthURL err : %v", err)
		return data.NewErrorResp(err), err
	}
	return data.NewFineResp("", res.Data), nil
}
