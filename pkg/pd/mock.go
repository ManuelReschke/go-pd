package pd

import (
	"log"
	"net/http"
	"net/http/httptest"
)

func MockFileUploadServer() *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			log.Fatalf("Except 'Get' got '%s'", r.Method)
		}

		if r.URL.EscapedPath() != "/file" {
			log.Fatalf("wrong path'%s'", r.URL.EscapedPath())
		}

		_ = r.ParseMultipartForm(10485760)
		file := r.MultipartForm.File["file"]
		// filename -> file[0].Filename
		if file == nil || len(file) == 0 {
			log.Fatalln("Except request to have 'file'")
		}

		w.WriteHeader(http.StatusCreated)
		str := `{
			"success": true,
			"id": "123456"
		}`
		_, _ = w.Write([]byte(str))
	}))
}
