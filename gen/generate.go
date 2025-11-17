package gen

//go:generate ogen --package openapi --target ./openapi --clean ../openapi/openapi.yaml
//go:generate go run github.com/sqlc-dev/sqlc/cmd/sqlc -- generate -f ../sqlc.yaml