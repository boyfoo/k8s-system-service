package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/shenyisyn/goft-gin/goft"
	"io"
	"k8s.io/api/core/v1"
	"k8s.io/client-go/kubernetes"
	"net/http"
)

type PodLogsCtl struct {
	Client *kubernetes.Clientset `inject:"-"`
}

func NewPodLogsCtl() *PodLogsCtl {
	return &PodLogsCtl{}
}
func (this *PodLogsCtl) GetLogs(c *gin.Context) (v goft.Void) {
	ns := c.DefaultQuery("ns", "default")
	podname := c.DefaultQuery("podname", "")
	cname := c.DefaultQuery("cname", "")
	req := this.Client.CoreV1().Pods(ns).GetLogs(podname, &v1.PodLogOptions{Container: cname, Follow: true})
	reader, err := req.Stream(c)
	goft.Error(err)
	for {
		buf := make([]byte, 1024)
		n, err := reader.Read(buf)
		if err != nil && err != io.EOF {
			break
		}
		// http分块发送
		if n > 0 {
			c.Writer.Write([]byte(string(buf[0:n])))
			c.Writer.(http.Flusher).Flush()
		}
	}

	return
}
func (*PodLogsCtl) Name() string {
	return "PodLogsCtl"
}

func (this *PodLogsCtl) Build(goft *goft.Goft) {
	goft.Handle("GET", "/pods/logs", this.GetLogs)
}
