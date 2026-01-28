package handlers

import (
	"GREEN-API/api"
	"GREEN-API/internal/logger"
	"encoding/json"
	"net/http"
)

// ValidateCredentials проверяет наличие обязательных учетных данных
func ValidateCredentials(idInstance, apiToken string) bool {
	return idInstance != "" && apiToken != ""
}

// GetSettings возвращает настройки инстанса
func GetSettings(w http.ResponseWriter, r *http.Request) {
	idInstance := r.URL.Query().Get("idInstance")
	apiToken := r.URL.Query().Get("apiToken")

	if !ValidateCredentials(idInstance, apiToken) {
		logger.Warn("GetSettings: missing credentials")
		http.Error(w, "Missing idInstance or apiToken", http.StatusBadRequest)
		return
	}

	logger.Info("GetSettings called for instance: %s", idInstance)

	client := api.NewClient()
	result, err := client.GetSettings(idInstance, apiToken)
	if err != nil {
		logger.Error("GetSettings error: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(result); err != nil {
		logger.Error("Failed to encode response: %v", err)
	}
}

// GetStateInstance возвращает состояние инстанса
func GetStateInstance(w http.ResponseWriter, r *http.Request) {
	idInstance := r.URL.Query().Get("idInstance")
	apiToken := r.URL.Query().Get("apiToken")

	if !ValidateCredentials(idInstance, apiToken) {
		logger.Warn("GetStateInstance: missing credentials")
		http.Error(w, "Missing idInstance or apiToken", http.StatusBadRequest)
		return
	}

	logger.Info("GetStateInstance called for instance: %s", idInstance)

	client := api.NewClient()
	result, err := client.GetStateInstance(idInstance, apiToken)
	if err != nil {
		logger.Error("GetStateInstance error: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(result); err != nil {
		logger.Error("Failed to encode response: %v", err)
	}
}
