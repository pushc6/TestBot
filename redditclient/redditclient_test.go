package redditclient

import "testing"

func TestNewRedditClient(t *testing.T) {

	r, err := NewRedditClient("redditclient.xml")
	if err != nil {
		t.Error("Should have been able to load reddit client, but couldn't.")
	}

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

	r2, err2 := NewRedditClient("thisisascrewedupfile.xml")
	if err2 == nil {
		t.Error("Shouldn't have been able to load reddit client, but did.")
	}
	if r2 != nil {
		t.Error("Shouldn't have been able to load reddit client, but did.")
	}

}

func TestMakeAPICall(t *testing.T) {
	r, err := NewRedditClient("redditclient.xml")
	if err != nil {
		t.Error("Should have been able to load reddit client, but couldn't.")
	}
	//Test Call With Wrong API
	_, err2 := r.MakeAPICall("sdfsfsdfsfsd", "test", nil)

	if err2 == nil {
		t.Error("Expected an error response for bad API call")
	}

}
