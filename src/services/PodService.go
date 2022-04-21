package services

import (
	"k8sapi/src/core"
)

//@Service
type PodService struct {
	PodMap *core.PodMap   `inject:"-"`
	Common *CommonService `inject:"-"`
}

func (p *PodService) ListByNs(ns string) interface{} {
	return p.PodMap.ListByNs(ns)
}

func NewPodService() *PodService {
	return &PodService{}
}
