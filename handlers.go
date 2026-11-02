package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type HealthResponse struct {
	Status string `json:"status"`
}

type InfoResponse struct {
	App     string `json:"app"`
	Version string `json:"version"`
	Status  string `json:"status"`
}

type Task struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Done  bool   `json:"done"`
}

var tasks = []Task{
	{ID: 1, Title: "یادگیری Go", Done: true},
	{ID: 2, Title: "نوشتن API تمیز", Done: false},
	{ID: 3, Title: "نوشتن API ", Done: false},
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprintln(w, welcomeMsg)
	if err != nil {
		return
	}
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, methodNotAllowedMsg, http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	resp := HealthResponse{
		Status: healthStatus,
	}

	if err := json.NewEncoder(w).Encode(resp); err != nil {
		http.Error(w, "خطا در ساخت پاسخ", http.StatusInternalServerError)
		return
	}
}

func infoHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, methodNotAllowedMsg, http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	resp := InfoResponse{
		App:     appName,
		Version: appVersion,
		Status:  appStatus,
	}

	if err := json.NewEncoder(w).Encode(resp); err != nil {
		http.Error(w, "خطا در ساخت پاسخ", http.StatusInternalServerError)
		return
	}
}

func tasksHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		getTasks(w, r)
	case http.MethodPost:
		createTask(w, r)
	default:
		http.Error(w, methodNotAllowedMsg, http.StatusMethodNotAllowed)
	}
}

func getTasks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(tasks); err != nil {
		http.Error(w, "خطا در ساخت پاسخ", http.StatusInternalServerError)
	}
}

func createTask(w http.ResponseWriter, r *http.Request) {
	var newTask Task

	if err := json.NewDecoder(r.Body).Decode(&newTask); err != nil {
		http.Error(w, "بدنه درخواست نامعتبر است", http.StatusBadRequest)
		return
	}

	newTask.ID = len(tasks) + 1
	tasks = append(tasks, newTask)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	if err := json.NewEncoder(w).Encode(newTask); err != nil {
		http.Error(w, "خطا در ساخت پاسخ", http.StatusInternalServerError)
	}
}

func doneHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, methodNotAllowedMsg, http.StatusMethodNotAllowed)
		return
	}

	var input struct {
		ID int `json:"id"`
	}

	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		http.Error(w, badRequest, http.StatusBadRequest)
		return
	}

	for i := range tasks {
		if tasks[i].ID == input.ID {
			tasks[i].Done = true

			w.Header().Set("Content-Type", "application/json")

			err = json.NewEncoder(w).Encode(tasks[i])
			if err != nil {
				http.Error(w, "خطا در ساخت پاسخ", http.StatusInternalServerError)
				return
			}
			return
		}
	}

	http.Error(w, notFound, http.StatusNotFound)
}

func countHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, methodNotAllowedMsg, http.StatusMethodNotAllowed)
		return
	}

	total := 0
	doneCount := 0

	for i := range tasks {
		total = total + 1
		if tasks[i].Done {
			doneCount = doneCount + 1
		}
	}

	type CountResponse struct {
		Total int `json:"total"`
		Done  int `json:"done"`
	}

	resp := CountResponse{
		Total: total,
		Done:  doneCount,
	}

	w.Header().Set("Content-Type", "application/json")

	err := json.NewEncoder(w).Encode(resp)
	if err != nil {
		http.Error(w, "خطا در ساخت پاسخ", http.StatusInternalServerError)
		return
	}
}
