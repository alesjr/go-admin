package main

import (
	"github.com/alesjr/go-admin/github.com/alesjr/go-admin/context"
	"github.com/alesjr/go-admin/github.com/alesjr/go-admin/modules/auth"
	c "github.com/alesjr/go-admin/github.com/alesjr/go-admin/modules/config"
	"github.com/alesjr/go-admin/github.com/alesjr/go-admin/modules/db"
	"github.com/alesjr/go-admin/github.com/alesjr/go-admin/modules/service"
	"github.com/alesjr/go-admin/github.com/alesjr/go-admin/plugins"
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
