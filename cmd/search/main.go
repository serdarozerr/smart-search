package main

import (
	"fmt"
	llm "search/internal/llm"
	"search/internal/search/duckduckgo"
	//"google.golang.org/api/googleapi/transport"

	//"net/http"
	"os"
	//"search/internal/service"
)

var (
	apiKey = os.Getenv("API_KEY")
	cx     = os.Getenv("CX")
)

func main() {

	//http.HandleFunc("/search", service.Search)
	//
	//err := http.ListenAndServe(":8080", nil)
	//if err != nil {
	//	panic(err)
	//}

	query := "What are the main components of the EU AI Act?"

	ddg := duckduckgo.DuckDuckGo{}
	ret, err := ddg.Query(query)
	if err != nil {
		fmt.Println(err)
	}
	//fmt.Printf("%s", t)

	//g := google2.GoogleSearch{}
	//ret, err := g.Query(query)
	//if err != nil {
	//	fmt.Println(err)
	//}

	llmEnhancer := llm.Enhancer{}
	res, err := llmEnhancer.FinalizeResult(query, ret)
	if err != nil {
		fmt.Println(err)
	}
	println(res)

}

//
//func gooleSearch() {
//
//	ctx := context.Background()
//
//	service, err := customsearch.NewService(ctx, option.WithAPIKey(apiKey))
//	if err != nil {
//		panic(err)
//	}
//	resp, err := service.Cse.List().Cx(cx).Q("what is the latest pixel phone").Do()
//	if err != nil {
//		panic(err)
//	}
//
//	for i, result := range resp.Items {
//		fmt.Printf("#%d: %s\n", i+1, result.Title)
//		fmt.Printf("\t%s\n", result.Snippet)
//	}
//
//}
//
//func duckDuckGoSearch() {
//
//}
