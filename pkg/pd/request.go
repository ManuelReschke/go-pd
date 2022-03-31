package pd

import "path/filepath"

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
	PathToFile string // path to the file "/home/user/cat.jpg"
	FileName   string // just the filename "test.jpg"
	Anonymous  bool   // if the upload is anonymous or with auth
	Auth       Auth
	URL        string // specific the upload endpoint, is set by default with the correct values
}

// GetFileName return the filename from the path if no specific filename in the params
func (r *RequestUpload) GetFileName() string {
	if r.FileName == "" {
		r.FileName = filepath.Base(r.PathToFile)
	}
	return r.FileName
}

// RequestDownload container for the file download
type RequestDownload struct {
	ID         string
	Download   bool
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
