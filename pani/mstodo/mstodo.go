package mstodo

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"golang.org/x/oauth2"
)

//Task struct
type Task struct {
	title string
	body  string
}

//Folder Task folders
type Folder struct {
	id        string
	Name      string
	IsDefault bool
}

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

func (c *Client) callAPI(method string, url string, payload map[string]interface{}) *map[string]interface{} {
	var resp *http.Response
	var err error
	var jdata map[string]interface{}
	switch method {
	case "GET":
		resp, err = c.client.Get(url)
	case "POST":
		resp, err = c.client.Get(url)
	}
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	json.NewDecoder(resp.Body).Decode(&jdata)
	return &jdata
}

//ListFolders lists all the folder for the user
func (c Client) ListFolders() []Folder {
	jdata := c.callAPI(URLS["ListFolders"].method, URLS["ListFolders"].url, nil)
	jd := *interface{}(jdata).(*map[string]interface{})
	folders := []Folder{}
	for _, f := range jd["value"].([]interface{}) {
		value := f.(map[string]interface{})
		folders = append(folders, Folder{
			id:        value["id"].(string),
			Name:      value["name"].(string),
			IsDefault: value["isDefaultFolder"].(bool),
		})
	}
	return folders
}
