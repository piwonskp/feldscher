package main

import (
	"github.com/piwonskp/feldscher/fetcher"
	"github.com/piwonskp/feldscher/models"
	"log"
)

func main() {
	fetcher.Channels = make(map[int]chan struct{})
	server, err := fetcher.StartServer()
	if err != nil {
		log.Println(err)
		return
	}
	worker := server.NewWorker("fetcher", 5)

	var pages []models.Page
	models.DB.Find(&pages)
	for _, page := range pages {
		fetcher.ScheduleFetch(page.ID,
			page.URL,
			page.Interval)
	}

	err = worker.Launch()
	if err != nil {
		log.Println(err)
		return
	}

}
