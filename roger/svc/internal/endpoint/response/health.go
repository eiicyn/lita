package response

type HealthCheckResponse struct {
	RogerResponse
	Data string `json:"data"`
}
