package handlers

import (
	"log"
	"net/http"
	"xml/content-service/data"
)

type Stories struct {
	l *log.Logger
}

func NewStories(l *log.Logger) *Stories {
	return &Stories{l}
}

func (p *Stories) GetStories(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle GET Request for Posts")

	ls := data.GetStories()

	err := ls.ToJSON(rw)

	if err != nil {
		http.Error(rw, "Unable to unmarshal stories json" , http.StatusInternalServerError)
	}
}
