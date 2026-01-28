package handlers

import (
	"GREEN-API/api"
	"GREEN-API/internal/logger"
	"encoding/json"
	"net/http"
)

// SendFileByURL отправляет файл по URL в чат
func SendFileByURL(w http.ResponseWriter, r *http.Request) {
	idInstance := r.URL.Query().Get("idInstance")
	apiToken := r.URL.Query().Get("apiToken")
	chatID := r.URL.Query().Get("chatId")
	fileURL := r.URL.Query().Get("fileUrl")
	fileName := r.URL.Query().Get("fileName")
	caption := r.URL.Query().Get("caption")

	// Валидация параметров
	if !ValidateCredentials(idInstance, apiToken) {
		logger.Warn("SendFileByURL: missing credentials")
		http.Error(w, "Missing idInstance or apiToken", http.StatusBadRequest)
		return
	}

	if chatID == "" || fileURL == "" {
		logger.Warn("SendFileByURL: missing chatId or fileUrl")
		http.Error(w, "Missing chatId or fileUrl", http.StatusBadRequest)
		return
	}

	// Если fileName не указан, используем последнюю часть URL
	if fileName == "" {
		fileName = "file"
	}

	logger.Info("SendFileByURL: instance=%s, chatId=%s, fileUrl=%s, fileName=%s",
		idInstance, chatID, fileURL, fileName)

	client := api.NewClient()
	result, err := client.SendFileByURL(idInstance, apiToken, chatID, fileURL, fileName, caption)
	if err != nil {
		logger.Error("SendFileByURL error: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	logger.Info("SendFileByURL success: idMessage=%s", result.IDMessage)

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(result); err != nil {
		logger.Error("Failed to encode response: %v", err)
	}
}
