package schema

import (
	"encoding/json"
	"net/http"
)

type QueryRequest struct {
	query string
}

type QueryResponse struct {
	query  string
	answer string
}

func (q *QueryRequest) Decode(r *http.Request) error {

	err := json.NewDecoder(r.Body).Decode(q)
	if err != nil {
		return err
	}
	return nil
}
