package services

import (
	"fmt"
	log "github.com/byReqz/slug"
	"net/http"
)

func init() {
	var s Service
	s.Name = "myip"
	s.Path = "/myip"
	s.Description = "echo back the requesting IP address"
	s.Handler = myip
	s.Register()
}

func myip(w http.ResponseWriter, req *http.Request) {
	if req.Header.Get("X-Forwarded-For") != "" {
		log.Debug("service: myip,", "remote_ip:", req.Header.Get("X-Forwarded-For"))
		fmt.Fprintln(w, req.Header.Get("X-Forwarded-For"))
	} else {
		log.Debug("service: myip,", "remote_ip:", req.RemoteAddr)
		fmt.Fprintln(w, req.RemoteAddr)
	}
}
