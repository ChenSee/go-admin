package main

import (
	_ "github.com/chenhg5/go-admin/adapter/echo"
	"github.com/chenhg5/go-admin/engine"
	"github.com/chenhg5/go-admin/examples/datamodel"
	"github.com/chenhg5/go-admin/modules/config"
	"github.com/chenhg5/go-admin/modules/db"
	"github.com/chenhg5/go-admin/modules/language"
	"github.com/chenhg5/go-admin/plugins/admin"
	"github.com/chenhg5/go-admin/plugins/example"
	"github.com/labstack/echo"
)

func main() {
	e := echo.New()

	eng := engine.Default()

	cfg := config.Config{
		DATABASE: config.DatabaseList{
			"default": {
				HOST:         "127.0.0.1",
				PORT:         "3306",
				USER:         "root",
				PWD:          "root",
				NAME:         "godmin",
				MAX_IDLE_CON: 50,
				MAX_OPEN_CON: 150,
				DRIVER:       db.DriverMysql,
			},
		},
		DOMAIN:   "localhost",
		PREFIX:   "admin",
		INDEX:    "/",
		DEBUG:    true,
		LANGUAGE: language.CN,
	}

	adminPlugin := admin.NewAdmin(datamodel.Generators)

	// add generator, first parameter is the url prefix of table when visit.
	// example:
	//
	// "user" => http://localhost:1323/admin/info/user
	//
	adminPlugin.AddGenerator("user", datamodel.GetUserTable)

	// customize a plugin

	examplePlugin := example.NewExample()

	// load from golang.Plugin
	//
	// examplePlugin := plugins.LoadFromPlugin("../datamodel/example.so")

	// customize the login page
	// example: https://github.com/chenhg5/go-admin/blob/master/demo/main.go#L30
	//
	// template.AddComp("login", datamodel.LoginPage)

	// load config from json file
	//
	// eng.AddConfigFromJson("../datamodel/config.json")

	if err := eng.AddConfig(cfg).AddPlugins(adminPlugin, examplePlugin).Use(e); err != nil {
		panic(err)
	}

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}
