package fasthttp

import (
	// add fasthttp adapter
	ada "alesjr/go-admin/go-admin/adapter/fasthttp"
	// add mysql driver
	_ "alesjr/go-admin/go-admin/modules/db/drivers/mysql"
	// add postgresql driver
	_ "alesjr/go-admin/go-admin/modules/db/drivers/postgres"
	// add sqlite driver
	_ "alesjr/go-admin/go-admin/modules/db/drivers/sqlite"
	// add mssql driver
	_ "alesjr/go-admin/go-admin/modules/db/drivers/mssql"
	// add adminlte ui theme
	"alesjr/go-admin/themes/adminlte"

	"os"

	"alesjr/go-admin/go-admin/engine"
	"alesjr/go-admin/go-admin/modules/config"
	"alesjr/go-admin/go-admin/modules/language"
	"alesjr/go-admin/go-admin/plugins/admin"
	"alesjr/go-admin/go-admin/plugins/admin/modules/table"
	"alesjr/go-admin/go-admin/template"
	"alesjr/go-admin/go-admin/template/chartjs"
	"alesjr/go-admin/go-admin/tests/tables"

	"github.com/buaazp/fasthttprouter"
	"github.com/valyala/fasthttp"
)

func internalHandler() fasthttp.RequestHandler {
	router := fasthttprouter.New()

	eng := engine.Default()

	adminPlugin := admin.NewAdmin(tables.Generators).AddDisplayFilterXssJsFilter()
	adminPlugin.AddGenerator("user", tables.GetUserTable)

	template.AddComp(chartjs.NewChart())

	if err := eng.AddConfigFromJSON(os.Args[len(os.Args)-1]).
		AddPlugins(adminPlugin).
		Use(router); err != nil {
		panic(err)
	}

	eng.HTML("GET", "/admin", tables.GetContent)

	return func(ctx *fasthttp.RequestCtx) {
		router.Handler(ctx)
	}
}

func NewHandler(dbs config.DatabaseList, gens table.GeneratorList) fasthttp.RequestHandler {
	router := fasthttprouter.New()

	eng := engine.Default()

	template.AddComp(chartjs.NewChart())

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
		AddAdapter(new(ada.Fasthttp)).
		AddGenerators(gens).
		Use(router); err != nil {
		panic(err)
	}

	eng.HTML("GET", "/admin", tables.GetContent)

	return func(ctx *fasthttp.RequestCtx) {
		router.Handler(ctx)
	}
}
