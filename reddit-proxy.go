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

	"github.com/SlyMarbo/rss"
	"github.com/gorilla/feeds"
)

func running(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Got to 15")

	genFeed(w, "https://www.reddit.com/r/running/")

	fmt.Println("Got to 16")

}

func trailrunning(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Got to 17")

	genFeed(w, "https://www.reddit.com/r/trailrunning/")

	fmt.Println("Got to 18")

}

func genFeed(w http.ResponseWriter, feedURL string) {

	fmt.Println("Got to 5")

	inputFeed, err := rss.Fetch(feedURL + ".rss")

	fmt.Println("Got to 6")

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

	fmt.Println("Got to 7")

	for _, inputItem := range inputFeed.Items {

		fmt.Println("Got to 8")

		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		outputItem := feeds.Item{
			Title:       inputItem.Title,
			Link:        &feeds.Link{Href: inputItem.Link},
			Description: inputItem.Content,
			Author:      &feeds.Author{Name: "conor", Email: "conor@conoroneill.com"},
			Created:     inputItem.Date,
		}

		fmt.Println("Got to 9")

		fmt.Println("Got to 10")

		RSSXML.Add(&outputItem)

		fmt.Println("Got to 11")

	}

	fmt.Println("Got to 12")

	rss, err := RSSXML.ToRss()

	fmt.Println("Got to 13")

	io.WriteString(w, rss)

	fmt.Println("Got to 14")

}

func main() {
	//TODO: Just change the sub-reddit name to a query param so it's fully dynamic
	//TODO: Figure out why genfeed seems to be called twice for every browser request
	//TODO: Figure out why it's so spectacularly slow on EC2
	fmt.Println("Got to 1")
	http.HandleFunc("/r/running", running)
	fmt.Println("Got to 2")
	http.HandleFunc("/r/trailrunning", trailrunning)
	fmt.Println("Got to 3")
	http.ListenAndServe(":8111", nil)
	fmt.Println("Got to 4")
}
