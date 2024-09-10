package duckduckgo

import (
	"fmt"
	"github.com/sap-nocops/duckduckgogo/client"
	"search/internal/search"
)

type DuckDuckGo struct{}

func (ddg DuckDuckGo) Query(query string) (search.ResponseList, error) {
	duckdg := client.NewDuckDuckGoSearchClient()
	res, err := duckdg.SearchLimited(query, 10)
	if err != nil {
		fmt.Printf("error: %v", err)
		return nil, err
	}
	responsesList := make(search.ResponseList, 0, 20)

	for _, result := range res {
		res := search.Response{Title: result.Title, Url: result.FormattedUrl, Snippet: result.Snippet}
		responsesList = append(responsesList, res)
	}

	return responsesList, nil

}
