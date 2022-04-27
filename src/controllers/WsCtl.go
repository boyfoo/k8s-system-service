package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/shenyisyn/goft-gin/goft"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/remotecommand"
	"k8sapi/src/helpers"
	"k8sapi/src/wscore"
	"log"
)

//@Controller
type WsCtl struct {
	Client *kubernetes.Clientset `inject:"-"`
	Config *rest.Config          `inject:"-"`
}

func NewWsCtl() *WsCtl {
	return &WsCtl{}
}

func (this *WsCtl) Connect(c *gin.Context) (v goft.Void) {
	client, err := wscore.Upgrader.Upgrade(c.Writer, c.Request, nil) //升级
	if err != nil {
		log.Println(err)
		return
	} else {
		wscore.ClientMap.Store(client)

		return
	}
}
func (this *WsCtl) PodConnect(c *gin.Context) (v goft.Void) {
	wsClient, err := wscore.Upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		return
	}
	fmt.Println("创建pod wsshell 链接")
	shellClient := wscore.NewWsShellClient(wsClient)
	err = helpers.HandleCommand(this.Client, this.Config, []string{"sh"}).
		Stream(remotecommand.StreamOptions{
			Stdin:  shellClient,
			Stdout: shellClient,
			Stderr: shellClient,
			Tty:    true,
		})
	if err != nil {
		fmt.Println(err)
	}
	return
}
func (this *WsCtl) Build(goft *goft.Goft) {
	goft.Handle("GET", "/ws", this.Connect)
	goft.Handle("GET", "/podws", this.PodConnect)
}
func (this *WsCtl) Name() string {
	return "WsCtl"
}
