package request

type UserCreateRequest struct {
	Name string `json:"name"`
}

type UserCountAddRequest struct {
	Count int `json:"count"`
}
