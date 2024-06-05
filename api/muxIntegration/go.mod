module github.com/Serares/curly-octo-enigma/api/muxIntegration

go 1.22.0

replace github.com/Serares/curly-octo-enigma/domain/repo => ../../domain/repo

require (
	github.com/Serares/curly-octo-enigma/domain/repo v0.0.0-00010101000000-000000000000
	github.com/akrylysov/algnhsa v1.1.0
	github.com/coreos/go-oidc v2.2.1+incompatible
	github.com/google/uuid v1.6.0
	github.com/joho/godotenv v1.5.1
	golang.org/x/oauth2 v0.21.0
)

require (
	github.com/antlr4-go/antlr/v4 v4.13.0 // indirect
	github.com/aws/aws-lambda-go v1.43.0 // indirect
	github.com/dustin/go-humanize v1.0.1 // indirect
	github.com/hashicorp/golang-lru/v2 v2.0.7 // indirect
	github.com/libsql/sqlite-antlr4-parser v0.0.0-20240327125255-dbf53b6cbf06 // indirect
	github.com/mattn/go-isatty v0.0.20 // indirect
	github.com/ncruces/go-strftime v0.1.9 // indirect
	github.com/remyoudompheng/bigfft v0.0.0-20230129092748-24d4a6f8daec // indirect
	github.com/tursodatabase/libsql-client-go v0.0.0-20240416075003-747366ff79c4 // indirect
	golang.org/x/exp v0.0.0-20240325151524-a685a6edb6d8 // indirect
	golang.org/x/sys v0.21.0 // indirect
	modernc.org/gc/v3 v3.0.0-20240107210532-573471604cb6 // indirect
	modernc.org/libc v1.50.9 // indirect
	modernc.org/mathutil v1.6.0 // indirect
	modernc.org/memory v1.8.0 // indirect
	modernc.org/sqlite v1.30.0 // indirect
	modernc.org/strutil v1.2.0 // indirect
	modernc.org/token v1.1.0 // indirect
	nhooyr.io/websocket v1.8.10 // indirect
)

require (
	github.com/Serares/curly-octo-enigma/shared v0.0.0 // indirect
	github.com/Serares/curly-octo-enigma/utils v0.0.0
	github.com/pquerna/cachecontrol v0.2.0 // indirect
	golang.org/x/crypto v0.24.0 // indirect
	gopkg.in/square/go-jose.v2 v2.6.0 // indirect
)

replace github.com/Serares/curly-octo-enigma/utils => ../../utils

replace github.com/Serares/curly-octo-enigma/shared => ../../shared
