package controller

import (
	"app/config"
	"app/storage"
)

type Connect struct {
	cfg     *config.Config_DB
	storage storage.StorageI
}

func NewController(cfg *config.Config_DB, strg storage.StorageI) *Connect {
	return &Connect{cfg: cfg, storage: strg}
}
