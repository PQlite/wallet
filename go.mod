module wallet

go 1.24.5

require (
	github.com/PQlite/core v0.0.0-20250801132404-2c57b31b44c5
	github.com/PQlite/crypto v0.0.5-0.20250720122300-266869b5900f
	github.com/spf13/cobra v1.9.1
	github.com/stretchr/testify v1.10.0
	golang.org/x/crypto v0.40.0
	golang.org/x/term v0.33.0
)

require (
	github.com/cloudflare/circl v1.6.1 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/inconshreveable/mousetrap v1.1.0 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/polydawn/refmt v0.89.0 // indirect
	github.com/spf13/pflag v1.0.6 // indirect
	golang.org/x/sys v0.34.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

replace github.com/PQlite/crypto => ../crypto
