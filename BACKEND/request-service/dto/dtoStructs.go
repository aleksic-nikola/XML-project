package dtoRequest

import (
	"encoding/json"
	"io"
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
	FollowToUsername string `json:"follow-to-username"`
}

func (ur *UsernameRoleDto) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(ur)
}

func (scr *SensitiveContentReportDto) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(scr)
}