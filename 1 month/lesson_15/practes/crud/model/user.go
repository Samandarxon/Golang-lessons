package model

type User struct {
	Id        string
	FirstName string
	LastName  string
	Age       int
	CreatedAt string
}

type GetListUserRequest struct {
	Offset int
	Limit  int
}

type GetListUserRespons struct {
	Count int
	Users []User
}
