package services

import "net/http"

var Active []*Service // list of registered services

type Service struct {
	Name        string                                   // name of the service
	Path        string                                   // url path of the service
	Description string                                   // short description of the service
	Handler     func(http.ResponseWriter, *http.Request) // the function that is run when the service is called
}

// Register registers a service to be executed
func (s *Service) Register() {
	Active = append(Active, s)
}
