package guard

import (
	"go-admin/go-admin/context"
	"go-admin/go-admin/modules/db"
	"go-admin/go-admin/modules/errors"
	"go-admin/go-admin/modules/service"
	"go-admin/go-admin/plugins/admin/modules/constant"
	"go-admin/go-admin/plugins/admin/modules/response"
	"go-admin/go-admin/plugins/admin/modules/table"
	"go-admin/go-admin/template"
	"go-admin/go-admin/template/types"
)

type Guard struct {
	services  service.List
	conn      db.Connection
	tableList table.GeneratorList
	navBtns   *types.Buttons
}

func New(s service.List, c db.Connection, t table.GeneratorList, b *types.Buttons) *Guard {
	return &Guard{
		services:  s,
		conn:      c,
		tableList: t,
		navBtns:   b,
	}
}

func (g *Guard) table(ctx *context.Context) (table.Table, string) {
	prefix := ctx.Query(constant.PrefixKey)
	return g.tableList[prefix](ctx), prefix
}

func (g *Guard) CheckPrefix(ctx *context.Context) {

	prefix := ctx.Query(constant.PrefixKey)

	if _, ok := g.tableList[prefix]; !ok {
		if ctx.Headers(constant.PjaxHeader) == "" && ctx.Method() != "GET" {
			response.BadRequest(ctx, errors.Msg)
		} else {
			response.Alert(ctx, errors.Msg, errors.Msg, "table model not found", g.conn, g.navBtns,
				template.Missing404Page)
		}
		ctx.Abort()
		return
	}

	ctx.Next()
}

const (
	editFormParamKey    = "edit_form_param"
	deleteParamKey      = "delete_param"
	exportParamKey      = "export_param"
	serverLoginParamKey = "server_login_param"
	deleteMenuParamKey  = "delete_menu_param"
	editMenuParamKey    = "edit_menu_param"
	newMenuParamKey     = "new_menu_param"
	newFormParamKey     = "new_form_param"
	updateParamKey      = "update_param"
	showFormParamKey    = "show_form_param"
	showNewFormParam    = "show_new_form_param"
)
