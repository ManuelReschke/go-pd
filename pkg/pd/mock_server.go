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
			// ##########################################
			// POST /file
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
			// ##########################################
			// PUT /file/{name}
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
			// ##########################################
			// GET /file/{id}
			if r.URL.EscapedPath() == "/file/K1dA8U5W" {
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
			}

			// ##########################################
			// GET /file/{id}/info
			if r.URL.EscapedPath() == "/file/K1dA8U5W/info" {
				_ = r.ParseForm()

				fileID := filepath.Base(r.URL.EscapedPath())
				if len(fileID) == 0 {
					log.Fatalf("empty file ID '%s'", fileID)
				}

				w.WriteHeader(http.StatusOK)
				str := `{
				  "id": "K1dA8U5W",
				  "name": "screenshot.png",
				  "size": 37621,
				  "views": 1234,
				  "bandwidth_used": 1234567890,
				  "bandwidth_used_paid": 1234567890,
				  "downloads": 1234,
				  "date_upload": "2020-02-04T18:34:05.706801Z",
				  "date_last_view": "2020-02-04T18:34:05.706801Z",
				  "mime_type": "image/png",
				  "thumbnail_href": "/file/1234abcd/thumbnail",
				  "hash_sha256": "1af93d68009bdfd52e1da100a019a30b5fe083d2d1130919225ad0fd3d1fed0b",
				  "can_edit": true
				}`
				_, _ = w.Write([]byte(str))
			}

			return
		default:
			return
		}
	}))
}
