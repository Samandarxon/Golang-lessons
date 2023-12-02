package model

type User struct {
	Id        string `json:"id"`
	Name      string `json:"name"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type UserPrimaryKey struct {
	Id string `json:"id"`
}

type UserCreate struct {
	Name string `json:"name"`
}

type UserUpdate struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type UserGetListRequest struct {
	Offset int `json:"offset"`
	Limit  int `json:"limit"`
}

type UserGetListResponse struct {
	Count int    `json:"count"`
	Users []User `json:"users"`
}
