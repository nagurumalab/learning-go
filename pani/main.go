package main

import (
	"fmt"

	"github.com/nagurumalab/learning-go/pani/mstodo"
)

// https://godoc.org/golang.org/x/oauth2#ex-Config--CustomHTTP
func main() {

	client := mstodo.NewClient()

	folders := client.ListFolders()
	for i, folder := range folders {
		fmt.Printf("%d - %s\n", i, folder.Name)
	}
	//fmt.Println(folders)
	// jd, _ := json.MarshalIndent(jdata, "", "")
	// fmt.Println(string(jd))
}
