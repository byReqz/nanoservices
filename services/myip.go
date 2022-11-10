package services

import (
	"fmt"
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
	Announce(req)
	if req.Header.Get("X-Forwarded-For") != "" {
		fmt.Fprintln(w, req.Header.Get("X-Forwarded-For"))
	} else {
		fmt.Fprintln(w, req.RemoteAddr)
	}
}
