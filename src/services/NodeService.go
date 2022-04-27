package services

import "k8sapi/src/models"

//@service
type NodeService struct {
	NodeMap *NodeMapStruct `inject:"-"`
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
			Name:       node.Name,
			IP:         node.Status.Addresses[0].Address,
			HostName:   node.Status.Addresses[1].Address,
			CreateTime: node.CreationTimestamp.Format("2006-01-02 15:04:05"),
		}
	}
	return ret
}
