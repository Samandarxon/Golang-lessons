package api

import (
	"net/http"

	"app/api/handler"
	"app/config"
	"app/storage"
)

func NewApi(cfg *config.Config, strg storage.StorageI) {
	handler := handler.NewHanler(cfg, strg)

	http.HandleFunc("/user", handler.User)
}
