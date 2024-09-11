package duckduckgo

import (
	"fmt"
	"github.com/sap-nocops/duckduckgogo/client"
	"search/internal/schema"
)

type DuckDuckGo struct{}

func (ddg DuckDuckGo) Query(query string) (schema.EngineResponseList, error) {
	ddgClient := client.NewDuckDuckGoSearchClient()
	res, err := ddgClient.SearchLimited(query, 10)
	if err != nil {
		fmt.Printf("error: %v", err)
		return nil, err
	}
	responsesList := make(schema.EngineResponseList, 0, 20)

	for _, result := range res {
		res := schema.EngineResponse{Title: result.Title, Url: result.FormattedUrl, Snippet: result.Snippet}
		responsesList = append(responsesList, res)
	}

	return responsesList, nil

}
