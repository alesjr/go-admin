package echo

import (
	// add echo adapter
	_ "go-admin/go-admin/adapter/echo"
	"go-admin/go-admin/modules/config"
	"go-admin/go-admin/modules/language"
	"go-admin/go-admin/plugins/admin/modules/table"

	// add mysql driver
	_ "go-admin/go-admin/modules/db/drivers/mysql"
	// add postgresql driver
	_ "go-admin/go-admin/modules/db/drivers/postgres"
	// add sqlite driver
	_ "go-admin/go-admin/modules/db/drivers/sqlite"
	// add mssql driver
	_ "go-admin/go-admin/modules/db/drivers/mssql"
	// add adminlte ui theme
	"go-admin/themes/adminlte"

	"net/http"
	"os"

	"go-admin/go-admin/engine"
	"go-admin/go-admin/plugins/admin"
	"go-admin/go-admin/plugins/example"
	"go-admin/go-admin/template"
	"go-admin/go-admin/template/chartjs"
	"go-admin/go-admin/tests/tables"

	"github.com/labstack/echo/v4"
)

func internalHandler() http.Handler {
	e := echo.New()

	eng := engine.Default()

	adminPlugin := admin.NewAdmin(tables.Generators)
	adminPlugin.AddGenerator("user", tables.GetUserTable)
	template.AddComp(chartjs.NewChart())

	examplePlugin := example.NewExample()

	if err := eng.AddConfigFromJSON(os.Args[len(os.Args)-1]).
		AddPlugins(adminPlugin, examplePlugin).Use(e); err != nil {
		panic(err)
	}

	eng.HTML("GET", "/admin", tables.GetContent)

	return e
}

func NewHandler(dbs config.DatabaseList, gens table.GeneratorList) http.Handler {
	e := echo.New()

	eng := engine.Default()

	adminPlugin := admin.NewAdmin(gens)

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
		AddPlugins(adminPlugin).Use(e); err != nil {
		panic(err)
	}

	eng.HTML("GET", "/admin", tables.GetContent)

	return e
}
