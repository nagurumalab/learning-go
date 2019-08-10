package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/microsoft"
)

// https://godoc.org/golang.org/x/oauth2#ex-Config--CustomHTTP
func main() {
	ctx := context.Background()

	conf := &oauth2.Config{
		ClientID: "c7bcff77-9645-458a-a39e-f5d3c2664ea8",
		//		ClientSecret: "DBM5A/pYGOzfD/Atodq*?oZi2uHT6hL0",
		Scopes:      []string{"Tasks.Read", "Tasks.Read.Shared", "Tasks.ReadWrite", "Tasks.ReadWrite.Shared", "User.Read"},
		Endpoint:    microsoft.AzureADEndpoint("common"),
		RedirectURL: "https://login.microsoftonline.com/common/oauth2/nativeclient",
	}

	// Redirect user to consent page to ask for permission
	// for the scopes specified above.
	url := conf.AuthCodeURL("state", oauth2.AccessTypeOffline)
	fmt.Printf("Visit the URL for the auth dialog: %v", url)

	// Use the authorization code that is pushed to the redirect
	// URL. Exchange will do the handshake to retrieve the
	// initial access token. The HTTP Client returned by
	// conf.Client will refresh the token as necessary.
	var code string
	if _, err := fmt.Scan(&code); err != nil {
		log.Fatal(err)
	}

	// Use the custom HTTP client when requesting a token.
	httpClient := &http.Client{Timeout: 2 * time.Second}
	ctx = context.WithValue(ctx, oauth2.HTTPClient, httpClient)

	tok, err := conf.Exchange(ctx, code)
	if err != nil {
		log.Fatal(err)
	}

	client := conf.Client(ctx, tok)
	resp, err := client.Get("https://graph.microsoft.com/beta/me/outlook/taskFolders")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	var jsonBody map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&jsonBody)
	fmt.Println(json.MarshalIndent(jsonBody, "", ""))
}
