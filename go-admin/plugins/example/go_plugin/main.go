package main

import (
	"alesjr/go-admin/go-admin/context"
	"alesjr/go-admin/go-admin/modules/auth"
	c "alesjr/go-admin/go-admin/modules/config"
	"alesjr/go-admin/go-admin/modules/db"
	"alesjr/go-admin/go-admin/modules/service"
	"alesjr/go-admin/go-admin/plugins"
)

type Example struct {
	*plugins.Base
}

var Plugin = &Example{
	Base: &plugins.Base{PlugName: "example"},
}

func (example *Example) InitPlugin(srv service.List) {
	example.InitBase(srv, "example")
	Plugin.App = example.initRouter(c.Prefix(), srv)
}

func (example *Example) initRouter(prefix string, srv service.List) *context.App {

	app := context.NewApp()
	route := app.Group(prefix)
	route.GET("/example", auth.Middleware(db.GetConnection(srv)), example.TestHandler)

	return app
}

func (example *Example) TestHandler(ctx *context.Context) {

}
