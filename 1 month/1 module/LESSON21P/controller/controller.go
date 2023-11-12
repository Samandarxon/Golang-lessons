package controller

import (
	"app/config"
	"app/storage"
)

type Connect struct {
	cfg     *config.Config
	storage storage.StorageI
}

func NewController(cfg *config.Config, strg storage.StorageI) *Connect {
	return &Connect{cfg: cfg, storage: strg}
}
