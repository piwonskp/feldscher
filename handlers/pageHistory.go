package handlers

import (
	"github.com/go-chi/render"
	"github.com/piwonskp/feldscher/models"
	"net/http"
)

func FetchedPagesHistory(w http.ResponseWriter, r *http.Request) {
	id := r.Context().Value("id")
	var page models.Page
	result := models.DB.First(&page, id)
	if result.RowsAffected == 0 {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	var pages []models.FetchedPage
	models.DB.Where("page_id = ?", id).Find(&pages)
	render.JSON(w, r, pages)
}
