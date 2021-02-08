package dto

type AnonyURLCountByUser struct {
	Name          string `json:"name" db:"name"`
	Email         string `json:"email" db:"email"`
	CntURLs       int64  `json:"count_urls" db:"count_urls"`
	CntActiveURLs int64  `json:"count_active_urls" db:"count_active_urls"`
}
