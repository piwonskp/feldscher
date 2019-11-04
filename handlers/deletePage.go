package handlers

import (
	"github.com/RichardKnop/machinery/v1/tasks"
	"github.com/go-chi/render"
	"github.com/piwonskp/feldscher/fetcher"
	"github.com/piwonskp/feldscher/models"
	"log"
	"net/http"
)

type DeletedPageResponse struct {
	ID int `json:"id"`
}

func DeletePage(w http.ResponseWriter, r *http.Request) {
	id := r.Context().Value("id").(int)

	deleted := DBRemovePage(id)
	if deleted == false {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	server, err := fetcher.StartServer()
	if err != nil {
		log.Println(err)
	}
	signature := &tasks.Signature{
		Name: "stop_fetch",
		Args: []tasks.Arg{
			{
				Type:  "int",
				Value: id,
			},
		},
	}
	_, err = server.SendTask(signature)
	if err != nil {
		log.Println(err)
	}

	render.JSON(w, r, DeletedPageResponse{ID: id})
}

func DBRemovePage(ID int) bool {
	page := models.Page{ID: ID}
	d := models.DB.Delete(&page)
	return d.RowsAffected != 0
}
