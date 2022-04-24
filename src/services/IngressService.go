package services

import (
	"context"
	"k8s.io/api/networking/v1beta1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
	"k8s.io/client-go/kubernetes"
	"k8sapi/src/models"
	"strconv"
)

//@Service
type IngressService struct {
	Common *CommonService        `inject:"-"`
	Client *kubernetes.Clientset `inject:"-"`
}

func (i *IngressService) PostIngress(post *models.IngressPost) error {

	className := "nginx"
	ingressRules := []v1beta1.IngressRule{}
	for _, r := range post.Rules {
		httpRuleValue := &v1beta1.HTTPIngressRuleValue{}
		rulePaths := make([]v1beta1.HTTPIngressPath, 0)

		for _, pathCfg := range r.Paths {
			port, err := strconv.Atoi(pathCfg.Port)
			if err != nil {
				return err
			}
			rulePaths = append(rulePaths, v1beta1.HTTPIngressPath{
				Path: pathCfg.Path,
				Backend: v1beta1.IngressBackend{
					ServiceName: pathCfg.SvcName,
					ServicePort: intstr.FromInt(port),
				},
			})
		}
		httpRuleValue.Paths = rulePaths
		rule := v1beta1.IngressRule{
			Host: r.Host,
			IngressRuleValue: v1beta1.IngressRuleValue{
				HTTP: httpRuleValue,
			},
		}
		ingressRules = append(ingressRules, rule)
	}

	i.Client.NetworkingV1beta1().Ingresses(post.Namespace).Create(context.Background(), &v1beta1.Ingress{
		TypeMeta: v1.TypeMeta{
			Kind:       "Ingress",
			APIVersion: "networking.k8s.io/v1beta1",
		},
		ObjectMeta: v1.ObjectMeta{
			Name:      post.Name,
			Namespace: post.Namespace,
		},
		Spec: v1beta1.IngressSpec{
			IngressClassName: &className,
			Rules:            ingressRules,
		},
	}, v1.CreateOptions{})
	return nil
}

func NewIngressService() *IngressService {
	return &IngressService{}
}
