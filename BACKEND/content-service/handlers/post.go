package handlers

import (
	"log"
	"net/http"
	"xml/content-service/data"
)

type Posts struct {
	l *log.Logger
}

func NewPosts(l *log.Logger) *Posts {
	return &Posts{l}
}

func (p *Posts) GetPosts(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle GET Request for Posts")

	lp := data.GetPosts()

	err := lp.ToJSON(rw)

	if err != nil {
		http.Error(rw, "Unable to unmarshal posts json" , http.StatusInternalServerError)
	}
}