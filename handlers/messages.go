package handlers

import (
	"GREEN-API/api"
	"GREEN-API/internal/logger"
	"encoding/json"
	"net/http"
)

// SendMessage отправляет сообщение в чат
func SendMessage(w http.ResponseWriter, r *http.Request) {
	idInstance := r.URL.Query().Get("idInstance")
	apiToken := r.URL.Query().Get("apiToken")
	chatID := r.URL.Query().Get("chatId")
	message := r.URL.Query().Get("message")

	// Валидация параметров
	if !ValidateCredentials(idInstance, apiToken) {
		logger.Warn("SendMessage: missing credentials")
		http.Error(w, "Missing idInstance or apiToken", http.StatusBadRequest)
		return
	}

	if chatID == "" || message == "" {
		logger.Warn("SendMessage: missing chatId or message")
		http.Error(w, "Missing chatId or message", http.StatusBadRequest)
		return
	}

	logger.Info("SendMessage: instance=%s, chatId=%s, messageLen=%d",
		idInstance, chatID, len(message))

	client := api.NewClient()
	result, err := client.SendMessage(idInstance, apiToken, chatID, message)
	if err != nil {
		logger.Error("SendMessage error: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	logger.Info("SendMessage success: idMessage=%s", result.IDMessage)

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(result); err != nil {
		logger.Error("Failed to encode response: %v", err)
	}
}
