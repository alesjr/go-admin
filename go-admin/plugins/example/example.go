package example

import (
	c "alesjr/go-admin/go-admin/modules/config"
	"alesjr/go-admin/go-admin/modules/service"
	"alesjr/go-admin/go-admin/plugins"
)

type Example struct {
	*plugins.Base
}

func NewExample() *Example {
	return &Example{
		Base: &plugins.Base{PlugName: "example"},
	}
}

func (e *Example) InitPlugin(srv service.List) {
	e.InitBase(srv, "example")
	e.App = e.initRouter(c.Prefix(), srv)
}
