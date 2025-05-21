package urls

type UrlDataBaseEntry struct {
	short_url  string `json:"shortUrl"`
	long_url   string `json:"longUrl"`
	created_by string `json:"createdBy"`
}
