package helpers

import (
	"fmt"
	v1 "k8s.io/api/core/v1"
	"regexp"
)

const hostPattern = "[a-zA-Z0-9][-a-zA-Z0-9]{0,62}(\\.[a-zA-Z0-9][-a-zA-Z0-9]{0,62})+\\.?"

func showLable(key string) bool {
	return !regexp.MustCompile(hostPattern).MatchString(key)
	//return true
}

func FilterTaints(taints []v1.Taint) (ret []string) {
	for _, taint := range taints {
		if showLable(taint.Key) {
			ret = append(ret, fmt.Sprintf("%s=%s:%s", taint.Key, taint.Value, taint.Effect))
		}
	}
	return
}

//过滤 要显示的标签
func FilterLables(labels map[string]string) (ret []string) {
	for k, v := range labels {
		if showLable(k) {
			ret = append(ret, fmt.Sprintf("%s=%s", k, v))
		}
	}
	return
}
