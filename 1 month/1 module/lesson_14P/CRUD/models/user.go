package models

import (
	"time"

	"github.com/google/uuid"
)

// User Crud
// Name
// Surname
// Age
// Dateof Birth

type User struct {
	Id          uuid.UUID
	Name        string
	Surname     string
	UserName    string
	Password    string
	Age         int
	DateOfBirth string
	CreatedAt   string
	UpdatedAt   string
}

type UserCreate struct {
	Name        string
	SurName     string
	UserName    string
	Password    string
	Age         int
	DateOfBirth string
}

type UserGetListRequest struct {
	OffSet int
	Limit  int
}

type UserGetListRespons struct {
	Count int
	Users []User
}

type UserPrimaryKey struct {
	Id uuid.UUID
}
type DateAndAge struct {
	Age  int
	Date string
}
type UserUpdateRequest struct {
	Id       uuid.UUID
	Name     string
	Surname  string
	Username string
	Password string
}

type UserUpdateResponse struct {
	User
}

type LoginRequest struct {
	UserName string
	Password string
}
type LoginRespons struct {
	Message string
}
type RegisterRequest struct {
	Name        string
	SurName     string
	UserName    string
	Password    string
	Age         int
	DateOfBirth string
}
type RegisterRespons struct {
	Message string
}

func (u *DateAndAge) AgeAndDate(date string) {
	x_parse, _ := time.Parse("2006-01-01", date)
	u.Age = int(time.Now().Sub(x_parse).Hours() / 24 / 365)
	u.Date = date
}
