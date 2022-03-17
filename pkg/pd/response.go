package pd

import (
	"fmt"
)

type ResponseUpload struct {
	StatusCode int    `json:"code"`
	ID         string `json:"id,omitempty"`
	Success    bool   `json:"success"`
	Value      string `json:"value,omitempty"`
	Message    string `json:"message,omitempty"`
}

// GetFileURL return the full URl to the uploaded file
func (rsp *ResponseUpload) GetFileURL() string {
	return fmt.Sprintf("%su/%s", BaseURL, rsp.ID)
}
