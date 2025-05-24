package urls

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
)

func RegisterUrlsRoutes(r *mux.Router, db *sqlx.DB) {
	urlsRouterInternal := r.PathPrefix("/v1/urls").Subrouter()
	urlsRouterExternal := r.PathPrefix("").Subrouter()

	urlsRouterInternal.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		CreateShortUrl(w, r, db)
	}).Methods("POST")
	urlsRouterExternal.HandleFunc("/{shortUrl}", func(w http.ResponseWriter, r *http.Request) {
		RedirectToLongUrl(w, r, db)
	}).Methods("GET")

}
