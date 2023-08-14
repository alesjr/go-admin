package gear

import (
	// add gin adapter
	ada "alesjr/go-admin/go-admin/adapter/gear"

	"github.com/teambition/gear"

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
	"alesjr/go-admin/go-admin/modules/config"
	"alesjr/go-admin/go-admin/modules/language"
	"alesjr/go-admin/go-admin/plugins/admin/modules/table"
	"alesjr/go-admin/go-admin/template"
	"alesjr/go-admin/go-admin/template/chartjs"
	"alesjr/go-admin/go-admin/tests/tables"
)

func internalHandler() http.Handler {
	app := gear.New()

	eng := engine.Default()

	template.AddComp(chartjs.NewChart())

	if err := eng.AddConfigFromJSON(os.Args[len(os.Args)-1]).
		AddGenerators(tables.Generators).
		AddGenerator("user", tables.GetUserTable).
		Use(app); err != nil {
		panic(err)
	}

	eng.HTML("GET", "/admin", tables.GetContent)

	return app
}

func NewHandler(dbs config.DatabaseList, gens table.GeneratorList) http.Handler {
	app := gear.New()

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
		AddAdapter(new(ada.Gear)).
		AddGenerators(gens).
		Use(app); err != nil {
		panic(err)
	}

	eng.HTML("GET", "/admin", tables.GetContent)

	return app
}
