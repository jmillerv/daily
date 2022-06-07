package helpers

import (
	"log"
	"net/url"
)

func ParseURL(urlStr string) *url.URL {
	link, err := url.Parse(urlStr)
	if err != nil {
		log.Println("Could not parse URL", err)
	}

	return link
}
