package models

type User struct {
	Id        string `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Balance   int    `json:"balance"`
}
type CreateUser struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Balance   int    `json:"balance"`
}

type PrimeryKeyUser struct {
	Id string `json:"id"`
}

type UpadetUser struct {
	Id        string `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Balance   int    `json:"balance"`
}

type GetListUserRequest struct {
	Limit  int `json:"limit"`
	Offset int `json:"offset"`
}

type GetListUserResponse struct {
	Count int     `json:"count"`
	Users []*User `json:"users"`
}
