package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/skdong/zeus/pkg/websocket"
)

type Message struct {
	DeviceId string
	Enable   bool
	Close    bool
}

type WindWSController struct {
	beego.Controller
}

func (c *WindWSController) Get() {
	handler, err := websocket.WebSocketManager.NewWebSocketHander(
		c.Ctx.ResponseWriter,
		c.Ctx.Request,
		nil)
	if err != nil {
		logs.Warn(err)
		return
	}
	defer handler.Close()
	handler.HandlerMessage()
	logs.Info("close", handler)
}
