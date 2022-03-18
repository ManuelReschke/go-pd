package pd

import "path/filepath"

// RequestUpload container for the upload information
type RequestUpload struct {
	PathToFile string // path to the file "/home/user/cat.jpg"
	FileName   string // just the filename "test.jpg"
	Anonymous  bool   // if the upload is anonymous or with auth
	Auth       Auth
	URL        string // specific the upload endpoint, is set by default with the correct values
}

// Auth hold the auth information
type Auth struct {
	APIKey string // if you have an account you can enter here your API Key for uploading in your account
}

// GetFileName return the filename from the path if no specific filename in the params
func (r *RequestUpload) GetFileName() string {
	if r.FileName == "" {
		r.FileName = filepath.Base(r.PathToFile)
	}
	return r.FileName
}
