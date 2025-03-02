package internals

type JWT struct {
	Expiration int64 `json:"exp"`
}

type User struct {
	ID int `json:"id"`
}
