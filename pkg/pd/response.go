package pd

import (
	"fmt"
)

type ResponseDefault struct {
	Success bool   `json:"success"`
	Value   string `json:"value,omitempty"`
	Message string `json:"message,omitempty"`
}

type ResponseUpload struct {
	StatusCode int    `json:"code"`
	ID         string `json:"id,omitempty"`
	ResponseDefault
}

type ResponseDownload struct {
	StatusCode int    `json:"code"`
	FilePath   string `json:"file_path"`
	FileName   string `json:"file_name"`
	FileSize   int64  `json:"file_size"`
	ResponseDefault
}

// GetFileURL return the full URl to the uploaded file
func (rsp *ResponseUpload) GetFileURL() string {
	return fmt.Sprintf("%su/%s", BaseURL, rsp.ID)
}
