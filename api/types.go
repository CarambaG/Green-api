package api

type SettingsResponse struct {
	WebhookURL                        string   `json:"webhookUrl"`
	WebhookURLToken                   string   `json:"webhookUrlToken"`
	DelaySendMessagesMilliseconds     int      `json:"delaySendMessagesMilliseconds"`
	MarkIncomingMessagesAsRead        bool     `json:"markIncomingMessagesAsRead"`
	MarkIncomingMessagesAsReadOnReply bool     `json:"markIncomingMessagesAsReadOnReply"`
	GroupsIgnoreMarkRecipientRead     bool     `json:"groupsIgnoreMarkRecipientRead"`
	SaveIncomingFiles                 bool     `json:"saveIncomingFiles"`
	AutoReplyOnMessages               bool     `json:"autoReplyOnMessages"`
	AutoReplyOnMessagesText           string   `json:"autoReplyOnMessagesText"`
	AutoReplyOnMessagesExceptKeywords []string `json:"autoReplyOnMessagesExceptKeywords"`
	StatusInstance                    string   `json:"statusInstance"`
	DisableGroupNotifications         bool     `json:"disableGroupNotifications"`
	DisableAllNotifications           bool     `json:"disableAllNotifications"`
}

type StateResponse struct {
	StateInstance string `json:"stateInstance"`
}

type MessageResponse struct {
	IDMessage string `json:"idMessage"`
}

type SendMessageRequest struct {
	ChatID  string `json:"chatId"`
	Message string `json:"message"`
}

type SendFileRequest struct {
	ChatID   string `json:"chatId"`
	FileURL  string `json:"urlFile"`
	FileName string `json:"fileName"`
	Caption  string `json:"caption"`
}

type APIError struct {
	Code    int    `json:"code"`
	Status  string `json:"status"`
	Message string `json:"message"`
}

type APIResponse struct {
	Result interface{} `json:"result"`
	Error  *APIError   `json:"error,omitempty"`
}
