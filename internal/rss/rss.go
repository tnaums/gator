package rss

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

func main() {
	// Create a context that can be canceled
	ctx, _ := context.WithCancel(context.Background())

	req, err := http.NewRequestWithContext(ctx, "GET", "https://www.wagslane.dev/index.xml", nil)
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
