package manager

import (
	"github.com/gin-gonic/gin"
	"github.com/stevenkitter/weilu/api/endpoint"
)

//Manager wrap gin engine
type Manager struct {
	Engine *gin.Engine
}

//NewManager init manager and use middleware
func NewManager() Manager {
	engine := gin.Default()
	engine.Use(CORSMiddleware())
	m := Manager{
		Engine: engine,
	}
	return m
}

//Route router the api
func (m *Manager) Route() {
	m.Engine.POST("/wx", WrapWXHandler(endpoint.WXReceiveEndpoint))
	m.Engine.POST("/wx/:appid", WrapWXHandler(endpoint.WXReceiveEndpoint))
	v1 := m.Engine.Group("/v1")
	{
		v1.POST("/login", WrapHandler(endpoint.WXReceiveEndpoint))
	}
	v2 := m.Engine.Group("/v2")
	{
		v2.POST("/login", WrapHandler(endpoint.WXReceiveEndpoint))
	}
}

//Run run engine
func (m *Manager) Run(port string) error {
	m.Route()
	return m.Engine.Run(port)
}
