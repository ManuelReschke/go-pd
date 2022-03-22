package pd

import (
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"path/filepath"
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

			if r.FormValue("anonymous") == "" {
				log.Fatalln("Except request to have form value 'anonymous'")
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

			//if r.FormValue("anonymous") == "" {
			//	log.Fatalln("Except request to have form value 'anonymous'")
			//}

			w.WriteHeader(http.StatusCreated)
			str := `{
				"id": "123456"
			}`
			_, _ = w.Write([]byte(str))
		case "GET":
			if !strings.Contains(r.URL.EscapedPath(), "/file/") {
				log.Fatalf("wrong path'%s'", r.URL.EscapedPath())
			}

			_ = r.ParseForm()

			fileID := filepath.Base(r.URL.EscapedPath())
			if len(fileID) == 0 {
				log.Fatalf("empty file ID '%s'", fileID)
			}

			fileContent, err := ioutil.ReadFile("testdata/cat.jpg")
			if err != nil {
				log.Fatalln(err)
			}

			w.WriteHeader(http.StatusOK)
			w.Write(fileContent)

			return
		default:
			return
		}
	}))
}
