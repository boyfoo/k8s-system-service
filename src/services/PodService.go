package services

//@Service
type PodService struct {
	PodMap *PodMapStruct  `inject:"-"`
	Common *CommonService `inject:"-"`
}

func (p *PodService) ListByNs(ns string) interface{} {
	return p.PodMap.ListByNs(ns)
}

func NewPodService() *PodService {
	return &PodService{}
}
