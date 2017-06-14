package main

import "fmt"
import "github.com/pushc6/testbot/redditclient"

func main() {
	fmt.Println("a")
	r, _ := redditclient.NewRedditClient("redditclient/redditclient.xml")

	//TODO: https://github.com/reddit/reddit/wiki/JSON
	//Need to build the basic reddit types as JSON structs for parsing and
	//using in requests. First get the basic GET /r/cincinnati thread working
	/*res, _ := r.MakeAPICall("https://oauth.reddit.com/r/cincinnati/new", "GET", nil)

	body := string(res)

	fmt.Println(body)
	*/
	fmt.Println("*************")
	params := make(map[string]string)
	params["limit"] = "1"
	params["raw_json"] = "1"
	res2, _ := r.MakeParsedAPICall("https://oauth.reddit.com/r/test/new?limit=1&raw_json=1", "GET", params, nil)

	//Still needs work parsing response
	fmt.Println(res2)
}
