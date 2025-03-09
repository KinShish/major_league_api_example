package controllers

import (
	"major_league_api_example/models"

	beego "github.com/beego/beego/v2/server/web"
)

type WebSocketController struct {
	beego.Controller
}

var hub = models.NewHub()

func init() {
	go hub.Run()
}

// @Title Get
// @Description
// @router /ws [get]
func (ws *WebSocketController) Get() {
	models.ServeWs(hub, ws.Ctx.ResponseWriter, ws.Ctx.Request)
}

// @Title Get
// @Description
// @router /chat [get]
func (ws *WebSocketController) ViewChat() {
	ws.TplName = "home.html"
	ws.Render()
}
