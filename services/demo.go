package services

import (
	"fmt"
	"net/http"
)

func init() {
	var demo Service
	demo.Name = "Nanoservices Demo Service"
	demo.Path = "/demo"
	demo.Description = "simple hello worldy service"
	demo.Handler = demohandler
	demo.Register()
}

func demohandler(w http.ResponseWriter, req *http.Request) {
	Announce(req)
	fmt.Fprintln(w, "nanoservices is running")
}
