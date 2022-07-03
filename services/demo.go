package services

import (
	"fmt"
	log "github.com/byReqz/slug"
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
	log.Debug("demoservice has been accessed by", req.RemoteAddr)
	fmt.Fprintln(w, "nanoservices is running")
}
