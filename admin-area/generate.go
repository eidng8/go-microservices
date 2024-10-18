package main

//go:generate go run -modfile=./tools/go.mod ./tools/entc.go
//go:generate go run -modfile=./tools/go.mod github.com/deepmap/oapi-codegen/v2/cmd/oapi-codegen --config=./tools/oapi-codegen.json ./ent/openapi.json
