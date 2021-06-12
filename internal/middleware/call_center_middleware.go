package middleware

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
	"strings"
)

func CheckParams(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		vars := mux.Vars(r)
		if vars["date"] == "" {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, "La fecha es obligatoria")
			return
		}
		if len(strings.Split(vars["date"], "-")) != 3 {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, "Formato invalido de fecha, el formato es YYYY-mm=DD, 2020-12-01")
			return
		}
		days := r.URL.Query().Get("dias")
		if days == "" {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, "El numero de dias es obligatorio")
			return
		}
		number, err := strconv.Atoi(days)
		if err != nil || 0 > number {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, "formato de parametro days invalido")
			return
		}
		if 0 == number {
			fmt.Fprintf(w, "No existen datos que mostrar")
			return
		}
		next(w, r)
	}

}
