package urls

import (
	"encoding/json"
	"log"
	"net/http"
	"os/exec"

	"github.com/aidenfine/tny/tny-src/models"
	"github.com/jmoiron/sqlx"
)

// TODO: create handler for these funcs
func CreateShortUrl(w http.ResponseWriter, r *http.Request, db *sqlx.DB) {
	var newShortUrl models.UrlDataBaseEntry
	if err := json.NewDecoder(r.Body).Decode(&newShortUrl); err != nil {
		log.Printf("decode error create url")
		http.Error(w, "invalid request payload", http.StatusAccepted)
		return
	}
	uuid := generateUUID()
	query := `INSERT INTO urls (short_url, long_url, created_by) VALUES ($1, $2, $3) RETURNING short_url`
	query_err := db.QueryRow(query, uuid, newShortUrl.LongUrl, newShortUrl.CreatedBy).Scan(&newShortUrl.ShortUrl)
	if query_err != nil {
		log.Printf("failed creating short url", query_err)
		http.Error(w, "Failed to create short url", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newShortUrl)
	// generate unique token and insert into data base
	// INSERT WILL RETURN ERR if item with same PK already exists
}

func generateUUID() string {
	uuid, err := exec.Command("uuidgen").Output()
	if err != nil {
		log.Panic(err)
	}
	return string(uuid)
}
