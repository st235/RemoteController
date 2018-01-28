package main

import (
	rtr "remotecontrol/http"
	ctrl "remotecontrol/controllers"
	"fmt"
	"remotecontrol/qr"
	"os"
	"net/http"
)

const PUBLIC_DIR = "./public"
const PUBLIC_PORT = ":8080"

func main() {
	router := rtr.NewInstance()

	if _, err := os.Stat(PUBLIC_DIR); os.IsNotExist(err) {
		fmt.Println("Directory created")
		os.Mkdir(PUBLIC_DIR, os.ModePerm)
	}

	ctrl.
		RegisterKeyboardController(router).
		Handle("/public/", http.StripPrefix("/public/", ctrl.ObtainPublicFilesHandler(PUBLIC_DIR))).
		Serve(PUBLIC_PORT, func(ipAddr string) {
		fmt.Println("Created server at " + ipAddr)
		err := qr.CreateQr(PUBLIC_DIR, ipAddr)

		if err != nil {
			fmt.Print(err)
		}
	})
}
