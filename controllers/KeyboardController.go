package controllers

import (
    "github.com/go-vgo/robotgo"
    router "remotecontrol/http"
    "net/http"
    "fmt"
)

var keys = []string{"left", "right", "up", "down", "space"}

func RegisterKeyboardController(router *router.Router) (*router.Router) {
    for _, key := range keys {
        router.
            Add(fmt.Sprintf("/keyboard/%s", key), obtainSingleKeyPressHandler(key))
    }
    return router
}

func obtainSingleKeyPressHandler(key string) (func(w http.ResponseWriter, r *http.Request)) {
    return func(w http.ResponseWriter, r *http.Request) {
        robotgo.KeyTap(key)

        w.WriteHeader(http.StatusOK)
        w.Write([]byte(key))
    }
}


