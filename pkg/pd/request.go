package pd

import "path/filepath"

type RequestUpload struct {
	PathToFile string
	Anonymous  bool
	FileName   string
}

func (r *RequestUpload) GetFileName() string {
	if r.FileName == "" {
		r.FileName = filepath.Base(r.PathToFile)
	}
	return r.FileName
}
