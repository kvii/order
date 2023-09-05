package main

import (
	"fmt"
	"net/http"

	"github.com/kvii/order/internal/order"
	"github.com/kvii/pkg/database"
)

func main() {
	s := order.Service{
		DB: &database.DB{},
	}
	http.HandleFunc("/create", func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		name := r.URL.Query().Get("name")

		msg, err := s.Create(ctx, name)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		} else {
			fmt.Fprintln(w, msg)
		}
	})
	http.ListenAndServe(":9091", nil)
}
