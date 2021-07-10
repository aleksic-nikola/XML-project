package data

import (
	"encoding/json"
	"gorm.io/gorm"
	"io"
)

type Query struct {
	gorm.Model
	Input     string    `json:"input" gorm:"uniqueText"`
	QueryType QueryType `json:"query_type" gorm:"type:int"`
}

func (q *Query) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(q)
}

func (q *Queries) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(q)
}

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
