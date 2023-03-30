package main

import (
	"github.com/licheng1013/panda-cloud/gateway"
	"github.com/licheng1013/panda-cloud/registers"
	"github.com/licheng1013/panda-cloud/router"
)

func main() {
	g := gateway.Gateway{}
	defaultRouter := router.DefaultRouter{}
	g.SetRouter(&defaultRouter)
	g.SetRegister(&registers.Demo{})
	g.Start()
}
