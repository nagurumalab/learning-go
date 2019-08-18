package main

import (
	"encoding/json"
	"fmt"
)

// https://godoc.org/golang.org/x/oauth2#ex-Config--CustomHTTP
func main() {

	resp, err := client.Get("https://graph.microsoft.com/beta/me/outlook/taskFolders")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	var taskFolders map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&taskFolders)
	jd, err := json.MarshalIndent(taskFolders, "", "")
	fmt.Println(string(jd))

}
