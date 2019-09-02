package mstodo

//Task struct
type Task struct {
	ID      string
	Subject string
	Body    struct {
		contentType string
		content     string
	}
	CreatedDateTime string
	DueDateTime     string
	IsReminderOn    bool
	ParentFolderID  string
	Status          string
}

//Tasks list of tasks
type Tasks struct {
	Tasks    []Task `json:"value"`
	NextLink string `json:"@odata.nextLink"`
}

//ListTasks lists tasks
func (t Tasks) ListTasks(listID string) []Task {
	if listID != "" {
		listID = ""
	}
	return Tasks{}.Tasks
}
