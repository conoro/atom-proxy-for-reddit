// Reddit Proxy
// This tool lets you request specific subreddits via RSS and generates its own RSS feed from them
// It exists to deal with Feedly being blocked for (presumably) excessive requests to Reddit
// V1 is completely hard-coded for me but could easily be made configurable.
// Next step is to get running on AWS Lambda
// Note I have vendored-in gofeed so I can add a needed User-Agent for the requests.
// A simple reverse proxy might do the same job but this was quicker for me as I had previous code
//
// Copyright © 2017 Conor O'Neill, conor@conoroneill.com
// License MIT

package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/feeds"
	"github.com/mmcdole/gofeed"
)

func running(w http.ResponseWriter, r *http.Request) {
	genFeed(w, "https://www.reddit.com/r/running/")
}

func trailrunning(w http.ResponseWriter, r *http.Request) {
	genFeed(w, "https://www.reddit.com/r/trailrunning/")
}

func genFeed(w http.ResponseWriter, feedURL string) {

	fp := gofeed.NewParser()
	inputFeed, err := fp.ParseURL(feedURL + ".rss")

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(inputFeed.Title)

	var RSSXML = &feeds.Feed{
		Title:       inputFeed.Title,
		Link:        &feeds.Link{Href: inputFeed.Link},
		Description: "Conors Proxy of " + inputFeed.Description,
		Author:      &feeds.Author{Name: "Conor", Email: "conor@conoroneill.com"},
	}

	for _, inputItem := range inputFeed.Items {

		layOut := "2006-01-02T15:04:05-07:00"
		timeStamp, err := time.Parse(layOut, inputItem.Updated)

		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		outputItem := feeds.Item{
			Title:       inputItem.Title,
			Link:        &feeds.Link{Href: inputItem.Link},
			Description: inputItem.Content,
			Author:      &feeds.Author{Name: "conor@conoroneill.com", Email: "conor@conoroneill.com"},
			Created:     timeStamp,
		}
		RSSXML.Add(&outputItem)

	}
	rss, err := RSSXML.ToRss()

	io.WriteString(w, rss)
}

func main() {
	http.HandleFunc("/r/running", running)
	http.HandleFunc("/r/trailrunning", trailrunning)
	http.ListenAndServe(":8111", nil)
}
