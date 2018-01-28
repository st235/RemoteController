package http

import (
	"net/http"
	"remotecontrol/ip"
	"fmt"
)

type Router struct {}

func NewInstance() (*Router) {
	return new(Router)
}

func (r *Router) Add(path string, handler func(http.ResponseWriter, *http.Request)) (*Router) {
	http.HandleFunc(path, handler)
	return r
}

func (r *Router) Serve(port string) (*Router) {
	ipaddr := ip.GetMachineIp(port)
	fmt.Println("serve: " + ipaddr)

	http.ListenAndServe(port, nil)
	return r
}
