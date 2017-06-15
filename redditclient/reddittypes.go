package redditclient

import (
	"encoding/json"
)

//Thing - Root of all the API call responses from Reddit
type Thing struct {
	ID   string          `json:"id"`
	Name string          `json:"name"`
	Kind string          `json:"kind"`
	Data json.RawMessage `json:"data"`
}

//Listing - Used for pagination of list items in reddit responses
type Listing struct {
	Before  string  `json:"before"`
	After   string  `json:"after"`
	ModHash string  `json:"modhash"`
	Things  []Thing `json:"children"`
}

//Link - A link (Reddit Thread) containing meta-data about a post
type Link struct {
	Title string `json:"title"`
}
