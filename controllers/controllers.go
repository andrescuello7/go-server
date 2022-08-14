package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/andrescuello7/go-server/models"
)

var tasks []models.Task

func GetController(w http.ResponseWriter, r *http.Request) {
	tasks = append(tasks, models.Task{
		Id:    1,
		Enum:  0,
		Task:  "arreglar",
		Title: "Piesa sin foco",
	})
	json.NewEncoder(w).Encode(tasks)
}
