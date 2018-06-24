package main

import (
	server "remotecontrol/http"
	ctrl "remotecontrol/controllers"
	"fmt"
	"remotecontrol/qr"
	"os"
	"net/http"
)

const PUBLIC_DIR = "./public"
const PUBLIC_PORT = ":8080"

func main() {
	router := server.NewInstance()

	if _, err := os.Stat(PUBLIC_DIR); os.IsNotExist(err) {
		fmt.Println("Public directory created")
		os.Mkdir(PUBLIC_DIR, os.ModePerm)
	}

	ctrl.RegisterKeyboardController(router).
		Handle("/public/", http.StripPrefix("/public/", ctrl.ObtainPublicFilesHandler(PUBLIC_DIR))).
		Serve(PUBLIC_PORT, func(ipAddr string) {

		fmt.Printf("Your server is started\nFor exit press ctrl+x button\nServer started at %s\n", ipAddr)
		fmt.Printf("You can find controls at:\n - %spublic\n - %spublic/connect.png\n", ipAddr, ipAddr)
		
		err := qr.CreateQr(PUBLIC_DIR, ipAddr)

		if err != nil {
			fmt.Print(err)
		}
	})
}
