package mstodo

type urlMethod struct {
	method string
	url    string
}

//URLS list of all the urls
var URLS = map[string]urlMethod{
	"ListFolders": {"GET", "https://graph.microsoft.com/beta/me/outlook/taskFolders"},
}
