package main

import (
	rtr "remotecontrol/http"
	keyctrl "remotecontrol/controllers"
)

func main() {
	router := rtr.NewInstance()

	keyctrl.
		RegisterKeyboardController(router).
		Serve(":8080")
}
