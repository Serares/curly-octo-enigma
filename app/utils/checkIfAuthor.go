package utils

import "github.com/Serares/curly-octo-enigma/app/services"

var CheckIfAuthor = func(claims *services.UserClaims, userSub string) bool {
	return claims.Sub == userSub
}
