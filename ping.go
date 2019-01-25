package main

import (
	"examplego/database"
	"examplego/models"
	"fmt"
	"github.com/joho/godotenv"
	"net/http"
	"time"
)

func main() {
	// load .env environment variables
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	// initializes database
	db, _ := database.Initialize()
	var urls []models.URL

	if err := db.Find(&urls).Error; err != nil {
		fmt.Println(err)
	} else {
		monitor(urls)
	}
}

func monitor(urls []models.URL) {
	ch := make(chan *models.URL, len(urls)) // buffered
	failureThresholdMap := make(map[string]int)
	for _, url := range urls {
		temp :=url
		go crawlUrl(&temp, ch, failureThresholdMap)
	}
	for {
		select {
		case r := <-ch:
			fmt.Printf("%s was fetched\n", r.Url)
		case <-time.After(50 * time.Millisecond):
			fmt.Printf(".")
		}
	}
}

func crawlUrl(url *models.URL, ch chan *models.URL, failureThresholdMap map[string]int) {
	for {
		if failureThresholdMap[url.Url] < url.FailureThreshold {
			fmt.Printf("Fetching %s \n", url.Url)
			resp, err := fetchUrl(url)
			if err != nil || resp.StatusCode != 200 {
				failureThresholdMap[url.Url] += 1
			}
			ch <- url
			time.Sleep(time.Duration(url.Frequency) * time.Second)
		} else {
			fmt.Printf("Threshold reached %s \n", url.Url)
			break
		}
	}
}

func fetchUrl(url *models.URL) (res *http.Response, err error) {
	timeout := time.Duration(time.Duration(url.CrawlTimeout) * time.Second)
	client := http.Client{
		Timeout: timeout,
	}
	fmt.Printf("Fetching %s \n", url.Url)
	resp, err := client.Get(url.Url)
	resp.Body.Close()
	return resp,err
}


