package services

import (
	"fmt"
	"k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
)

//@Service
type CommonService struct {
}

func NewCommonService() *CommonService {
	return &CommonService{}
}
func (this *CommonService) GetImages(dep v1.Deployment) string {
	return this.GetImagesByPod(dep.Spec.Template.Spec.Containers)
}
func (*CommonService) GetImagesByPod(containers []corev1.Container) string {
	images := containers[0].Image
	if imgLen := len(containers); imgLen > 1 {
		images += fmt.Sprintf("+其他%d个镜像", imgLen-1)
	}
	return images
}

// pod 是否准备好
func (c *CommonService) PodIsReady(pod *corev1.Pod) bool {
	// pod阶段值
	if pod.Status.Phase != corev1.PodRunning {
		return false
	}
	// 是否已经被调度到podScheduled    pod中容器是否准备就绪containerReady    所有init容器已启动initialized   POD可以提供服务Ready
	for _, condition := range pod.Status.Conditions {
		if condition.Status != corev1.ConditionTrue {
			return false
		}
	}

	for _, gate := range pod.Spec.ReadinessGates {
		for _, condition := range pod.Status.Conditions {
			if condition.Type == gate.ConditionType && condition.Status != corev1.ConditionTrue {
				return false
			}
		}
	}

	return true
}
