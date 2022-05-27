package main

import (
	"context"
	"fmt"
	v1 "k8s.io/api/apps/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/util/json"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {
	// 修改替换操作
	// 新增了一个containers 修改的时候会自动把原本的和新增的合并成两个 containers的 pod
	// "$patch": "delete" 这个加上去删除 去掉就合并新增
	// language=json
	s := `{
		"spec": {
			"template": {
				"spec": {
					"containers": [
						{
							"name": "redis",
							"image": "redis:5-alpine",
							"$patch": "delete"
						}
					]
				}
			}
		}
	}`
	data := []byte(s)
	clientset := kubernetes.NewForConfigOrDie(GetConfig())
	_, err := clientset.AppsV1().Deployments("default").Patch(context.Background(), "nginx-deployment-name01", types.StrategicMergePatchType, data, metav1.PatchOptions{})
	fmt.Println(err)
}

func 动态客户端() {

	//clientset := ClientSet()
	//
	//list, err := clientset.CoreV1().Pods("istio-system").List(context.Background(), v1.ListOptions{})
	//if err != nil {
	//	panic(err)
	//}
	//for _, item := range list.Items {
	//	fmt.Println(item.Name)
	//}

	client := DynamicClient()
	list, err := client.Resource(v1.SchemeGroupVersion.WithResource("deployments")).List(context.Background(), metav1.ListOptions{})
	if err != nil {
		panic(err)
	}
	//for _, item := range list.Items {
	//	fmt.Println(item.GetName())
	//}
	b, _ := list.MarshalJSON()

	depList := &v1.DeploymentList{}
	// 结果转换成结构体
	json.Unmarshal(b, depList)

	for _, item := range depList.Items {
		fmt.Println(item.Name)
	}
}

func ClientSet() *kubernetes.Clientset {
	clientset, err := kubernetes.NewForConfig(GetConfig())
	if err != nil {
		panic(err)
	}

	return clientset
}

func GetConfig() *rest.Config {
	config, err := clientcmd.BuildConfigFromFlags("", clientcmd.RecommendedHomeFile)
	if err != nil {
		panic(err)
	}
	return config
}

// DynamicClient 动态客户端
func DynamicClient() dynamic.Interface {
	return dynamic.NewForConfigOrDie(GetConfig())
}
