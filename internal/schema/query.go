package schema

import (
	"encoding/json"
	"net/http"
)

type Request struct {
	Query        string `json:"query"`
	SearchEngine string `json:"search_engine"`
}
type Response struct {
	Query  string `json:"query"`
	Answer string `json:"answer"`
}

func (q *Request) Decode(r *http.Request) error {

	err := json.NewDecoder(r.Body).Decode(q)
	if err != nil {
		return err
	}
	return nil
}

func (q Response) Encode(result EngineResponseList) ([]byte, error) {
	res, err := json.Marshal(result)
	return res, err
}
