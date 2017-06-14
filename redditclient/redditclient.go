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
	token  *tokenResponse
	config *configClient
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
func NewRedditClient(filePath string) (*RedditClient, error) {
	client := &RedditClient{}

	configFile, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Println("There was a problem loading the configuration file.")
		return nil, errors.New("couldn't load reddit client configuration file")
	}

	client.config = &configClient{}
	xml.Unmarshal(configFile, client.config)

	client.token = client.getClientToken()
	fmt.Println(client.token.AccessToken)

	return client, nil
}

func (r RedditClient) getClientToken() *tokenResponse {

	if r.token != nil {
		return r.token
	}

	//Build call to get token
	client := &http.Client{}
	body := strings.NewReader("grant_type=password&username=" + r.config.Username + "&password=" + r.config.Password)
	req, _ := http.NewRequest("POST", r.config.AuthURL, body)
	req.Header.Set("User-Agent", r.config.UserAgent)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.SetBasicAuth(r.config.ClientID, r.config.Secret)

	//Call authorization
	resp, _ := client.Do(req)

	//Parse response JSON
	bodyVal, _ := ioutil.ReadAll(resp.Body)
	tr := tokenResponse{}
	json.Unmarshal(bodyVal, &tr)

	return &tr

}

//MakeAPICall - Calls a Reddit API with the given method POST\GET and returns a response
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
	userAgent := r.config.UserAgent
	req, err := http.NewRequest(method, apiURL, payload)
	if err != nil {
		log.Fatal("There was a problem building the request")
		return nil, errors.New("There was a problem building the request")
	}
	req.Header.Set("User-Agent", userAgent)
	fmt.Println("setting user agent to: ", userAgent)
	req.Header.Set("Authorization", "bearer "+r.getClientToken().AccessToken)
	return req, nil
}
