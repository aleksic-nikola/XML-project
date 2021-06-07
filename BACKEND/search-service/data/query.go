package data

import (
	"encoding/json"
	"io"
)

type Query struct {
	Input     string    `json:"input"`
	QueryType QueryType `json:"query_type"`
}

func (q *Query) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(q)
}

func (q *Queries) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(q)
}

// declaring the collection
type Queries []*Query

func GetQueries() Queries {
	return queryList
}

var queryList = []*Query{

	{
		Input: "dparipovic98",
		QueryType: PROFILE,
	},
}
