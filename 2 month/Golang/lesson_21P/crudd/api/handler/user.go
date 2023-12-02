package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/spf13/cast"

	"app/model"
)

func (h *Handler) User(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		h.CreateUser(w, r)
	case "GET":
		var (
			values = r.URL.Query()
			method = values.Get("method")
		)
		fmt.Println(method)
		if method == "GET" {
			h.UserGetById(w, r)
		} else if method == "GET_LIST" {
			h.UserGetList(w, r)
		}
	case "PUT":
		h.UserUpdate(w, r)

	case "DELETE":
		h.UserDelete(w, r)
	}
}

func (h *Handler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var userRequest = model.UserCreate{}
	req, err := ioutil.ReadAll(r.Body)
	if err != nil {
		h.handlerResponse(w, 500, "Error while read: "+err.Error(), nil)
		return
	}

	if err := json.Unmarshal(req, &userRequest); err != nil {
		h.handlerResponse(w, 500, "error while unmarsh request: "+err.Error(), nil)
		return
	}

	resp, err := h.strg.User().Create(userRequest)
	if err != nil {
		h.handlerResponse(w, http.StatusBadRequest, err.Error(), nil)
		return
	}

	h.handlerResponse(w, http.StatusCreated, "User Created", resp)
}

func (h *Handler) UserUpdate(w http.ResponseWriter, r *http.Request) {
	var userId = r.URL.Query().Get("id")
	var userName = r.URL.Query().Get("name")

	fmt.Println(userId, userName)

	resp, err := h.strg.User().Update(model.UserUpdate{
		Id:   userId,
		Name: userName,
	})

	if err != nil {
		h.handlerResponse(w, http.StatusInternalServerError, "User does not exist: "+err.Error(), nil)
		return
	}

	h.handlerResponse(w, http.StatusOK, "User Updated:", resp)
}
func (h *Handler) UserDelete(w http.ResponseWriter, r *http.Request) {
	var userId = r.URL.Query().Get("id")
	resp, err := h.strg.User().Delete(model.UserPrimaryKey{
		Id: userId,
	})

	if err != nil {
		h.handlerResponse(w, http.StatusInternalServerError, "User does not exist: "+err.Error(), nil)
		return
	}

	h.handlerResponse(w, http.StatusOK, "User Deleted:", resp)
}
func (h *Handler) UserGetById(w http.ResponseWriter, r *http.Request) {
	var userId = r.URL.Query().Get("id")

	resp, err := h.strg.User().GetById(model.UserPrimaryKey{Id: userId})
	if err != nil {
		h.handlerResponse(w, 500, "User does not exist: "+err.Error(), nil)
		return
	}

	h.handlerResponse(w, http.StatusOK, "User:", resp)
}

func (h *Handler) UserGetList(w http.ResponseWriter, r *http.Request) {
	var offsetstr = r.URL.Query().Get("offset")
	var limitstr = r.URL.Query().Get("limit")

	var offset = cast.ToInt(offsetstr)
	var limit = cast.ToInt(limitstr)

	resp, err := h.strg.User().GetList(model.UserGetListRequest{Offset: offset, Limit: limit})
	if err != nil {
		h.handlerResponse(w, 500, "User does not exist: "+err.Error(), nil)
		return
	}

	h.handlerResponse(w, http.StatusOK, "User:", resp)
}
