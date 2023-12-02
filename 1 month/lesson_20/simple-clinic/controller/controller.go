package controller

import (
	"simple-clinic/config"
	"simple-clinic/storage"
)

type Conn struct {
	cfg     *config.Config
	storage storage.StorageI
}

func NewController(cfg *config.Config, strg storage.StorageI) *Conn {
	return &Conn{cfg: cfg, storage: strg}
}
