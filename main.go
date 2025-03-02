package main

import (
	"github.com/PS-Wizard/JWT/internals"
)

func main() {
	internals.Base64_Encode("something")
	user := internals.User{
		ID: 0,
	}
	user.GenerateJWT()
}
