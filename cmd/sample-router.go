package cmd

import (
	"fmt"
	"github.com/spf13/viper"
	"gitlab.com/M.darvish/funtory/api/routes"
	"gitlab.com/M.darvish/funtory/configs"
	"log"
	"net/http"
)

func runServer() {
	// read all config files
	configs.Setup()
	appPort := viper.GetString("app.port")
	container := NewDig()

	// keep application up
	go func() {

		if err := container.Invoke(func(router *routes.APIRouteHandler) {
			invokeFunc(appPort, router)
		}); err != nil {
			panic(err)
		}
	}()

	fmt.Printf("Server is listening on port %s...\n", appPort)
	// blocks app indefinitely and never exits.
	select {}
}

func invokeFunc(appPort string, router *routes.APIRouteHandler) {
	if err := http.ListenAndServe(":"+appPort, router.Router); err != nil {
		log.Fatalf("Error starting server: %s", err)
	}
}
