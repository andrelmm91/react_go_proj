package graphql

import (
	"backend/internal/models"
	"strings"

	"github.com/graphql-go/graphql"
)

type GraphQL struct {
	Movies []*models.Movie
	QueryString string
	Config graphql.SchemaConfig
	Fields graphql.Fields
	movieType graphql.Object
}

func New(movies []*models.Movie) *GraphQL {
	
	// describe the data as it displays in the database
	var movieType = graphql.NewObject(
		graphql.ObjectConfig{
			Name: "Movie",
			Fields: graphql.Fields{
				"id": &graphql.Field{
					Type: graphql.Int,
				},
				"title": &graphql.Field{
					Type: graphql.String,
				},
				"description": &graphql.Field{
					Type: graphql.String,
				},
				"release_date": &graphql.Field{
					Type: graphql.DateTime,
				},
				"runtime": &graphql.Field{
					Type: graphql.Int,
				},
				"mpaa_rating": &graphql.Field{
					Type: graphql.String,
				},
				"image": &graphql.Field{
					Type: graphql.String,
				},
				"created_at": &graphql.Field{
					Type: graphql.DateTime,
				},
				"updated_at": &graphql.Field{
					Type: graphql.DateTime,
				},
			},
		},
	) 

	// define the available actions it will performs
	var fields = graphql.Fields{

		"list": &graphql.Field{
			Type: graphql.NewList(movieType),
			Description: "get all movies",
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				return movies, nil
			},
		},

		"search":&graphql.Field{
			Type: graphql.NewList(movieType),
			Description: "search movies by title",
			Args: graphql.FieldConfigArgument{
				"titleContains": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
			},
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				var theList []*models.Movie
				search, ok := params.Args["titleContains"].(string)
				if ok {
					for _, currentMovie := range movies {
						if strings.Contains(strings.ToLower(currentMovie.Title), strings.ToLower(search)) {
							theList = append(theList, currentMovie)
						}
					}
				}
				return theList, nil
			},
		},

		"get": &graphql.Field{
			Type: movieType,
			Description: "get movie by id",
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{
					Type: graphql.Int,
				},
			},
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				id, ok := params.Args["id"].(int)
				if ok {
					for _, movie := range movies {
						if movie.ID == id {
							return movie, nil
						}
					}
				}
				return nil, nil
			},
		},
	}

	return &GraphQL{
		Movies: movies,
		Fields: fields,
		movieType: *movieType,
	}
}

