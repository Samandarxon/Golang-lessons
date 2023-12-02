package handler

import (
	"essy_travel/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) CreateAirport(c *gin.Context) {
	var airport = models.CreateAirport{}
	err := c.ShouldBindJSON(&airport)
	if err != nil {
		handleResponse(c, http.StatusBadRequest, "Error while json decoding"+err.Error())
		return
	}
	_, err = h.strg.Airport().Create(airport)
	if err != nil {
		handleResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	handleResponse(c, http.StatusCreated, "Airport created")
}

func (h *Handler) AirportGetById(c *gin.Context) {
	var airport = models.AirportPrimaryKey{}
	err := c.ShouldBindJSON(&airport)
	if err != nil {
		handleResponse(c, http.StatusBadRequest, "Error while json decoding"+err.Error())
		return
	}
	resp, err := h.strg.Airport().GetById(airport)
	if err != nil {
		handleResponse(c, 500, "Airport does not exist: "+err.Error())
		return
	}

	handleResponse(c, http.StatusOK, resp)
}

func (h *Handler) AirportGetList(c *gin.Context) {
	var airport = models.GetListAirportRequest{}
	err := c.ShouldBindJSON(&airport)
	if err != nil {
		handleResponse(c, http.StatusBadRequest, "Error while json decoding"+err.Error())
		return
	}
	resp, err := h.strg.Airport().GetList(airport)
	if err != nil {
		handleResponse(c, 500, "Airport does not exist: "+err.Error())
		return
	}

	handleResponse(c, http.StatusOK, resp)
}

func (h *Handler) AirportUpdate(c *gin.Context) {
	var airport = models.UpdateAirport{}
	err := c.ShouldBindJSON(&airport)
	if err != nil {
		handleResponse(c, http.StatusBadRequest, "Error while json decoding"+err.Error())
		return
	}
	_, err = h.strg.Airport().Update(airport)
	if err != nil {
		handleResponse(c, 500, "Airport does not update: "+err.Error())
		return
	}

	handleResponse(c, http.StatusAccepted, "Updated")
}

func (h *Handler) AirportDelete(c *gin.Context) {
	var airport = models.AirportPrimaryKey{}
	err := c.ShouldBindJSON(&airport)
	if err != nil {
		handleResponse(c, http.StatusBadRequest, "Error while json decoding"+err.Error())
		return
	}
	_, err = h.strg.Airport().Delete(airport)
	if err != nil {
		handleResponse(c, 500, "Airport does not delete: "+err.Error())
		return
	}

	handleResponse(c, http.StatusAccepted, "Deleted:")
}

func (h *Handler) AirportSearch(c *gin.Context) {
	var airport = models.Airport{}
	err := c.ShouldBindJSON(&airport)
	if err != nil {
		handleResponse(c, http.StatusBadRequest, "Error while json decoding"+err.Error())
		return
	}
	resp, err := h.strg.Airport().Search(airport)
	if err != nil {
		handleResponse(c, 500, "Airport does not find: "+err.Error())
		return
	}

	handleResponse(c, http.StatusAccepted, resp)
}
