package controller

import (
	"simple-clinic/config"
	"simple-clinic/storage"
)

type conn struct {
	cfg     *config.Config
	storage storage.StorageI
}

func NewController(cfg *config.Config, strg storage.StorageI) *conn {
	return &conn{cfg: cfg, storage: strg}
}
