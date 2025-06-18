package apiv1

type AddURLBody struct {
	LongURL  string  `json:"longurl"`
	ShortURL *string `json:"shorturl"`
}

type ShortURLS struct {
	Data []AddURLBody
}
