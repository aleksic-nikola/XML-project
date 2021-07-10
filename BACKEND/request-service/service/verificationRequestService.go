package service

import (
	"fmt"
	"xml/request-service/data"
	dtoRequest "xml/request-service/dto"
	"xml/request-service/repository"
)

type VerificationRequestService struct {
	Repo *repository.VerificationRequestRepository
}

func (service *VerificationRequestService) CreateVerificationRequest(verificationRequest *data.VerificationRequest) error {
	error := service.Repo.CreateVerificationRequest(verificationRequest)
	return error
}

func (service *VerificationRequestService) GetInProgressVerificationRequests() data.VerificationRequests {
	requests := service.Repo.GetInProgressVerificationRequests()
	return requests
}

func (service *VerificationRequestService) CreateInProgressVerificationRequest(verificationRequestDto *dtoRequest.VerificationRequestDto, userDto *dtoRequest.UsernameRoleDto ) error {
	var verificationRequest data.VerificationRequest
	verificationRequest.Request.SentBy = userDto.Username
	verificationRequest.Request.Status = 0
	verificationRequest.Category = verificationRequestDto.Category
	verificationRequest.Image = verificationRequestDto.Image
	verificationRequest.Name = verificationRequestDto.Name
	verificationRequest.LastName = verificationRequestDto.LastName

	num := service.Repo.CheckIfUserHasActiveVR(userDto.Username)

	if num > 0 {
		return fmt.Errorf("User has verification request in process or is already verified!")
	}

	err := service.Repo.CreateVerificationRequest(&verificationRequest)

	return err
}

func (service *VerificationRequestService) UpdateVerificationRequest(dto *dtoRequest.UpdateVerificationRequestDto) (error, *data.VerificationRequest) {

	err, verifiedRequest := service.Repo.FindById(dto.Id)

	if err != nil {
		return err, verifiedRequest
	}

	verifiedRequest.Request.Status = dto.NewStatus

	verifiedRequest = service.Repo.UpdateVerificationRequest(verifiedRequest)

	fmt.Println(verifiedRequest)

	return nil, verifiedRequest
}

