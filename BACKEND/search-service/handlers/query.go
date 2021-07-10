package handlers

import (
	"fmt"
	"log"
	"net/http"
	"xml/search-service/data"
	"xml/search-service/service"
)


type QueryHandler struct {
	L *log.Logger
	Service *service.QueryService
}

func NewQueries(l *log.Logger, service *service.QueryService) *QueryHandler {
	return &QueryHandler{l, service}
}

func (u *QueryHandler) CreateQuery(rw http.ResponseWriter, r *http.Request) {
	fmt.Println("creating")
	var query data.Query
	err := query.FromJSON(r.Body)
	if err != nil {
		u.L.Println(err)
		rw.WriteHeader(http.StatusBadRequest)
		return
	}
	fmt.Println(query)

	err = u.Service.CreateQuery(&query)
	if err != nil {
		fmt.Println(err)
		rw.WriteHeader(http.StatusExpectationFailed)
	}
	rw.WriteHeader(http.StatusCreated)
	rw.Header().Set("Content-Type", "application/json")
}

func (u *QueryHandler) GetQueries(rw http.ResponseWriter, r *http.Request) {
	u.L.Println("Handle GET Request for Queries")

	lp := data.GetQueries()

	err := lp.ToJSON(rw)

	if err != nil {
		http.Error(rw, "Unable to unmarshal queries json" , http.StatusInternalServerError)
	}
}
