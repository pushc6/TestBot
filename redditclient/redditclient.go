package redditclient

import (
	"encoding/json"
	"encoding/xml"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

//RedditClient - Handles authorization and requests to reddit API
type RedditClient struct {
	token *tokenResponse
}

type tokenResponse struct {
	AccessToken string  `json:"access_token"`
	TokenType   string  `json:"token_type"`
	ExpiresIn   float64 `json:"expires_in"`
	Scope       string  `json:"scope"`
}

type configClient struct {
	Username  string `xml:"username"`
	Password  string `xml:"password"`
	ClientID  string `xml:"clientID"`
	Secret    string `xml:"secret"`
	UserAgent string `xml:"userAgent"`
	AuthURL   string `xml:"authURL"`
}

//NewRedditClient - Creates a new reddit client instance with a valid token
func NewRedditClient() *RedditClient {
	client := &RedditClient{}
	client.token = client.getClientToken()
	fmt.Println(client.token.AccessToken)
	return client
}

func (r RedditClient) getClientToken() *tokenResponse {

	if r.token != nil {
		return r.token
	}

	configFile, err := ioutil.ReadFile("redditclient.xml")
	if err != nil {
		log.Fatal("Couldn't read client configuration file.")
	}

	config := &configClient{}
	xml.Unmarshal(configFile, config)

	user := config.Username
	password := config.Password
	secret := config.Secret
	clientID := config.ClientID
	userAgent := config.UserAgent
	authURL := config.AuthURL

	//Build call to get token
	client := &http.Client{}
	body := strings.NewReader("grant_type=password&username=" + user + "&password=" + password)
	req, _ := http.NewRequest("POST", authURL, body)
	req.Header.Set("User-Agent", userAgent)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.SetBasicAuth(clientID, secret)

	//Call authorization
	resp, _ := client.Do(req)

	//Parse response JSON
	bodyVal, _ := ioutil.ReadAll(resp.Body)
	tr := tokenResponse{}
	json.Unmarshal(bodyVal, &tr)

	return &tr

}

//MakeAPICall - Calls a Reddit API with the given method POST\GET and returns
// a response
func (r RedditClient) MakeAPICall(api, method string, request io.Reader) ([]byte, error) {
	req, err2 := r.buildRequest(api, method, request)
	if err2 != nil {
		log.Print("Failed building request for API call")
		return nil, errors.New("Failed building request for API call")
	}
	client := &http.Client{}

	resp, err := client.Do(req)

	if err != nil {
		log.Print("API Call Failed: ", err)
		return nil, errors.New("API Call Failed: ")
	}

	bodyVal, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("Failed reading API call response")
		return nil, errors.New("Failed reading API call response")
	}
	return bodyVal, nil
}

func (r RedditClient) buildRequest(apiURL, method string, payload io.Reader) (*http.Request, error) {
	userAgent := "windows:golang.reddit.bot.TestBot:.1 (by /u/realityman_"
	req, err := http.NewRequest(method, apiURL, payload)
	if err != nil {
		log.Fatal("There was a problem building the request")
		return nil, errors.New("There was a problem building the request")
	}
	req.Header.Set("User-Agent", userAgent)
	req.Header.Set("Authorization", "bearer "+r.getClientToken().AccessToken)
	return req, nil
}
