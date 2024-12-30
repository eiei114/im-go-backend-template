package response

type UserCreateResponse struct {
	Token string `json:"token"`
}

type UserGetResponse struct {
	Id    string `json:"id"`
	Name  string `json:"name"`
	Count int    `json:"count"`
}

type UserCountAddResponse struct {
	Count int `json:"count"`
}

type UserDestroyResponse struct {
	Message string `json:"message"`
}
