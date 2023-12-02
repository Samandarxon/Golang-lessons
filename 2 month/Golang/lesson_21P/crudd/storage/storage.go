package storage

import "app/model"

type StorageI interface {
	User() UserRepoI
}

type UserRepoI interface {
	Create(model.UserCreate) (*model.User, error)
	GetById(model.UserPrimaryKey) (*model.User, error)
	GetList(model.UserGetListRequest) (*model.UserGetListResponse, error)
	Update(model.UserUpdate) (int64, error)
	Delete(model.UserPrimaryKey) (int64, error)
}
