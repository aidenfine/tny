package urls

import (
	"context"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v5/pgxpool"
)

func RegisterUrlsRoutes(ctx context.Context, r *mux.Router, db *pgxpool.Pool) {
	urlsRouterInternal := r.PathPrefix("/v1/urls").Subrouter()
	urlsRouterExternal := r.PathPrefix("").Subrouter()

	urlsRouterInternal.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		CreateShortUrl(ctx, w, r, db)
	}).Methods("POST")
	urlsRouterExternal.HandleFunc("/{shortUrl}", func(w http.ResponseWriter, r *http.Request) {
		RedirectToLongUrl(ctx, w, r, db)
	}).Methods("GET")

}
