package main

import (
	_ "github.com/chenhg5/go-admin/adapter/buffalo"
	"github.com/chenhg5/go-admin/engine"
	"github.com/chenhg5/go-admin/examples/datamodel"
	"github.com/chenhg5/go-admin/modules/config"
	"github.com/chenhg5/go-admin/modules/db"
	"github.com/chenhg5/go-admin/modules/language"
	"github.com/chenhg5/go-admin/plugins/admin"
	"github.com/chenhg5/go-admin/plugins/example"
	"github.com/chenhg5/go-admin/template/adminlte"
	"github.com/chenhg5/go-admin/template/types"
	"github.com/gobuffalo/buffalo"
	"net/http"
)

func main() {
	bu := buffalo.New(buffalo.Options{
		Env:  "test",
		Addr: "127.0.0.1:9033",
	})

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
		DOMAIN: "localhost",
		PREFIX: "admin",
		STORE: config.Store{
			PATH:   "./uploads",
			PREFIX: "uploads",
		},
		LANGUAGE:    language.EN,
		INDEX:       "/",
		DEBUG:       true,
		COLORSCHEME: adminlte.COLORSCHEME_SKIN_BLACK,
	}

	adminPlugin := admin.NewAdmin(datamodel.Generators)

	// add generator, first parameter is the url prefix of table when visit.
	// example:
	//
	// "user" => http://localhost:9033/admin/info/user
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

	if err := eng.AddConfig(cfg).AddPlugins(adminPlugin, examplePlugin).Use(bu); err != nil {
		panic(err)
	}

	bu.ServeFiles("/uploads", http.Dir("./uploads"))

	// you can custom your pages like:

	bu.GET("/admin/custom", func(ctx buffalo.Context) error {
		engine.Content(ctx, func() types.Panel {
			return datamodel.GetContent()
		})
		return nil
	})

	_ = bu.Serve()
}
