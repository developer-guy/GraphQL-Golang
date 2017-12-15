package main

import (
	"github.com/graphql-go/graphql"
	"github.com/developer-guy/grapql-sample"
	"fmt"
	"log"
	"encoding/json"
)

func main() {

	humanType := graphql.NewObject(graphql.ObjectConfig{
		Name: "human", Fields: graphql.Fields{
			"Name": &graphql.Field{
				Type: graphql.String,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if human, ok := p.Source.(human.Human); ok {
						return human.Name, nil
					}
					return nil, nil
				},
			},
			"Surname": &graphql.Field{
				Type: graphql.String,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if human, ok := p.Source.(human.Human); ok {
						return human.Surname, nil
					}
					return []interface{}{}, nil
				},
			},
			"Gender": &graphql.Field{
				Type: graphql.String,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if human, ok := p.Source.(human.Human); ok {
						return human.Gender, nil
					}
					return []interface{}{}, nil
				},
			},
			"Age": &graphql.Field{
				Type: graphql.Int,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if human, ok := p.Source.(human.Human); ok {
						return human.Age, nil
					}
					return []interface{}{}, nil
				},
			},
		},
	})

	fields := graphql.Fields{
		"human": &graphql.Field{
			Type: humanType,
			Args: graphql.FieldConfigArgument{
				"Name": &graphql.ArgumentConfig{
					Type:         graphql.String,
					DefaultValue: "Batuhan",
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return human.HumanByName(p.Args["Name"].(string)), nil
			},
		},
	}

	rootQuery := graphql.ObjectConfig{
		Name:   "RootQuery",
		Fields: fields,
	}

	schemaConfig := graphql.SchemaConfig{
		Query: graphql.NewObject(rootQuery),
	}

	schema, err := graphql.NewSchema(schemaConfig)

	if err != nil {
		panic(err)
	}

	query := `query HumanByName($Name: String){
				human(Name: $Name){
					Name
					Surname
					Gender
					Age
				}
                }

			`

	params := graphql.Params{
		Schema:         schema,
		RequestString:  query,
		VariableValues: map[string]interface{}{"Name": "Asena"}}

	result := graphql.Do(params)
	if len(result.Errors) > 0 {
		log.Fatalf("failed to execute graphql operation, errors: %+v", result.Errors)
	}

	resultJson, _ := json.Marshal(result)
	fmt.Printf("%s \n", resultJson) // {“data”:{“hello”:”world”}}
}
