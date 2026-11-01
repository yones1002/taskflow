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

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, welcomeMsg)
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
