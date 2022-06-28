package pd

import (
	"io"
	"path/filepath"
)

// Auth hold the auth information
type Auth struct {
	APIKey string // if you have an account you can enter here your API Key for uploading in your account
}

// IsAuthAvailable checks if an API Key available
func (a *Auth) IsAuthAvailable() bool {
	auth := false
	if a.APIKey != "" {
		auth = true
	}

	return auth
}

// RequestUpload container for the upload information
type RequestUpload struct {
	File       io.ReadCloser
	PathToFile string // path to the file "/home/user/cat.jpg"
	FileName   string // just the filename "test.jpg"
	Anonymous  bool   // if the upload is anonymous or with auth
	Auth       Auth
	URL        string // specific the upload endpoint, is set by default with the correct values
}

// GetFileName return the filename from the path if no specific filename in the params
func (r *RequestUpload) GetFileName() string {
	if r.FileName == "" {
		if r.PathToFile != "" {
			r.FileName = filepath.Base(r.PathToFile)
		}
	}

	return r.FileName
}

// RequestDownload container for the file download
type RequestDownload struct {
	ID         string
	PathToSave string
	Auth       Auth
	URL        string // specific the API endpoint, is set by default with the correct values
}

// RequestFileInfo the FileInfo request needs only an ID
type RequestFileInfo struct {
	ID   string
	Auth Auth
	URL  string
}

// RequestThumbnail the Thumbnail request needs the ID and width and height
type RequestThumbnail struct {
	ID         string
	Width      string
	Height     string
	PathToSave string
	Auth       Auth
	URL        string
}

// RequestDelete delete the file if you are the owner with the given ID
type RequestDelete struct {
	ID   string
	Auth Auth
	URL  string
}

// RequestCreateList parameters for creating new list
type RequestCreateList struct {
	Title     string     `json:"title"`
	Anonymous bool       `json:"anonymous"`
	Files     []ListFile `json:"files"`
	Auth      Auth
	URL       string
}

// ListFile a file inside a CreateList request
type ListFile struct {
	ID          string `json:"id"`
	Description string `json:"description"`
}

// RequestGetList request to a retrieve a list
type RequestGetList struct {
	ID   string `json:"id"`
	Auth Auth
	URL  string
}

// RequestGetUserFiles ...
type RequestGetUserFiles struct {
	Auth Auth
	URL  string
}

// RequestGetUserLists ...
type RequestGetUserLists struct {
	Auth Auth
	URL  string
}
