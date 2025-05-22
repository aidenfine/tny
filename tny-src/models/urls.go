package models

type UrlDataBaseEntry struct {
	ShortUrl  string  `json:"shortUrl"`
	LongUrl   string  `json:"longUrl"`
	Domain    *string `json:"domain"`
	CreatedBy string  `json:"createdBy"`
}

type UrlsDataBaseItem struct {
	ShortUrl  string  `json:"shortUrl"`
	LongUrl   string  `json:"longUrl"`
	Domain    *string `json:"domain"`
	CreatedBy string  `json:"createdBy"`
	CreatedAt string  `json:"createdAt"`
}
