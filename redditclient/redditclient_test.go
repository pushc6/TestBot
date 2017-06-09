package redditclient

import "testing"

func TestNewRedditClient(t *testing.T) {

	r := NewRedditClient()

	if r.token == nil {
		t.Error("Didn't receieve a token in the generated client")
	} else if r.token.AccessToken == "" {
		t.Error("No access token received")
	} else if r.token.ExpiresIn == 0.0 {
		t.Error("No expiration received")
	} else if r.token.Scope == "" {
		t.Error("No scope received")
	} else if r.token.TokenType == "" {
		t.Error("No token type received")
	}

}

func TestMakeAPICall(t *testing.T) {
	r := NewRedditClient()

	//Test Call With Wrong API
	_, err := r.MakeAPICall("sdfsfsdfsfsd", "test", nil)

	if err == nil {
		t.Error("Expected an error response for bad API call")
	}

}
