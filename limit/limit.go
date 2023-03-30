package limit

// Limit 限制器，暂时废弃
type Limit interface {
	// Proxy ctx.AbortWithStatusJSON(http.StatusOK, data) 中断处理
}

type RequestInfo struct {
	// 请求路径
	Path string
	// 请求方法
	Method string
}
