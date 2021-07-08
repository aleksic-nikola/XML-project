package dtoRequest

import (
	"encoding/json"
	"io"
	"xml/request-service/data"
)

type FollowRequestDto struct {
	SentBy string `json:"sentBy"`
	ForWho     string `json:"forWho"`
}

type SensitiveContentReportDto struct {
	PostID string `json:"post_id"`
	Note string `json:"note"`
}

type UsernameRoleDto struct {
	Username string `json:"username"`
	Role     string `json:"role"`
}

type ProfileForFollow struct {
	FollowToUsername string `json:"followToUsername"`
}

type VerificationRequestDto struct {
	Category data.VerifiedType `json:"verifiedType"`
	Image string `json:"image"`
	Name string `json:"name"`
	LastName string `json:"lastname"`
}

type UpdateVerificationRequestDto struct {
	Id uint `json:"id"`
	NewStatus data.RequestStatus `json:"new_status"`
}

type NewVerified struct {
	Username string `json:"username"`
	VerifiedType data.VerifiedType `json:"verified_type"`
}




func (ur *UsernameRoleDto) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(ur)
}

func (scr *SensitiveContentReportDto) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(scr)
}

func (vrdto *VerificationRequestDto) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(vrdto)
}

func (uvrdto *UpdateVerificationRequestDto) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(uvrdto)
}


