package services

import (
	"crypto/rand"
	"fmt"
	log "github.com/byReqz/slug"
	"math/big"
	"net/http"
)

func init() {
	var s Service
	s.Name = "random password generator"
	s.Path = "/pw"
	s.Description = "generate a random 32-char password on query"
	s.Handler = passwordhandler
	s.Register()
}

func passwordhandler(w http.ResponseWriter, req *http.Request) {
	Announce(req)
	var (
		pw      string
		charset string = "abcdedfghijklmnopqrstABCDEFGHIJKLMNOPQRSTUVWXYZ!#$%*_-0123456789"
	)
	for i := 0; i < 32; i++ {
		rng, err := rand.Int(rand.Reader, big.NewInt(int64((len(charset) - 1))))
		if err != nil {
			w.WriteHeader(500)
			pw = "500 internal server error: " + err.Error()
			log.Error("error generating password:", err.Error())
			break
		}
		pw = pw + string(charset[rng.Int64()])
	}
	fmt.Fprintln(w, pw)
}
