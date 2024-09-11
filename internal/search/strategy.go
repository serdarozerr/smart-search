package search

import "search/internal/schema"

type Strategy interface {
	Query(query string) (schema.EngineResponseList, error)
}
