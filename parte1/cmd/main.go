package main

import (
	"github.com/gorilla/mux"
	"github.com/memeoAmazonas/imuko-test-golang/parte1/internal/call_center"
	"github.com/memeoAmazonas/imuko-test-golang/parte1/internal/middleware"
	"net/http"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/resumen/{date}", middleware.CheckParams(call_center.Service))
	http.ListenAndServe(":8080", r)
}
