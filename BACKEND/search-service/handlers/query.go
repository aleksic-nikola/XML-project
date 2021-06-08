package handlers

import (
	"log"
	"net/http"
	"xml/search-service/data"
)


type Queries struct {
	l *log.Logger
}

func NewQueries(l *log.Logger) *Queries {
	return &Queries{l}
}

func (u *Queries) GetQueries(rw http.ResponseWriter, r *http.Request) {
	u.l.Println("Handle GET Request for Queries")

	lp := data.GetQueries()

	err := lp.ToJSON(rw)

	if err != nil {
		http.Error(rw, "Unable to unmarshal queries json" , http.StatusInternalServerError)
	}
}
