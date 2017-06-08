package main

import "encoding/json"
import "fmt"

type TokenResponse struct {
	Access_token string  `json:"access_token"`
	Token_types  string  `json:"token_type"`
	Expires_in   float64 `json:"expires_in"`
	Scope        string  `json:"scope"`
}

func stuff() {
	tr := TokenResponse{}
	tr2 := TokenResponse{"blah", "bearer", 4000, "*"}
	jsonSt := "{\"access_token\": \"n0cIgfYK6yBYSg1k2zjd-1UO0sQ\", \"token_type\": \"bearer\", \"expires_in\": 3600, \"scope\": \"*\"}"

	err := json.Unmarshal([]byte(jsonSt), &tr)
	ress, err2 := json.Marshal(tr2)
	fmt.Println(string(ress), err2, tr2.Access_token)
	fmt.Println("ERROR? ", err)
	fmt.Println(tr.Scope)

}
