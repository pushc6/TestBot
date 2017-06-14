package main

import "fmt"
import "github.com/pushc6/testbot/redditclient"

func main() {
	fmt.Println("a")
	r, _ := redditclient.NewRedditClient("redditclient/redditclient.xml")

	//TODO: https://github.com/reddit/reddit/wiki/JSON
	//Need to build the basic reddit types as JSON structs for parsing and
	//using in requests. First get the basic GET /r/cincinnati thread working
	res, _ := r.MakeAPICall("https://outh.reddit.com/r/cincinnati/new", "GET", nil)

	body := string(res)

	fmt.Println(body)
}
