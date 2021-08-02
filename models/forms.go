package models

type OptionSet struct {
	Id    string `json:"id"`
	Type  string `json:"type"`
	Name  string `json:"name"`
	Value string `json:"value"`
	Changeable
	Creatable
	TeamId string `json:"teamId"`
}

type Team struct {
	Id     string   `json:"id"`
	Name   string   `json:"name"`
	Owner  string   `json:"owner"`
	Admins []string `json:"admins"`
	Changeable
	Creatable
}

type Form struct {
	Id          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Changeable
	Creatable
	Fields            []interface{} `json:"fields"`
	SubmissionCount   int64         `json:"submissionCount"`
	UnreadSubmissions int64         `json:"unreadSubmissions"`
	SubmissionConfig  interface{}   `json:"submissionConfig"`
	Versioned
	Folder   string   `json:"folder"`
	TeamId   string   `json:"teamId"`
	Workflow Workflow `json:"workflow"`
}

type FormField struct {
	Name  string `json:"name"`
	Label string `json:"label"`
	Type  string `json:"type"`
	Value string `json:"value"`
}

type Workflow struct {
	Integrations []Integration `json:"integrations"`
}

type Integration struct {
	Name       string            `json:"name"`
	Config     map[string]string `json:"config"`
	Submission Submission        `json:"submission"`
	Form       Form              `json:"form"`
	Enabled    bool              `json:"enabled"`
}

type Folder struct {
	Id   string `json:"id"`
	Name string `json:"name"`
	Changeable
	Creatable
	Parent string `json:"parent"`
	Versioned
	TeamId string `json:"teamId"`
}

type LeanForm struct {
	Id string
	Changeable
	Title string
}

type File struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	FileType string `json:"type"`
	Size     int64  `json:"size"`
	FormId   string `json:"formId"`
}

type LeanSubmission struct {
	Id string
	Creatable
	Status string
}

type SubmissionMeta struct {
	Env       interface{} `json:"env"`
	IpAddress string      `json:"ipAddress"`
}

type Submission struct {
	FormId    string         `json:"formId"`
	Id        string         `json:"id"`
	Details   interface{}    `json:"details"`
	Meta      SubmissionMeta `json:"meta"`
	FieldMeta interface{}    `json:"fieldMeta"`
	Status    string         `json:"status"`
	Processed bool           `json:"processed"`
	Creatable
	ReadBy map[string]bool `json:"readBy"`
}

type SubmissionsDeleted struct {
	FormId string   `json:"formId"`
	Ids    []string `json:"ids"`
}

type ContentBlock struct {
	Id   string `json:"id"`
	Name string `json:"name"`
	Changeable
	Creatable
	Value string `json:"value"`
}
