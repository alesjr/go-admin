package example

import (
	"go-admin/go-admin/context"
	"go-admin/go-admin/modules/auth"
	"go-admin/go-admin/modules/db"
	"go-admin/go-admin/modules/service"
)

func (e *Example) initRouter(prefix string, srv service.List) *context.App {

	app := context.NewApp()
	route := app.Group(prefix)
	route.GET("/example", auth.Middleware(db.GetConnection(srv)), e.TestHandler)

	return app
}
