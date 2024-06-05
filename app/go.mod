module github.com/Serares/curly-octo-enigma/app

go 1.22.0

require (
	github.com/Serares/curly-octo-enigma/domain/repo v0.0.0-00010101000000-000000000000
	github.com/a-h/templ v0.2.707
	github.com/akrylysov/algnhsa v1.1.0
	github.com/coreos/go-oidc v2.2.1+incompatible
	github.com/joho/godotenv v1.5.1
	golang.org/x/oauth2 v0.20.0
)

require (
	github.com/aws/aws-lambda-go v1.43.0 // indirect
	github.com/pquerna/cachecontrol v0.2.0 // indirect
	github.com/stretchr/testify v1.9.0 // indirect
	golang.org/x/crypto v0.23.0 // indirect
	gopkg.in/square/go-jose.v2 v2.6.0 // indirect
)

replace github.com/Serares/curly-octo-enigma/domain/repo => ../domain/repo

replace github.com/Serares/curly-octo-enigma/shared => ../shared
