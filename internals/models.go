package internals

type JWTConfig struct {
	Alg    string `json:"alg"`
	Typ    string `json:"typ"`
	Claims Claims `json:"claims"`
	Secret string
}

type Claims struct {
	Id         string `json:"id"` // Unique token ID for the JWT itself
	Sub        string `json:"sub"` // User Id (most likeley a primary key from the database)
	Expiration uint64 `json:"exp"`
}
