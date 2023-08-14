package main

import (
	"log"
	"os"
	"os/signal"
	"time"

	_ "github.com/alesjr/go-admin/github.com/alesjr/go-admin/adapter/iris"
	_ "github.com/alesjr/go-admin/github.com/alesjr/go-admin/modules/db/drivers/mysql"
	_ "github.com/alesjr/go-admin/themes/adminlte"

	"github.com/alesjr/go-admin/github.com/alesjr/go-admin/engine"
	"github.com/alesjr/go-admin/github.com/alesjr/go-admin/examples/datamodel"
	"github.com/alesjr/go-admin/github.com/alesjr/go-admin/modules/config"
	"github.com/alesjr/go-admin/github.com/alesjr/go-admin/modules/language"
	"github.com/alesjr/go-admin/github.com/alesjr/go-admin/plugins/example"
	"github.com/alesjr/go-admin/github.com/alesjr/go-admin/template"
	"github.com/alesjr/go-admin/github.com/alesjr/go-admin/template/chartjs"

	"github.com/kataras/iris/v12"
)

func main() {
	app := iris.Default()

	eng := engine.Default()

	cfg := config.Config{
		Env: config.EnvLocal,
		Databases: config.DatabaseList{
			"default": {
				Host:            "127.0.0.1",
				Port:            "3306",
				User:            "root",
				Pwd:             "root",
				Name:            "godmin",
				MaxIdleConns:    50,
				MaxOpenConns:    150,
				ConnMaxLifetime: time.Hour,
				Driver:          config.DriverMysql,
			},
		},
		UrlPrefix: "admin",
		Store: config.Store{
			Path:   "./uploads",
			Prefix: "uploads",
		},
		IndexUrl: "/",
		Debug:    true,
		Language: language.CN,
	}

	template.AddComp(chartjs.NewChart())

	// customize a plugin

	examplePlugin := example.NewExample()

	// load from golang.Plugin
	//
	// examplePlugin := plugins.LoadFromPlugin("../datamodel/example.so")

	// customize the login page
	// example: https://github.com/alesjr/demo.go-admin.cn/blob/master/main.go#L39
	//
	// template.AddComp("login", datamodel.LoginPage)

	// load config from json file
	//
	// eng.AddConfigFromJSON("../datamodel/config.json")

	if err := eng.AddConfig(&cfg).
		AddGenerators(datamodel.Generators).
		AddDisplayFilterXssJsFilter().
		// add generator, first parameter is the url prefix of table when visit.
		// example:
		//
		// "user" => http://localhost:9033/admin/info/user
		//
		AddGenerator("user", datamodel.GetUserTable).
		AddPlugins(examplePlugin).
		Use(app); err != nil {
		panic(err)
	}

	app.HandleDir("/uploads", "./uploads", iris.DirOptions{
		IndexName: "/index.html",
		Gzip:      false,
		ShowList:  false,
	})

	// you can custom your pages like:

	eng.HTML("GET", "/admin", datamodel.GetContent)

	go func() {
		_ = app.Run(iris.Addr(":8099"))
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Print("closing database connection")
	eng.MysqlConnection().Close()
}
