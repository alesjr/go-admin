package gorilla

import (
	// add gorilla adapter
	_ "go-admin/go-admin/adapter/gorilla"
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

	"github.com/gorilla/mux"
)

func internalHandler() http.Handler {
	app := mux.NewRouter()
	eng := engine.Default()

	examplePlugin := example.NewExample()
	template.AddComp(chartjs.NewChart())

	if err := eng.AddConfigFromJSON(os.Args[len(os.Args)-1]).
		AddPlugins(admin.NewAdmin(tables.Generators).
			AddGenerator("user", tables.GetUserTable), examplePlugin).
		Use(app); err != nil {
		panic(err)
	}

	eng.HTML("GET", "/admin", tables.GetContent)

	return app
}

func NewHandler(dbs config.DatabaseList, gens table.GeneratorList) http.Handler {
	app := mux.NewRouter()
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
		AddPlugins(admin.NewAdmin(gens)).
		Use(app); err != nil {
		panic(err)
	}

	eng.HTML("GET", "/admin", tables.GetContent)

	return app
}
