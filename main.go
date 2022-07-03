package main

import (
	"github.com/byReqz/nanoservices/services"
	log "github.com/byReqz/slug"
	"net/http"
)

func init() {
	log.DefaultLogger.SetLevel(log.InfoLevel)
}

func main() {
	for _, service := range services.Active {
		http.HandleFunc(service.Path, service.Handler)
		log.Info("enabling service " + service.Name + " on " + service.Handle)
	}
	log.Fatal(http.ListenAndServe(":8080", nil))
}
