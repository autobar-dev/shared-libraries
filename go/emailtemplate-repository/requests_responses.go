package emailtemplaterepository

// Render email template
type ServiceRenderTemplateRequestBody struct {
	TemplateName    string                 `json:"template_name"`
	TemplateVersion *string                `json:"template_version"`
	Locale          string                 `json:"locale"`
	Params          map[string]interface{} `json:"params"`
}

type ServiceRenderTemplateResponse struct {
	Status string            `json:"status"`
	Error  *string           `json:"error"`
	Data   *RenderedTemplate `json:"data"`
}
