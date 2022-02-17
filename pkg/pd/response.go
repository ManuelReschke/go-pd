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

func (rsp *ResponseUpload) GetFileURL() string {
	return fmt.Sprintf("%s/u/%s", BaseURL, rsp.ID)
}