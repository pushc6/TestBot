package main

import "fmt"
import "github.com/pushc6/testbot/redditclient"

func main() {
	fmt.Println("a")
	r := redditclient.NewRedditClient()

	fmt.Println(r)
}
