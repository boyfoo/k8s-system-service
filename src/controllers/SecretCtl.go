package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/shenyisyn/goft-gin/goft"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8sapi/src/models"
	"k8sapi/src/services"
)

//@restcontroller
type SecretCtl struct {
	SecretMap     *services.SecretMapStruct `inject:"-"`
	SecretService *services.SecretService   `inject:"-"`
	Client        *kubernetes.Clientset     `inject:"-"`
}

func NewSecretCtl() *SecretCtl {
	return &SecretCtl{}
}
func (*SecretCtl) Name() string {
	return "SecretCtl"
}

//DELETE /ingress?ns=xx&name=xx
func (this *SecretCtl) RmSecret(c *gin.Context) goft.Json {
	ns := c.DefaultQuery("ns", "default")
	name := c.DefaultQuery("name", "")
	goft.Error(this.Client.CoreV1().Secrets(ns).
		Delete(c, name, v1.DeleteOptions{}))
	return gin.H{
		"code": 20000,
		"data": "OK",
	}
}
func (this *SecretCtl) ListAll(c *gin.Context) goft.Json {
	ns := c.DefaultQuery("ns", "default")
	return gin.H{
		"code": 20000,
		"data": this.SecretService.ListSecret(ns), //暂时 不分页
	}
}
func (this *SecretCtl) PostSecret(c *gin.Context) goft.Json {
	postModel := &models.PostSecretModel{}
	err := c.ShouldBindJSON(postModel)
	goft.Error(err)
	_, err = this.Client.CoreV1().Secrets(postModel.NameSpace).Create(c, &corev1.Secret{
		ObjectMeta: v1.ObjectMeta{
			Name:      postModel.Name,
			Namespace: postModel.NameSpace,
		},
		Type:       corev1.SecretType(postModel.Type),
		StringData: postModel.Data,
	}, v1.CreateOptions{})
	goft.Error(err)
	return gin.H{
		"code": 20000,
		"data": "OK",
	}
}

func (this *SecretCtl) Build(goft *goft.Goft) {
	goft.Handle("GET", "/secrets", this.ListAll)
	goft.Handle("DELETE", "/secrets", this.RmSecret)
	goft.Handle("POST", "/secrets", this.PostSecret)
}
