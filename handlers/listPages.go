package handlers

import (
	"github.com/go-chi/render"
	"github.com/piwonskp/feldscher/models"
	"net/http"
)

func ListPages(w http.ResponseWriter, r *http.Request) {
	var pages []models.Page
	models.DB.Find(&pages)
	render.JSON(w, r, pages)
}
