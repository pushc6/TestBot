package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type TokenResponse struct {
	Access_Token string  `json:"access_token"`
	Token_Type   string  `json:"token_type"`
	Expires_In   float64 `json:"expires_in"`
	Scope        string  `json:"scope"`
}

func main() {
	user := ""
	password := ""
	secret := ""
	clientID := ""
	user_agent := "windows:golang.reddit.bot.TestBot:.1 (by /u/realityman_)"
	authUrl := "https://www.reddit.com/api/v1/access_token"

	//Build call to get token
	client := &http.Client{}
	body := strings.NewReader("grant_type=password&username=" + user + "&password=" + password)
	req, _ := http.NewRequest("POST", authUrl, body)
	req.Header.Set("User-Agent", user_agent)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.SetBasicAuth(clientID, secret)

	//Call authorization
	resp, _ := client.Do(req)

	//Parse response JSON
	bodyVal, _ := ioutil.ReadAll(resp.Body)
	tr := TokenResponse{}
	json.Unmarshal(bodyVal, &tr)

	//Let's try an API Call

	newURL := "https://oauth.reddit.com/message/inbox"
	req2, _ := http.NewRequest("GET", newURL, nil)
	req2.Header.Set("User-Agent", user_agent)
	req2.Header.Set("Authorization", "bearer "+tr.Access_Token)
	fmt.Println(" HEADER", req2.Header.Get("Authorization"))

	resp, _ = client.Do(req2)
	bod, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(bod))

}
