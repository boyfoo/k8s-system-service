package main

import (
	"github.com/shenyisyn/goft-gin/goft"
	"k8s-system-service/src/configs"
	"k8s-system-service/src/controllers"
)

func main() {
	goft.Ignite().Config(configs.NewK8sConfig()).Mount("v1", controllers.NewDeploymentCtl()).Launch()
}
