package search

import (
	"errors"
	"search/internal/schema"
	"search/internal/search/duckduckgo"
	"search/internal/search/google"
)

type Context struct {
	SearchEngine Strategy
}

var InvalidChoiceError = errors.New("invalid choice error.Select either google or duckduckgo")

func (c *Context) SelectStrategy(policy string) error {
	switch policy {
	case "google":
		c.SearchEngine = google.GoogleSearch{}

	case "duckduckgo":
		c.SearchEngine = duckduckgo.DuckDuckGo{}
	default:
		return InvalidChoiceError
	}
	return nil
}

func (c Context) ExecuteSearchEngine(query string) (schema.EngineResponseList, error) {

	res, err := c.SearchEngine.Query(query)
	return res, err
}
