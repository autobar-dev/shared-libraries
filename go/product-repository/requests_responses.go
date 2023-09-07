package productrepository

// Get product
type ServiceGetProductResponseData struct {
	Type         ProductResponseType `json:"type"`
	Product      *Product            `json:"product"`
	RedirectSlug *string             `json:"redirect_slug"`
}

type ServiceGetProductResponse struct {
	Status string                         `json:"status"`
	Error  *string                        `json:"error"`
	Data   *ServiceGetProductResponseData `json:"data"`
}
