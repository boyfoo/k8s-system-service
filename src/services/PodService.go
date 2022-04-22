package services

import "k8sapi/src/models"

//@Service
type PodService struct {
	PodMap *PodMapStruct  `inject:"-"`
	Common *CommonService `inject:"-"`
}

func (p *PodService) ListByNs(ns string) interface{} {
	list := p.PodMap.ListByNs(ns)
	ret := make([]*models.Pod, 0)
	for _, pod := range list {
		ret = append(ret, &models.Pod{
			Name:       pod.Name,
			NameSpace:  pod.Namespace,
			Images:     p.Common.GetImagesByPod(pod.Spec.Containers),
			NodeName:   pod.Spec.NodeName,
			IP:         []string{pod.Status.PodIP, pod.Status.HostIP},
			Phase:      string(pod.Status.Phase), // pod 当前所处的阶段
			CreateTime: pod.CreationTimestamp.Format("2006-01-02 15:04:05"),
			IsReady:    p.Common.PodIsReady(pod),
		})
	}
	return ret
}

func NewPodService() *PodService {
	return &PodService{}
}
