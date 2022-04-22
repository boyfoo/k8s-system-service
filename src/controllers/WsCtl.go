package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/shenyisyn/goft-gin/goft"
	"k8sapi/src/wscore"
	"log"
)

type WsCtl struct {
}

func (w *WsCtl) Connect(c *gin.Context) (v goft.Void) {
	client, err := wscore.Upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Println(err)
		return
	} else {
		wscore.ClientMap.Store(client)
		return
	}
}

func NewWsCtl() *WsCtl {
	return &WsCtl{}
}

func (this *WsCtl) Build(goft *goft.Goft) {
	//路由
	goft.Handle("GET", "/ws", this.Connect)
}
func (*WsCtl) Name() string {
	return "WsCtl"
}
