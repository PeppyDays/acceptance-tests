package main

import (
	"net/http"

	"github.com/quii/go-specs-greet/internal/interface/controllers"
)

func main() {
	handler := controllers.GetHandler()
	_ = http.ListenAndServe(":8080", handler)
}
