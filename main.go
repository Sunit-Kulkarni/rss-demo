package main

import (
	"encoding/xml"
	"fmt"
	"github.com/eduncan911/podcast"
	"time"
)

func main() {
	fmt.Println("Testing here")
	rightNow := time.Now()
	p := podcast.New("We are the champions", "patriots.win", "Awww yeah", &rightNow, &rightNow)

	// Making a new item (episode) for the podcast
	exampleEpisode := podcast.Item{
		XMLName:            xml.Name{},
		GUID:               "",
		Title:              "The Podcasters Best Pal",
		Link:               "PodcasterForEver.com",
		Description:        "Peace and Love",
		Author:             nil,
		AuthorFormatted:    "",
		Category:           "CoExistence",
		Comments:           "",
		Source:             "",
		PubDate:            nil,
		PubDateFormatted:   "",
		Enclosure:          nil,
		IAuthor:            "",
		ISubtitle:          "",
		ISummary:           nil,
		IImage:             nil,
		IDuration:          "",
		IExplicit:          "",
		IIsClosedCaptioned: "",
		IOrder:             "",
	}
	_, err := p.AddItem(exampleEpisode)
	fmt.Println("the errors are", err)
	out, _ := xml.MarshalIndent(p, " ", "  ")
	fmt.Println(string(out))
}
