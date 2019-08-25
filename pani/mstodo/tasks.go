package mstodo

//Task struct
type Task struct {
	id      string
	Subject string
	Body    struct {
		contentType string
		content     string
	}
	CreatedDateTime string
	DueDateTime     string
	IsReminderOn    bool
	parentFolderId  string
	Status          string
}

//Tasks list of tasks
type Tasks struct {
	Tasks    []Task `json:value`
	nextLink string `json:"@odata.nextLink"`
}

//ListTasks lists tasks
func (c Client) ListTasks(listID string) Tasks {
	if listID != "" {
		listID = ""
	}
	return Tasks{}
}
