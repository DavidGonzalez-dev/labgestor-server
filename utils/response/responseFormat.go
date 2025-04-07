package response

// Estructura estandarizada para la respuesta de las APIS
type Response struct {
	Message string `json:"message,omitempty"`
	Error   string `json:"error,omitempty"`
	Body    any    `json:"body,omitempty"`
}
