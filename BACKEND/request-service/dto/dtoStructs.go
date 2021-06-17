package dtoRequest


type FollowRequestDto struct {
	SentBy string `json:"sentBy"`
	ForWho     string `json:"forWho"`
}

type UsernameRoleDto struct {
	Username string `json:"username"`
	Role     string `json:"role"`
}

type ProfileForFollow struct {
	FollowToUsername string `json:"follow-to-username"`
}