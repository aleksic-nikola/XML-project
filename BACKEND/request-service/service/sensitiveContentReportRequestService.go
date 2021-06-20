package service

import (
	"fmt"
	"xml/request-service/data"
	dtoRequest "xml/request-service/dto"
	"xml/request-service/repository"
)

type SensitiveContentReportRequestService struct {
	Repo *repository.SensitiveContentReportRequestRepository
}

func (service *SensitiveContentReportRequestService) CreateSensitiveContentReportRequest(sensitiveContentReportRequest *data.SensitiveContentReportRequest) error {
	error := service.Repo.CreateSensitiveContentReportRequest(sensitiveContentReportRequest)
	return error
}

func (service *SensitiveContentReportRequestService) CreateSensitiveContentReport(scrd *dtoRequest.SensitiveContentReportDto, user *dtoRequest.UsernameRoleDto) error {
	var sensitiveContentReportRequest data.SensitiveContentReportRequest
	sensitiveContentReportRequest.PostID = scrd.PostID
	sensitiveContentReportRequest.Note = scrd.Note
	sensitiveContentReportRequest.Request.SentBy = user.Username
	sensitiveContentReportRequest.Request.Status = data.INPROCESS
	fmt.Println(sensitiveContentReportRequest)

	error := service.Repo.CreateSensitiveContentReportRequest(&sensitiveContentReportRequest)
	return error
}