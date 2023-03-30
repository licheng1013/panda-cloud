package gateway

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/licheng1013/panda-cloud/limit"
	"github.com/licheng1013/panda-cloud/registers"
	"github.com/licheng1013/panda-cloud/router"
	"io"
	"net/http"
)

type Gateway struct {
	// 注册中心
	register registers.Register
	// 路由
	router router.Router
	// 限流器
	flowLimit limit.FlowLimit
	// 熔断器
	fuse limit.Limit
	// 降级器
	overload limit.Limit
}

func (g *Gateway) SetRegister(register registers.Register) {
	g.register = register
}

func (g *Gateway) SetRouter(router router.Router) {
	g.router = router
}

func (g *Gateway) Start() {
	r := gin.Default()
	r.Use(g.Proxy)
	_ = r.Run()
}

func (g *Gateway) Proxy(ctx *gin.Context) {
	data := gin.H{"message": "Hello, world!"}
	// 获取请求的url
	url := ctx.Request.URL
	fmt.Println("请求路径 -> ", url)

	// 限流处理 1 - 判断是否超过限流
	trigger, message := g.flowLimit.Control(limit.RequestInfo{Path: url.Path, Method: ctx.Request.Method})
	if trigger {
		ctx.AbortWithStatusJSON(http.StatusOK, message)
		return
	}
	for _, info := range g.router.GetAllRouters() {
		if info.Match(url.Path) { // 匹配成功路由
			rInfo, _ := g.register.GetIp(info.ServiceName)
			request, _ := http.NewRequestWithContext(context.Background(), ctx.Request.Method, "http://"+rInfo.Addr()+url.Path, ctx.Request.Body)
			// 发送请求
			response, _ := http.DefaultClient.Do(request)
			bytes, _ := io.ReadAll(response.Body)

			// 熔断处理 2 - 判断是否触发熔断,则不触发请求
			//g.fuse.Proxy(ctx.Copy())
			// 降级处理 3 - 判断熔断关闭,判断是否需要降级,则返回降级数据
			//g.overload.Proxy(ctx.Copy())
			// 熔断是直接拒绝请求,降级是继续请求,但是返回降级数据

			var d gin.H
			_ = json.Unmarshal(bytes, &d)
			ctx.AbortWithStatusJSON(http.StatusOK, d)
			return
		}
	}
	ctx.AbortWithStatusJSON(http.StatusOK, data)
	return
}
