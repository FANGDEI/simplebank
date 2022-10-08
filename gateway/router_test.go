package gateway

import "github.com/kataras/iris/v12"

func (m *Manager) RouteTest() {
	m.handler.Get("/test", m.routerTest)
}

func (m *Manager) routerTest(ctx iris.Context) {
	ctx.JSON("test")
}
