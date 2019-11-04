package main

import (
	"context"
	"github.com/go-chi/chi"
	"github.com/piwonskp/feldscher/handlers"
	"net/http"
	"strconv"
)

func main() {
	r := chi.NewRouter()
	r.Route("/api/fetcher", func(r chi.Router) {
		r.Post("/", handlers.CreatePage)
		r.Get("/", handlers.ListPages)
		r.Route("/{id}", func(r chi.Router) {
			r.Use(IntegerId)
			r.Delete("/", handlers.DeletePage)
			r.Get("/history", handlers.FetchedPagesHistory)
		})
	})
	http.ListenAndServe(":8080", r)
}

func IntegerId(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(chi.URLParam(r, "id"))
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		ctx := context.WithValue(r.Context(), "id", id)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
