package chi

import (
	// add chi adapter
	_ "alesjr/go-admin/go-admin/adapter/chi"
	"alesjr/go-admin/go-admin/modules/config"
	"alesjr/go-admin/go-admin/modules/language"
	"alesjr/go-admin/go-admin/plugins/admin/modules/table"

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

	"net/http"
	"os"

	"alesjr/go-admin/go-admin/engine"
	"alesjr/go-admin/go-admin/plugins/admin"
	"alesjr/go-admin/go-admin/plugins/example"
	"alesjr/go-admin/go-admin/template"
	"alesjr/go-admin/go-admin/template/chartjs"
	"alesjr/go-admin/go-admin/tests/tables"

	"github.com/go-chi/chi"
)

func internalHandler() http.Handler {
	r := chi.NewRouter()

	eng := engine.Default()

	adminPlugin := admin.NewAdmin(tables.Generators)
	adminPlugin.AddGenerator("user", tables.GetUserTable)
	examplePlugin := example.NewExample()
	template.AddComp(chartjs.NewChart())

	if err := eng.AddConfigFromJSON(os.Args[len(os.Args)-1]).
		AddPlugins(adminPlugin, examplePlugin).Use(r); err != nil {
		panic(err)
	}

	eng.HTML("GET", "/admin", tables.GetContent)

	return r
}

func NewHandler(dbs config.DatabaseList, gens table.GeneratorList) http.Handler {
	r := chi.NewRouter()

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
		AddPlugins(adminPlugin).Use(r); err != nil {
		panic(err)
	}

	eng.HTML("GET", "/admin", tables.GetContent)

	return r
}
