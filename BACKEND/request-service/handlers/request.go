package handlers

import (
	//"context"
	"log"
	"net/http"
	"xml/request-service/data"

	//"net/http"
	//"strconv"
	//"github.com/gorilla/mux"
	//"auth-service/data"
)

type Requests struct {
	l *log.Logger
}

func NewRequest(l *log.Logger) *Requests {
	return &Requests{l}
}

func (p *Requests) GetRequests(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle GET Request")

	lp := data.GetRequests()

	err := lp.ToJSON(rw)

	if err != nil {
		http.Error(rw, "Unable to marshal json", http.StatusInternalServerError)
	}
}





