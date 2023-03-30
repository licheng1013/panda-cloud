package router

type DefaultRouter struct {
}

// GetAllRouters 默认路由配置
func (d DefaultRouter) GetAllRouters() []RInfo {
	return []RInfo{
		{ServiceName: "user-service", Path: "/user/*"},
		{ServiceName: "goods-service", Path: "/goods/*"},
		{ServiceName: "order-service", Path: "/order/*"},
	}
}
