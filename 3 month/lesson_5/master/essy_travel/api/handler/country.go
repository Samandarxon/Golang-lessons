package handler

import (
	"essy_travel/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) CreateCountry(c *gin.Context) {
	var country = models.CreateCountry{}
	err := c.ShouldBindJSON(&country)
	if err != nil {
		c.JSON(400, "ShouldBindJSON err:"+err.Error())
		return
	}
	resp, err := h.strg.Country().Create(country)
	if err != nil {
		handleResponse(c, http.StatusInternalServerError, err)
		return
	}

	handleResponse(c, http.StatusCreated, resp)
}

func (h *Handler) CountryGetById(c *gin.Context) {
	var country = models.CountryPrimaryKey{}
	err := c.ShouldBindJSON(&country)
	if err != nil {
		handleResponse(c, http.StatusBadRequest, err)
		return
	}
	resp, err := h.strg.Country().GetById(country)
	if err != nil {
		handleResponse(c, 500, "Country does not exist: "+err.Error())
		return
	}

	handleResponse(c, http.StatusOK, resp)
}

func (h *Handler) CountryGetList(c *gin.Context) {
	var country = models.GetListCountryRequest{}
	err := c.ShouldBindJSON(&country)
	if err != nil {
		handleResponse(c, http.StatusBadRequest, err)
		return
	}
	resp, err := h.strg.Country().GetList(country)
	if err != nil {
		handleResponse(c, 500, "Country does not exist: "+err.Error())
		return
	}

	handleResponse(c, http.StatusOK, resp)
}

func (h *Handler) CountryUpdate(c *gin.Context) {
	var country = models.UpdateCountry{}
	err := c.ShouldBindJSON(&country)
	if err != nil {
		handleResponse(c, http.StatusBadRequest, err)
		return
	}
	_, err = h.strg.Country().Update(country)
	if err != nil {
		handleResponse(c, 500, "Country does not update: "+err.Error())
		return
	}

	handleResponse(c, http.StatusAccepted, "Updated:")
}

func (h *Handler) CountryDelete(c *gin.Context) {
	var country = models.CountryPrimaryKey{}
	err := c.ShouldBindJSON(&country)
	if err != nil {
		handleResponse(c, http.StatusBadRequest, err)
		return
	}
	_, err = h.strg.Country().Delete(country)
	if err != nil {
		handleResponse(c, 500, "Country does not delete: "+err.Error())
		return
	}

	handleResponse(c, http.StatusAccepted, "Deleted")
}

func (h *Handler) CountrySearch(c *gin.Context) {
	var country = models.Country{}
	err := c.ShouldBindJSON(&country)
	if err != nil {
		handleResponse(c, http.StatusBadRequest, "Error while json decoding"+err.Error())
		return
	}
	resp, err := h.strg.Country().Search(country)
	if err != nil {
		handleResponse(c, 500, "Country does not find: "+err.Error())
		return
	}

	handleResponse(c, http.StatusAccepted, resp)
}
