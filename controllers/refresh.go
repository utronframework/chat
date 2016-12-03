package controllers

import "github.com/gernest/utron/controller"

type Refresh struct {
	controller.BaseController
	Routes []string
}

func (c Refresh) Index() {
}

func (c Refresh) Room() {
}

func (c Refresh) Say(user, message string) {
}

func (c Refresh) Leave(user string) {
}

func NewRefresh() controller.Controller {
	return &Refresh{
		Routes: []string{
			"get;/refresh;Index",
			"get;/refresh/room;Room",
			"post;/refresh/room;Say",
			"post;/refresh/room/leave;Leave",
		},
	}
}
