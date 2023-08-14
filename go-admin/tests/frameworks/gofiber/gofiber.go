package gofiber

import (
	// add fasthttp adapter
	ada "alesjr/go-admin/go-admin/adapter/gofiber"
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

	"github.com/gofiber/fiber/v2"
	"github.com/valyala/fasthttp"
)

func internalHandler() fasthttp.RequestHandler {
	app := fiber.New(fiber.Config{
		ServerHeader: "Fiber",
	})

	eng := engine.Default()

	adminPlugin := admin.NewAdmin(tables.Generators).AddDisplayFilterXssJsFilter()
	adminPlugin.AddGenerator("user", tables.GetUserTable)

	template.AddComp(chartjs.NewChart())

	if err := eng.AddConfigFromJSON(os.Args[len(os.Args)-1]).
		AddPlugins(adminPlugin).
		Use(app); err != nil {
		panic(err)
	}

	eng.HTML("GET", "/admin", tables.GetContent)

	return app.Handler()
}

func NewHandler(dbs config.DatabaseList, gens table.GeneratorList) fasthttp.RequestHandler {
	app := fiber.New(fiber.Config{
		ServerHeader: "Fiber",
	})

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
		AddAdapter(new(ada.Gofiber)).
		AddGenerators(gens).
		Use(app); err != nil {
		panic(err)
	}

	eng.HTML("GET", "/admin", tables.GetContent)

	return app.Handler()
}
