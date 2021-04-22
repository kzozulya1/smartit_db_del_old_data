package types

// JSONResponse is a common response
type JSONResponse struct {
	Error   bool   `json:"error"`
	Message string `json:"message"`
}

// NewJSONResponse constructs new response
func NewJSONResponse(e bool, m string) JSONResponse {
	return JSONResponse{
		Error:   e,
		Message: m,
	}
}
