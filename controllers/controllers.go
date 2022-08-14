package controllers

import (
	"encoding/json"
	"net/http"
)

type Task struct {
	Id    int    `json:"Id,omitempty"`
	Enum  int    `json:"Enum,omitempty"`
	Task  string `json:"Task,omitempty"`
	Title string `json:"Title,omitempty"`
}

var tasks []Task

func GetController(w http.ResponseWriter, r *http.Request) {
	tasks = append(tasks, Task{
		Id:    1,
		Enum:  0,
		Task:  "arreglar",
		Title: "Piesa sin foco",
	})
	json.NewEncoder(w).Encode(tasks)
}
