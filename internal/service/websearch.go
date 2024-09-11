package service

import (
	"encoding/json"
	"net/http"
	"search/internal/llm"
	"search/internal/schema"

	"search/internal/search"
)

func WebSearch(w http.ResponseWriter, r *http.Request) {

	request := schema.Request{}
	err := request.Decode(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	context := search.Context{}
	err = context.SelectStrategy(request.SearchEngine)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	res, err := context.ExecuteSearchEngine(request.Query)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	llmEnhancer := llm.Enhancer{}
	llmAnswer, err := llmEnhancer.FinalizeResult(request.Query, res)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := schema.Response{Query: request.Query, Answer: llmAnswer}
	resByte, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(resByte)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}

//curl  -X POST http://localhost:8080/search -H "Content-Type: application/json" -d '{"query": "Who is the jony Bravo?", "search_engine":"google"}'
