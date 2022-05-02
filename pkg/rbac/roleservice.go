package rbac

//@Service
type RoleService struct {
	RoleMap        *RoleMapStruct        `inject:"-"`
	RoleBindingMap *RoleBindingMapStruct `inject:"-"`
}

func NewRoleService() *RoleService {
	return &RoleService{}
}
func (this *RoleService) ListRoles(ns string) []*RoleModel {
	list := this.RoleMap.ListAll(ns)
	ret := make([]*RoleModel, len(list))
	for i, item := range list {
		ret[i] = &RoleModel{
			Name:       item.Name,
			CreateTime: item.CreationTimestamp.Format("2006-01-02 15:04:05"),
			NameSpace:  item.Namespace,
		}
	}
	return ret
}

func (this *RoleService) ListRoleBindings(ns string) []*RoleBindingModel {
	list := this.RoleBindingMap.ListAll(ns)
	ret := make([]*RoleBindingModel, len(list))
	for i, item := range list {

		ret[i] = &RoleBindingModel{

			Name:       item.Name,
			CreateTime: item.CreationTimestamp.Format("2006-01-02 15:04:05"),
			NameSpace:  item.Namespace,
			Subject:    item.Subjects,
		}
	}
	return ret
}
