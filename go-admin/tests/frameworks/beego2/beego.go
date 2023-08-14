package beego

import (
	"net/http"
	"os"

	// add beego adapter
	_ "go-admin/go-admin/adapter/beego2"
	// add mysql driver
	_ "go-admin/go-admin/modules/db/drivers/mysql"
	// add postgresql driver
	_ "go-admin/go-admin/modules/db/drivers/postgres"
	// add sqlite driver
	_ "go-admin/go-admin/modules/db/drivers/sqlite"
	// add mssql driver
	_ "go-admin/go-admin/modules/db/drivers/mssql"

	"go-admin/go-admin/engine"
	"go-admin/go-admin/modules/config"
	"go-admin/go-admin/modules/language"
	"go-admin/go-admin/plugins/admin"
	"go-admin/go-admin/plugins/admin/modules/table"
	"go-admin/go-admin/plugins/example"
	"go-admin/go-admin/template"
	"go-admin/go-admin/template/chartjs"
	"go-admin/go-admin/tests/tables"
	"go-admin/themes/adminlte"

	"github.com/beego/beego/v2/server/web"
)

func internalHandler() http.Handler {

	app := web.NewHttpSever()

	eng := engine.Default()
	adminPlugin := admin.NewAdmin(tables.Generators)
	adminPlugin.AddGenerator("user", tables.GetUserTable)

	examplePlugin := example.NewExample()

	if err := eng.AddConfigFromJSON(os.Args[len(os.Args)-1]).
		AddPlugins(adminPlugin, examplePlugin).Use(app); err != nil {
		panic(err)
	}

	template.AddComp(chartjs.NewChart())

	eng.HTML("GET", "/admin", tables.GetContent)

	app.Cfg.Listen.HTTPAddr = "127.0.0.1"
	app.Cfg.Listen.HTTPPort = 9087

	return app.Handlers
}

func NewHandler(dbs config.DatabaseList, gens table.GeneratorList) http.Handler {

	app := web.NewHttpSever()

	eng := engine.Default()
	adminPlugin := admin.NewAdmin(gens)

	if err := eng.AddConfig(&config.Config{
		Databases: dbs,
		UrlPrefix: "admin",
		Store: config.Store{
			Path:   "./uploads",
			Prefix: "uploads",
		},
		Language:    language.EN,
		IndexUrl:    "/",
		Debug:       true,
		ColorScheme: adminlte.ColorschemeSkinBlack,
	}).
		AddPlugins(adminPlugin).Use(app); err != nil {
		panic(err)
	}

	template.AddComp(chartjs.NewChart())

	eng.HTML("GET", "/admin", tables.GetContent)

	app.Cfg.Listen.HTTPAddr = "127.0.0.1"
	app.Cfg.Listen.HTTPPort = 9087

	return app.Handlers
}
