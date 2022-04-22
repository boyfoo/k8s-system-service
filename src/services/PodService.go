package services

//@Service
type PodService struct {
	PodMap *PodMap        `inject:"-"`
	Common *CommonService `inject:"-"`
}

func (p *PodService) ListByNs(ns string) interface{} {
	return p.PodMap.ListByNs(ns)
}

func NewPodService() *PodService {
	return &PodService{}
}
