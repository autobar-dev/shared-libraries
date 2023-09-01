package emailrepository

// Send email
type ServiceSendRequestBodyMessage struct {
	Plain string `json:"plain"`
	Html  string `json:"html"`
}

type ServiceSendRequestBody struct {
	From    string                        `json:"from"`
	To      string                        `json:"to"`
	Subject string                        `json:"subject"`
	Message ServiceSendRequestBodyMessage `json:"message"`
}

type ServiceSendResponse struct {
	Status string  `json:"status"`
	Error  *string `json:"error"`
}
