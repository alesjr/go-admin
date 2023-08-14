package gf

import (
	"net/http"
	"os"

	// add gf adapter
	_ "github.com/alesjr/go-admin/github.com/alesjr/go-admin/adapter/gf"
	// add mysql driver
	_ "github.com/alesjr/go-admin/github.com/alesjr/go-admin/modules/db/drivers/mysql"
	// add postgresql driver
	_ "github.com/alesjr/go-admin/github.com/alesjr/go-admin/modules/db/drivers/postgres"
	// add sqlite driver
	_ "github.com/alesjr/go-admin/github.com/alesjr/go-admin/modules/db/drivers/sqlite"
	// add mssql driver
	_ "github.com/alesjr/go-admin/github.com/alesjr/go-admin/modules/db/drivers/mssql"
	// add adminlte ui theme
	"github.com/alesjr/go-admin/themes/adminlte"

	"github.com/alesjr/go-admin/github.com/alesjr/go-admin/engine"
	"github.com/alesjr/go-admin/github.com/alesjr/go-admin/modules/config"
	"github.com/alesjr/go-admin/github.com/alesjr/go-admin/modules/language"
	"github.com/alesjr/go-admin/github.com/alesjr/go-admin/plugins/admin"
	"github.com/alesjr/go-admin/github.com/alesjr/go-admin/plugins/admin/modules/table"
	"github.com/alesjr/go-admin/github.com/alesjr/go-admin/template"
	"github.com/alesjr/go-admin/github.com/alesjr/go-admin/template/chartjs"
	"github.com/alesjr/go-admin/github.com/alesjr/go-admin/tests/tables"

	"github.com/gogf/gf/frame/g"
)

func internalHandler() http.Handler {
	s := g.Server(8103)

	eng := engine.Default()

	adminPlugin := admin.NewAdmin(tables.Generators).AddDisplayFilterXssJsFilter()

	template.AddComp(chartjs.NewChart())

	adminPlugin.AddGenerator("user", tables.GetUserTable)

	if err := eng.AddConfigFromJSON(os.Args[len(os.Args)-1]).
		AddPlugins(adminPlugin).
		Use(s); err != nil {
		panic(err)
	}

	eng.HTML("GET", "/admin", tables.GetContent)

	return s
}

func NewHandler(dbs config.DatabaseList, gens table.GeneratorList) http.Handler {

	s := g.Server(8103)

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
		AddPlugins(adminPlugin).Use(s); err != nil {
		panic(err)
	}

	template.AddComp(chartjs.NewChart())

	eng.HTML("GET", "/admin", tables.GetContent)

	return s
}
