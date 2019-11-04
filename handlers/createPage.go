package handlers

import (
	"github.com/RichardKnop/machinery/v1/tasks"
	"github.com/go-chi/render"
	"github.com/piwonskp/feldscher/fetcher"
	"github.com/piwonskp/feldscher/models"
	"log"
	"net/http"
)

type PageAddedResponse struct {
	ID int `json:"id"`
}

func CreatePage(w http.ResponseWriter, r *http.Request) {
	const maxSizeInBytes = 1024 * 1024
	r.Body = http.MaxBytesReader(w, r.Body, maxSizeInBytes)
	var page models.Page
	err := render.DecodeJSON(r.Body, &page)
	if err != nil {
		// Seriously Go?
		// https://github.com/golang/go/blob/3409ce39bfd7584523b7a8c150a310cea92d879d/src/net/http/request.go#L1161
		if err.Error() == "http: request body too large" {
			w.WriteHeader(http.StatusRequestEntityTooLarge)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	DBAddPage(&page)

	server, err := fetcher.StartServer()
	if err != nil {
		log.Println(err)
	}
	signature := &tasks.Signature{
		Name: "schedule_fetch",
		Args: []tasks.Arg{
			{
				Type:  "int",
				Value: page.ID,
			},
			{
				Type:  "string",
				Value: page.URL,
			},
			{
				Type:  "int",
				Value: page.Interval,
			},
		},
	}

	_, err = server.SendTask(signature)
	if err != nil {
		log.Println(err)
	}
	render.JSON(w, r, PageAddedResponse{ID: page.ID})
}

func DBAddPage(page *models.Page) {
	models.DB.Create(page)
}
