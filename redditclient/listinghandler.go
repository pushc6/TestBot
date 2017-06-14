package redditclient

type ListingHandler struct {
}

type thing struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Kind string `json:"kind"`
	Data string `json:"data"`
}

type listing struct {
	Before  string  `json:"before"`
	After   string  `json:"after"`
	ModHash string  `json:"modhash"`
	Thing   []thing `json:"children"`
}

type response struct {
	Listing *listing
}

func NewListingHandler() *ListingHandler {
	return &ListingHandler{}
}
