package controllers

import (
	"net/http"

	"github.com/gernest/utron/controller"
)

type App struct {
	controller.BaseController
	Routes []string
}

func (a *App) Index() {
	a.Ctx.Template = "Application/Index"
	a.Ctx.Data["title"] = "Sign in"
	a.HTML(http.StatusOK)
}

func (a *App) Demo() {
	r := a.Ctx.Request()
	r.ParseForm()
	user := r.FormValue("user")
	demo := r.FormValue("demo")
	switch demo {
	case "websocket":
		a.Ctx.Redirect("/websocket/room?user="+user, http.StatusFound)
	}
}

func NewApp() controller.Controller {
	return &App{
		Routes: []string{
			"get;/;Index",
			"get;/demo;Demo",
		},
	}
}
