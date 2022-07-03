package main

import (
	"fmt"
	"github.com/byReqz/nanoservices/services"
	log "github.com/byReqz/slug"
	flag "github.com/spf13/pflag"
	"net/http"
)

var listenport int

func init() {
	verbose := flag.BoolP("verbose", "v", false, "print debug messages")
	quiet := flag.BoolP("quiet", "q", false, "only print errors")
	quietquiet := flag.BoolP("silent", "s", false, "dont print anything")
	flag.IntVarP(&listenport, "port", "p", 8080, "port to listen on")
	flag.Parse()
	if *verbose {
		log.DefaultLogger.SetLevel(log.DebugLevel)
		log.Debug("logger set to debug level")
	} else if *quiet {
		log.DefaultLogger.SetLevel(log.WarningLevel)
	} else if *quietquiet {
		log.DefaultLogger.SetLevel(log.Disabled)
	}

}

func main() {
	for _, service := range services.Active {
		http.HandleFunc(service.Path, service.Handler)
		log.Info("enabling service " + service.Name + " on " + service.Path)
	}
	log.Info("listening on port", listenport)
	log.Fatal(http.ListenAndServe(fmt.Sprint(":", listenport), nil))
}
