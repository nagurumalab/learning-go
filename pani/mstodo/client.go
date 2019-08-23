package mstodo

import (
	"context"
	"net/http"
	"time"

	"golang.org/x/oauth2"
)

//Client struct to hold the client deals with api calls
type Client struct {
	client *http.Client
}

//NewClient will init and give you back a todo client to work with
func NewClient() Client {
	ctx := context.Background()
	httpClient := &http.Client{Timeout: 2 * time.Second}
	ctx = context.WithValue(ctx, oauth2.HTTPClient, httpClient)
	//TODO: Error handling code
	config := getOauthConfig()
	client, _ := userOAuthClient(ctx, config)
	return Client{client: client}
}

func (c *Client) callAPI(method string, url string, payload map[string]interface{}) *http.Response {
	var resp *http.Response
	var err error
	switch method {
	case "GET":
		// log.Printf("Calling URL - %s\n", url)
		resp, err = c.client.Get(url)
		// log.Printf("Response - %s\n", resp.Status)
	case "POST":
		resp, err = c.client.Get(url)
	}
	if err != nil {
		panic(err)
	}
	return resp
}
