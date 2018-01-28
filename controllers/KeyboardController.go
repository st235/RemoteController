package controllers

import (
	"github.com/go-vgo/robotgo"
	router "remotecontrol/http"
	"net/http"
)

func RegisterKeyboardController(router *router.Router) (*router.Router) {
	return router.
		Add("/keyboard/left", obtainSingleKeyPressHandler("left")).
		Add("/keyboard/right", obtainSingleKeyPressHandler("right"))
}

func obtainSingleKeyPressHandler(key string) (func(w http.ResponseWriter, r *http.Request)) {
	return func(w http.ResponseWriter, r *http.Request) {
		robotgo.KeyTap(key)

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(key))
	}
}


