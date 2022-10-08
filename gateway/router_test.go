package gateway

import "github.com/kataras/iris/v12"

func (m *Manager) RouteTest() {
	m.handler.Get("/test", routerTest)
}

func routerTest(ctx iris.Context) {
	ctx.JSON("test")
}
