package mstodo

import (
	"encoding/json"
	"fmt"
)

//Folder Task folder
type Folder struct {
	id        string `json:"id"`
	Name      string `json:"name"`
	IsDefault bool   `json:"isDefaultFolder"`
}

//Folders Task folders from api
type Folders struct {
	Folders  []Folder `json:"value"`
	nextLink string   `json:"@odata.nextLink"`
}

//ListFolders lists all the folder for the user
func (c Client) ListFolders() Folders {
	// jdata := c.callAPI(URLS["ListFolders"].method, URLS["ListFolders"].url, nil)
	// jd := *interface{}(jdata).(*map[string]interface{})
	// folders := []Folder{}
	// for _, f := range jd["value"].([]interface{}) {
	// 	value := f.(map[string]interface{})
	// 	folders = append(folders, Folder{
	// 		id:        value["id"].(string),
	// 		Name:      value["name"].(string),
	// 		IsDefault: value["isDefaultFolder"].(bool),
	// 	})
	// }
	resp := c.callAPI(URLS["ListFolders"].method, URLS["ListFolders"].url, nil)
	var folders = Folders{}

	defer resp.Body.Close()
	err := json.NewDecoder(resp.Body).Decode(&folders)
	if err != nil {
		panic(err)
	}
	return folders
}

//Print prints the folders struct to console
func (f Folders) Print(detailed bool) {
	var err error
	for i, folder := range f.Folders {
		if detailed {

			_, err = fmt.Printf("%d - %s\nDefaut: %t ID: %s\n\n",
				i+1, folder.Name, folder.IsDefault, folder.id)
		} else {
			_, err = fmt.Printf("%d - %s\n", i+1, folder.Name)
		}

		if err != nil {
			panic(err)
		}
	}
}

//GetDefaultFolder loops over and get the default folders
//TODO: Cache the default folder id
func (f Folders) GetDefaultFolder() *Folder {
	for _, folder := range f.Folders {
		if folder.IsDefault {
			return &folder
		}
	}
	return nil
}

func (f Folder) GetTasks() Tasks {
	return Tasks{}
}
