package http

import (
	"net/http"
	"remotecontrol/ip"
)

type Router struct {}

func NewInstance() (*Router) {
	return new(Router)
}

func (r *Router) Add(path string, handler func(http.ResponseWriter, *http.Request)) (*Router) {
	http.HandleFunc(path, handler)
	return r
}

func (r *Router) Handle(path string, handler http.Handler) (*Router) {
	http.Handle(path, handler)
	return r
}

func (r *Router) Serve(port string, callback func(ipAddr string)) (*Router) {
	ipaddr := ip.GetMachineIp(port)
	callback(ipaddr)

	http.ListenAndServe(port, nil)
	return r
}
