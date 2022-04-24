package services

import (
	"k8s.io/client-go/kubernetes"
	"k8sapi/src/models"
)

//@service
type SecretService struct {
	Client    *kubernetes.Clientset `inject:"-"`
	SecretMap *SecretMapStruct      `inject:"-"`
}

func NewSecretService() *SecretService {
	return &SecretService{}
}

//前台用于显示Secret列表
func (this *SecretService) ListSecret(ns string) []*models.SecretModel {

	list := this.SecretMap.ListAll(ns)
	ret := make([]*models.SecretModel, len(list))
	for i, item := range list {
		ret[i] = &models.SecretModel{
			Name:       item.Name,
			CreateTime: item.CreationTimestamp.Format("2006-01-02 15:04:05"),
			NameSpace:  item.Namespace,
			Type:       models.SECRET_TYPE[string(item.Type)], // 类型的翻译
		}
	}
	return ret
}
