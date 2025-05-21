package router

import (
	"log"
	"net/http"

	"github.com/aidenfine/tny/tny-src/services/urls"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	"github.com/rs/cors"
)

func StartRouter(db *sqlx.DB) error {
	r := mux.NewRouter()

	// setup cors
	c := cors.New(cors.Options{
		AllowedOrigins: []string{
			"http://localhost:8080",
		}, AllowedMethods: []string{
			"GET", "POST", "PUT", "DELETE", "OPTIONS",
		}, AllowedHeaders: []string{
			"Content-Type", "Authorization",
		},
		AllowCredentials: true,
	})

	registerRoutes(r, db)

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Server is running"))
	}).Methods("GET")
	handler := c.Handler(r)
	log.Println("Server running on port 8080...")

	return http.ListenAndServe(":8080", handler)
}
func registerRoutes(r *mux.Router, db *sqlx.DB) {
	urls.RegisterUrlsRoutes(r, db)
}
