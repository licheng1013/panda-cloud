package router

import "strings"

type Router interface {
	// GetAllRouters 获取所有路由信息
	GetAllRouters() []RInfo
}

// RInfo 路由信息
type RInfo struct {
	// 路由名称
	ServiceName string
	// 路由地址
	Path string
}

// Match 匹配路由
func (r *RInfo) Match(path string) bool {
	// 正则匹配
	if strings.HasSuffix(r.Path, "*") {
		p := strings.ReplaceAll(r.Path, "*", "")
		return strings.HasPrefix(path, p)
	}
	return r.Path == path
}
