package models

type Task struct {
	Id    int    `json:"id,omitempty"`
	Enum  int    `json:"enum,omitempty"`
	Task  string `json:"task,omitempty"`
	Title string `json:"title,omitempty"`
}
