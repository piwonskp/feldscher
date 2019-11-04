package fetcher

import (
	"github.com/piwonskp/feldscher/models"
	"io/ioutil"
	"net/http"
	"time"
)

var Channels map[int]chan struct{}

func ScheduleFetch(pageID int, url string, interval int) error {
	ticker := time.NewTicker(time.Duration(interval) * time.Second)
	Channels[pageID] = make(chan struct{})
	go fetchPage(pageID, url, ticker, Channels[pageID])
	return nil
}

func StopFetch(pageID int) error {
	close(Channels[pageID])
	return nil
}

func makeRequest(id int, url string) models.FetchedPage {
	client := http.Client{Timeout: 5 * time.Second}

	start := time.Now()
	resp, err := client.Get(url)
	var body *string
	if err == nil {
		defer resp.Body.Close()
		bytes, _ := ioutil.ReadAll(resp.Body)
		content := string(bytes)
		body = &content
	}

	duration := time.Since(start).Seconds()

	return models.FetchedPage{PageID: id,
		Response: body,
		Duration: duration, CreatedAt: start.Unix()}
}

func fetchPage(pageId int, url string, ticker *time.Ticker, quit chan struct{}) {
	for {
		select {
		case <-ticker.C:
			page := makeRequest(pageId, url)
			models.DB.Create(&page)
		case <-quit:
			ticker.Stop()
			return
		}
	}
}
