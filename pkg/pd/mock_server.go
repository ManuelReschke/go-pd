package pd

import (
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
)

func MockFileUploadServer() *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		switch r.Method {
		case "POST":
			if r.URL.EscapedPath() != "/file" {
				log.Fatalf("wrong path'%s'", r.URL.EscapedPath())
			}

			_ = r.ParseMultipartForm(10485760)
			file := r.MultipartForm.File["file"]

			if file == nil || len(file) == 0 {
				log.Fatalln("Except request to have 'file'")
			}

			w.WriteHeader(http.StatusCreated)
			str := `{
				"success": true,
				"id": "123456"
			}`
			_, _ = w.Write([]byte(str))

		case "PUT":
			if !strings.Contains(r.URL.EscapedPath(), "/file/") {
				log.Fatalf("wrong path'%s'", r.URL.EscapedPath())
			}

			_ = r.ParseForm()

			if r.Body == nil || r.ContentLength == 0 {
				log.Fatalln("Empty body in PUT request")
			}

			if r.FormValue("anonymous") == "" {
				log.Fatalln("Except request to have form value 'anonymous'")
			}

			if r.FormValue("name") == "" {
				log.Fatalln("Except request to have form value 'name'")
			}

			w.WriteHeader(http.StatusCreated)
			str := `{
				"id": "123456"
			}`
			_, _ = w.Write([]byte(str))
		case "GET":
			return
		default:
			return
		}
	}))
}
