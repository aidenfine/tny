package models

type UrlDataBaseEntry struct {
	ShortUrl  string  `json:"shortUrl" db:"short_url"`
	LongUrl   string  `json:"longUrl" db:"long_url"`
	Domain    *string `json:"domain" db:"domain"`
	CreatedBy string  `json:"createdBy" db:"created_by"`
	ExpiresAt string  `json:"expiresAt" db:"expires_at"`
}

type UrlsDataBaseItem struct {
	ShortUrl  string  `json:"shortUrl" db:"short_url"`
	LongUrl   string  `json:"longUrl" db:"long_url"`
	Domain    *string `json:"domain" db:"domain"`
	CreatedBy string  `json:"createdBy" db:"created_by"`
	CreatedAt string  `json:"createdAt" db:"created_at"`
	ExpiresAt string  `json:"expiresAt" db:"expires_at"`
}
