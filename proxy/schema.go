package main

import (
	graphql "github.com/graphql-go/graphql"
	"github.com/vektah/gqlparser"
	"github.com/vektah/gqlparser/ast"
)

// GetSchema grabs the resources needed and initiates the GraphQL schema object
func GetSchema() (schema *graphql.Schema, err error) {
	defer Recovery(&err)

	sources, err := getSdl()

	if err != nil {
		return nil, err
	}

	astSchema, err := parseSdl(sources)

	if err != nil {
		return nil, err
	}

	schema, err = ConvertSchema(astSchema)

	if err != nil {
		return nil, err
	}

	return
}

type source struct {
	name string
	sdl  string
}

func getSdl() ([]source, error) {
	return []source{
		source{
			name: "test.gql",
			sdl: `type Query { 
				hello: String @stub(value: "world")  
			}`,
		},
		source{
			name: "stub.gql",
			sdl:  `directive @stub(value: String) on FIELD_DEFINITION`,
		},
	}, nil
}

func parseSdl(sources []source) (*ast.Schema, error) {
	astSources := make([]*ast.Source, len(sources))

	for i := range sources {
		astSources[i] = &ast.Source{
			Name:  sources[i].name,
			Input: sources[i].sdl,
		}
	}

	schema, err := gqlparser.LoadSchema(astSources...)

	if err != nil {
		return nil, err
	}

	return schema, nil
}