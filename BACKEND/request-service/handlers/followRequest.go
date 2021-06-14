package handlers

import (
	_ "encoding/json"
	"fmt"
	"log"
	"net/http"
	"xml/request-service/data"
	"xml/request-service/service"
)

type FollowRequestHandler struct {
	L *log.Logger
	Service *service.FollowRequestService
}

func NewFollowRequest(l *log.Logger, service *service.FollowRequestService) *FollowRequestHandler {
	return &FollowRequestHandler{l, service}
}


func (handler *FollowRequestHandler) CreateFollowRequest(rw http.ResponseWriter, r *http.Request) {
	fmt.Println("creating")
	var followReq data.FollowRequest
	err := followReq.FromJSON(r.Body)
	if err != nil {
		handler.L.Println(err)
		rw.WriteHeader(http.StatusBadRequest)
		return
	}
	fmt.Println(followReq)

	err = handler.Service.CreateFollowRequest(&followReq)
	if err != nil {
		fmt.Println(err)
		rw.WriteHeader(http.StatusExpectationFailed)
	}
	rw.WriteHeader(http.StatusCreated)
	rw.Header().Set("Content-Type", "application/json")
}

func (handler *FollowRequestHandler) GetFollowRequestsDB(rw http.ResponseWriter, r *http.Request) {
	fmt.Println("getting")
	var followReqs []data.FollowRequest
	var err error

	followReqs, err = handler.Service.GetAllRequests()
	if err != nil {
		fmt.Println(err)
		rw.WriteHeader(http.StatusExpectationFailed)
	}

	fmt.Println(followReqs)

	rw.WriteHeader(http.StatusOK)
	//err = followReqs.ToJSON(rw)
	if err != nil {
		http.Error(rw, "Unable to marshal json", http.StatusInternalServerError)
	}

	rw.Header().Set("Content-Type", "application/json")
}




func (p *FollowRequestHandler) GetFollowRequests(rw http.ResponseWriter, r *http.Request) {
	p.L.Println("Handle GET Request")

	lp := data.GetFollowRequests()

	err := lp.ToJSON(rw)

	if err != nil {
		http.Error(rw, "Unable to marshal json", http.StatusInternalServerError)
	}
}
