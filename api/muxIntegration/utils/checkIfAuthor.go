package utils

type UserClaims struct {
	Sub   string `json:"sub"`
	Email string `json:"email"`
	Name  string `json:"name"`
}

var CheckIfAuthor = func(claims *UserClaims, userSub string) bool {
	return claims.Sub == userSub
}
