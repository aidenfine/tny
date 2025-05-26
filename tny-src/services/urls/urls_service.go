package urls

import (
	"context"
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"time"

	"github.com/aidenfine/tny/tny-src/models"
	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v5/pgxpool"
)

// TODO: create handler for these funcs
func CreateShortUrl(ctx context.Context, w http.ResponseWriter, r *http.Request, db *pgxpool.Pool) {
	var newShortUrl models.UrlDataBaseEntry
	if err := json.NewDecoder(r.Body).Decode(&newShortUrl); err != nil {
		http.Error(w, "invalid request payload", http.StatusAccepted)
		return
	}
	var shortUrl string
	var doesExist bool

	for {
		shortUrl = generateShortUrl(6)
		doesExist = doesShortUrlExist(ctx, shortUrl, db)
		if !doesExist {
			break
		}
	}

	query := `INSERT INTO urls (short_url, long_url, created_by) VALUES ($1, $2, $3) RETURNING short_url`
	query_err := db.QueryRow(ctx, query, shortUrl, newShortUrl.LongUrl, newShortUrl.CreatedBy).Scan(&newShortUrl.ShortUrl)
	if query_err != nil {
		http.Error(w, "Failed to create short url", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newShortUrl)
	// generate unique token and insert into data base
	// INSERT WILL RETURN ERR if item with same PK already exists
}

func getShortUrlItem(ctx context.Context, short_url string, db *pgxpool.Pool) (models.UrlsDataBaseItem, error) {
	var databaseItem models.UrlsDataBaseItem

	query := `SELECT short_url, long_url, domain, created_by, created_at FROM urls WHERE short_url = $1`
	fmt.Println(query, "query")
	fmt.Println(short_url, "short url ")
	err := db.QueryRow(ctx, query, short_url).Scan(&databaseItem)
	if err != nil {
		fmt.Println("DB error:", err)
	}
	fmt.Println(databaseItem, "getShortUrlItemQueryResponse")
	return databaseItem, err
}

func RedirectToLongUrl(ctx context.Context, w http.ResponseWriter, r *http.Request, db *pgxpool.Pool) {
	vars := mux.Vars(r)
	shortUrl := vars["shortUrl"]

	if shortUrl == "" {
		http.Error(w, "Short URL is missing", http.StatusBadRequest)
		return
	}
	databaseItem, err := getShortUrlItem(ctx, shortUrl, db)
	fmt.Println(databaseItem, "databae item")
	if err != nil {
		http.Error(w, "An issue occured when getting long url", http.StatusBadRequest)
		return
	}
	http.Redirect(w, r, databaseItem.LongUrl, http.StatusFound)

}

func generateShortUrl(length int) string {
	rand.New(rand.NewSource(time.Now().UnixNano()))
	characters := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ123456790")
	res := make([]rune, length)
	for i := range res {
		res[i] = characters[rand.Intn(len(characters))]
	}
	return string(res)
}

// Returns True if exists
func doesShortUrlExist(ctx context.Context, shortUrl string, db *pgxpool.Pool) bool {
	res := false
	query := `SELECT EXISTS (SELECT 1 FROM urls WHERE short_url = $1)`
	err := db.QueryRow(ctx, query, shortUrl).Scan(&res)
	if err != nil {
		fmt.Println("Issue checking for short url")
	}
	return res
}
