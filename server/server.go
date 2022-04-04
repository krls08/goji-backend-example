// Package server does things :D
package server

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/carles/repo/authz"
	goji "goji.io/v3"
	"goji.io/v3/pat"
)

// RunServer is yada yada...
func RunServer() {
	mux := goji.NewMux()
	mux.HandleFunc(pat.Get("/hello/:name"), jsonMw(hello))
	mux.HandleFunc(pat.Get("/bye/:name"), jsonMw(bye))

	http.ListenAndServe("localhost:8000", mux)
}

func hello(w http.ResponseWriter, r *http.Request) {
	name := pat.Param(r, "name")

	user := authz.FetchUser(name)

	json.NewEncoder(w).Encode(user)
}

func bye(w http.ResponseWriter, r *http.Request) {
	name := pat.Param(r, "name")
	fmt.Fprintf(w, "Bye, %s!", name)
}

func jsonMw(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("content-type", "application/json")
		w.WriteHeader(http.StatusOK)

		next(w, r)
	}
}
