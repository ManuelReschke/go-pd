package pd

import (
	"fmt"
	"time"
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

// GetFileURL return the full URl to the uploaded file
func (rsp *ResponseUpload) GetFileURL() string {
	return fmt.Sprintf("%su/%s", BaseURL, rsp.ID)
}

type ResponseDownload struct {
	StatusCode int    `json:"code"`
	FilePath   string `json:"file_path"`
	FileName   string `json:"file_name"`
	FileSize   int64  `json:"file_size"`
	ResponseDefault
}

type ResponseFileInfo struct {
	ID                string    `json:"id"`
	Name              string    `json:"name"`
	Size              int       `json:"size"`
	Views             int       `json:"views"`
	BandwidthUsed     int       `json:"bandwidth_used"`
	BandwidthUsedPaid int       `json:"bandwidth_used_paid"`
	Downloads         int       `json:"downloads"`
	DateUpload        time.Time `json:"date_upload"`
	DateLastView      time.Time `json:"date_last_view"`
	MimeType          string    `json:"mime_type"`
	ThumbnailHref     string    `json:"thumbnail_href"`
	HashSha256        string    `json:"hash_sha256"`
	CanEdit           bool      `json:"can_edit"`
	StatusCode        int       `json:"code"`
	ResponseDefault
}
