package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gernest/utron"
	"github.com/utronframework/chat/controllers"
)

func main() {

	// Start the MVC App
	app, err := utron.NewMVC()
	if err != nil {
		log.Fatal(err)
	}

	// Register Controllers
	app.AddController(controllers.NewApp)
	app.AddController(controllers.NewWebsocket)
	app.AddController(controllers.NewRefresh)

	// Start the server
	port := fmt.Sprintf(":%d", app.Config.Port)
	app.Log.Info("staring server on port", port)
	log.Fatal(http.ListenAndServe(port, app))
}
