package main
import (
	"fmt"
	"net/http"
	"github.com/byReqz/nanoservices/demo"
)

var Active []Service

type Service struct {
	Name string
	Handle string
	Description string
	Handler Handler
}

type Handler interface {
	Run(w http.ResponseWriter, req *http.Request)
	Register(*[]Service)
}

func main() {
	
}