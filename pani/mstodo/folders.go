package mstodo

import (
	"encoding/json"
	"fmt"
	"net/http"
)

//Folder Task folder
type Folder struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	IsDefault bool   `json:"isDefaultFolder"`
}

//Folders Task folders from api
type Folders struct {
	Client   *http.Client `json:"-"`
	Value    []Folder     `json:"value"`
	NextLink string       `json:"@odata.nextLink"`
}

//List lists all the folder for the user
func (f Folders) List() Folders {
	// log.Println("Calling url - ", URLS["ListFolders"].url)
	resp, err := f.Client.Get(URLS["ListFolders"].url)
	if err != nil {
		panic(err)
	}
	// log.Println("Response Status - ", resp.Status)
	defer resp.Body.Close()
	err = json.NewDecoder(resp.Body).Decode(&f)
	if err != nil {
		panic(err)
	}
	return f
}

//Print prints the folders struct to console
func (f Folders) Print(detailed bool) {
	var err error
	for i, folder := range f.Value {
		defaultFolder := ""
		if folder.IsDefault {
			defaultFolder = "(Default)"
		}

		if detailed {
			_, err = fmt.Printf("%d - %s %s\n\tID: %s\n\n",
				i+1, folder.Name, defaultFolder, folder.ID)
		} else {
			_, err = fmt.Printf("%d - %s %s\n", i+1, folder.Name, defaultFolder)
		}

		if err != nil {
			panic(err)
		}
	}
}

//GetDefaultFolder loops over and get the default folders
//TODO: Cache the default folder id
func (f Folders) GetDefaultFolder() *Folder {
	for _, folder := range f.Value {
		if folder.IsDefault {
			return &folder
		}
	}
	return nil
}

func (f Folder) GetTasks() Tasks {
	return Tasks{}
}
