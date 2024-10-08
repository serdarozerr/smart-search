package google

import (
	"context"
	"google.golang.org/api/customsearch/v1"
	"google.golang.org/api/option"
	"os"
	"search/internal/schema"
)

var (
	apiKey = os.Getenv("API_KEY")
	cx     = os.Getenv("CX")
)

type GoogleSearch struct{}

func (g GoogleSearch) Query(query string) (schema.EngineResponseList, error) {

	ctx := context.TODO()

	service, err := customsearch.NewService(ctx, option.WithAPIKey(apiKey))
	if err != nil {
		panic(err)
	}
	resp, err := service.Cse.List().Cx(cx).Q(query).Do()
	if err != nil {
		return nil, err
	}

	responsesList := make(schema.EngineResponseList, 0, 20)

	for _, result := range resp.Items {
		res := schema.EngineResponse{Title: result.Title, Url: result.Link, Snippet: result.Snippet}
		responsesList = append(responsesList, res)
	}

	return responsesList, nil

}
