package models

import "github.com/google/uuid"

type Product struct {
	Id            uuid.UUID
	Name          string
	Price         int
	Quantity      int
	Category_id   uuid.UUID
	Category_Name string
	CreatedAt     string
	UpdatedAt     string
}

type ProductPrimaryKey struct {
	Id uuid.UUID
}

type ProductCreate struct {
	Id            uuid.UUID
	Name          string
	Price         int
	Quantity      int
	Category_id   uuid.UUID
	Category_Name string
}

type ProductUpdate struct {
	Id            uuid.UUID
	Price         int
	Quantity      int
	Category_id   uuid.UUID
	Category_Name string
}

type ProductGetListRequest struct {
	Offset int
	Limit  int
}

type ProductGetListResponse struct {
	Count    int
	Products []Product
}
