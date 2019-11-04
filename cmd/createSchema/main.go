package main

import (
	"github.com/piwonskp/feldscher/models"
)

func main() {
	defer models.DB.Close()

	models.DB.AutoMigrate(&models.Page{}, &models.FetchedPage{})
	models.DB.Model(&models.FetchedPage{}).AddForeignKey(
		"page_id", "pages(id)", "CASCADE", "RESTRICT")
}
