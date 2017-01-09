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

	"github.com/SlyMarbo/rss"
	"github.com/gorilla/feeds"
)

func rRunning(w http.ResponseWriter, r *http.Request) {
	genFeed(w, "https://www.reddit.com/r/running/")
}

func rTrailRunning(w http.ResponseWriter, r *http.Request) {
	genFeed(w, "https://www.reddit.com/r/trailrunning/")
}

func everythingElse(w http.ResponseWriter, r *http.Request) {
	subReddit := r.URL.Query().Get("r")

	genFeed(w, "https://www.reddit.com/r/"+subReddit)
}

func genFeed(w http.ResponseWriter, feedURL string) {

	inputFeed, err := rss.Fetch(feedURL + ".rss")

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(inputFeed.Title)

	var RSSXML = &feeds.Feed{
		Title:       inputFeed.Title,
		Link:        &feeds.Link{Href: inputFeed.Link},
		Description: "Conors Proxy of " + inputFeed.Description,
		Author:      &feeds.Author{Name: "Conor", Email: "conor@conoroneill.com"},
	}

	for _, inputItem := range inputFeed.Items {

		if err != nil {
			fmt.Println(err)
		}
		outputItem := feeds.Item{
			Title:       inputItem.Title,
			Link:        &feeds.Link{Href: inputItem.Link},
			Description: inputItem.Content,
			Author:      &feeds.Author{Name: "conor", Email: "conor@conoroneill.com"},
			Created:     inputItem.Date,
		}

		RSSXML.Add(&outputItem)

	}

	rss, err := RSSXML.ToRss()

	io.WriteString(w, rss)

}

func main() {
	//TODO: Better error handling
	http.HandleFunc("/r/running", rRunning)
	http.HandleFunc("/r/trailrunning", rTrailRunning)
	http.HandleFunc("/", everythingElse)
	http.ListenAndServe(":8111", nil)
}
