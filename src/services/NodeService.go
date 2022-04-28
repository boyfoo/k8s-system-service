package services

import (
	"k8sapi/src/helpers"
	"k8sapi/src/models"
)

//@service
type NodeService struct {
	NodeMap *NodeMapStruct `inject:"-"`
	PodMap  *PodMapStruct  `inject:"-"`
}

func NewNodeService() *NodeService {
	return &NodeService{}
}

//显示所有节点
func (this *NodeService) ListAllNodes() []*models.NodeModel {

	list := this.NodeMap.ListAll()
	ret := make([]*models.NodeModel, len(list))
	for i, node := range list {

		ret[i] = &models.NodeModel{
			Name:     node.Name,
			IP:       node.Status.Addresses[0].Address,
			HostName: node.Status.Addresses[1].Address,
			Lables:   helpers.FilterLables(node.Labels),
			Taints:   helpers.FilterTaints(node.Spec.Taints),
			Capacity: models.NewNodeCapacity(node.Status.Capacity.Cpu().Value(),
				node.Status.Capacity.Memory().Value(), node.Status.Capacity.Pods().Value()),
			Usage:      models.NewNodeUsage(this.PodMap.GetNum(node.Name)),
			CreateTime: node.CreationTimestamp.Format("2006-01-02 15:04:05"),
		}
	}
	return ret
}
