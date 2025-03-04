package internals

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/PS-Wizard/JWT/internals/encryptions"
)

type JWTConfig struct {
	Alg string `json:"alg"`
	Typ string `json:"typ"`
}

func (j *JWTConfig) GenerateHeader() string {
	data, err := json.Marshal(j)
	if err != nil {
		log.Fatalf("Couldnt Marshol Json: %v", err)
	}
	fmt.Println(string(data))
	return encryptions.EncodeBase64(string(data))
}
