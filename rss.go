package main

import (
	"context"
	"fmt"
	"net/http"
	"log"
	"io"
	"encoding/xml"
	"html"
)


type RSSFeed struct {
	Channel struct {
		Title       string    `xml:"title"`
		Link        string    `xml:"link"`
		Description string    `xml:"description"`
		Item        []RSSItem `xml:"item"`
	} `xml:"channel"`
}

type RSSItem struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Description string `xml:"description"`
	PubDate     string `xml:"pubDate"`
}
//func fetchFeed(ctx context.Context, feedURL string) (*RSSFeed, error) {
func fetchFeed(ctx context.Context, feedURL string) {
	// Create a context that can be canceled


	req, err := http.NewRequestWithContext(ctx, "GET", feedURL, nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}
	req.Header.Set("User-Agent", "gator")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	b, err := io.ReadAll(resp.Body)
	v := RSSFeed{}
	err = xml.Unmarshal(b, &v)
	if err != nil {
		fmt.Printf("error: %v", err)
		return
	}
	fmt.Printf("Title: %#v\n", html.UnescapeString(v.Channel.Title))
	fmt.Printf("Link: %#v\n", v.Channel.Link)
	fmt.Printf("Description: %#v\n", v.Channel.Description)
	for _, item := range v.Channel.Item {
		fmt.Printf("Title: %#v\n", html.UnescapeString(item.Title))
		fmt.Printf("Title: %#v\n", html.UnescapeString(item.Description))
	}
	//	fmt.Printf("%s\n", b)

}
