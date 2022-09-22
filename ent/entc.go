//go:build ignore
// +build ignore

package main

import (
	"log"

	"entgo.io/contrib/entgql"
	"entgo.io/ent/entc"
	"entgo.io/ent/entc/gen"
)

func main() {
	ex, err := entgql.NewExtension(
		entgql.WithWhereFilters(true),
		entgql.WithConfigPath("../gqlgen.yml"),
		entgql.WithSchemaPath("../graphql/ent.graphqls"),
	)
	if err != nil {
		log.Fatalf("error creating entgql extension: %v", err)
	}
	err = entc.Generate("./schema", &gen.Config{}, entc.Extensions(ex))
	if err != nil {
		log.Fatalf("running ent codegen failed: %v", err)
	}
}
