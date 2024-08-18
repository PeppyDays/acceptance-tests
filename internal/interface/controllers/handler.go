package controllers

import (
	"fmt"
	"net/http"

	"github.com/quii/go-specs-greet/internal/domain/interactions"
)

var NameQueryParameterKey = "name"

func GetHandler() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/greet", handleGreet)
	mux.HandleFunc("/curse", handleCurse)
	return mux
}

func handleGreet(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get(NameQueryParameterKey)
	fmt.Fprint(w, interactions.Greet(name))
}

func handleCurse(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get(NameQueryParameterKey)
	fmt.Fprint(w, interactions.Curse(name))
}
