package models

type UrlDataBaseEntry struct {
	ShortUrl  string `json:"shortUrl"`
	LongUrl   string `json:"longUrl"`
	CreatedBy string `json:"createdBy"`
}
