package keys

import (
	"fmt"
	"strings"
)

var prefix = "restdoc/"

// Redis Key
var (
	ZSearchKeywords = prefix + "ZSearchKeywords" // 搜索关键字统计
)

// Join 组合Key
func Join(key string, args ...interface{}) string {
	a := []string{key}
	for _, arg := range args {
		a = append(a, fmt.Sprint(arg))
	}
	return strings.Join(a, ":")
}
