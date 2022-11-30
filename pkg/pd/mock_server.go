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
			if r.URL.EscapedPath() == "/file" {
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
			}

			// ##########################################
			// POST /list
			if r.URL.EscapedPath() == "/list" {
				_ = r.ParseForm()

				w.WriteHeader(http.StatusOK)
				str := `{
					"success": true,
					"id": "123456"
				}`
				_, _ = w.Write([]byte(str))
			}

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

			// ##########################################
			// GET /file/{id}/thumbnail?width=x&height=x
			if r.URL.EscapedPath() == "/file/K1dA8U5W/thumbnail" {
				_ = r.ParseForm()

				fileContent, err := ioutil.ReadFile("testdata/cat_thumbnail.jpg")
				if err != nil {
					log.Fatalln(err)
				}

				w.WriteHeader(http.StatusOK)
				w.Write(fileContent)
			}

			// ##########################################
			// GET /list/{id}
			if r.URL.EscapedPath() == "/list/123" {
				_ = r.ParseForm()

				w.WriteHeader(http.StatusOK)
				str := `{
				  "success": true,
				  "id": "123",
				  "title": "Rust in Peace",
				  "date_created": "2020-02-04T18:34:13.466276Z",
				  "files": [
					{
					  "detail_href": "/file/_SqVWi/info",
					  "description": "",
					  "success": true,
					  "id": "_SqVWi",
					  "name": "01 Holy Wars... The Punishment Due.mp3",
					  "size": 123456,
					  "date_created": "2020-02-04T18:34:13.466276Z",
					  "date_last_view": "2020-02-04T18:34:13.466276Z",
					  "mime_type": "audio/mp3",
					  "views": 1,
					  "bandwidth_used": 1234567890,
					  "thumbnail_href": "/file/_SqVWi/thumbnail"
					},
					{
					  "detail_href": "/file/RKwgZb/info",
					  "description": "",
					  "success": true,
					  "id": "RKwgZb",
					  "name": "02 Hangar 18.mp3",
					  "size": 123456,
					  "date_created": "2020-02-04T18:34:13.466276Z",
					  "date_last_view": "2020-02-04T18:34:13.466276Z",
					  "mime_type": "audio/mp3",
					  "views": 2,
					  "bandwidth_used": 1234567890,
					  "thumbnail_href": "/file/RKwgZb/thumbnail"
					}
				  ]
				}`
				_, _ = w.Write([]byte(str))
			}

			// ##########################################
			// GET /user/files
			if r.URL.EscapedPath() == "/user/files" {
				_ = r.ParseForm()

				w.WriteHeader(http.StatusOK)
				str := `{
				  "files": [
					{
					  "id": "tUxgDCoQ",
					  "name": "test_post_cat.jpg",
					  "size": 37621,
					  "views": 0,
					  "bandwidth_used": 0,
					  "bandwidth_used_paid": 0,
					  "downloads": 0,
					  "date_upload": "2022-03-30T16:30:17.152Z",
					  "date_last_view": "2022-03-30T16:30:17.152Z",
					  "mime_type": "image/jpeg",
					  "thumbnail_href": "/file/tUxgDCoQ/thumbnail",
					  "hash_sha256": "1af93d68009bdfd52e1da100a019a30b5fe083d2d1130919225ad0fd3d1fed0b",
					  "availability": "",
					  "availability_message": "",
					  "abuse_type": "",
					  "abuse_reporter_name": "",
					  "can_edit": true,
					  "show_ads": false,
					  "allow_video_player": true,
					  "download_speed_limit": 0
					}
				  ]
				}`
				_, _ = w.Write([]byte(str))
			}

			// ##########################################
			// GET /user
			if r.URL.EscapedPath() == "/user" {
				_ = r.ParseForm()

				w.WriteHeader(http.StatusOK)
				str := `{
				   "username":"TestTest",
				   "email":"test@test.de",
				   "subscription":{
					  "id":"",
					  "name":"Free",
					  "type":"",
					  "file_size_limit":20000000000,
					  "file_expiry_days":60,
					  "storage_space":-1,
					  "price_per_tb_storage":0,
					  "price_per_tb_bandwidth":0,
					  "monthly_transfer_cap":0,
					  "file_viewer_branding":false
				   },
				   "storage_space_used":18834,
				   "is_admin":false,
				   "balance_micro_eur":0,
				   "hotlinking_enabled":true,
				   "monthly_transfer_cap":0,
				   "monthly_transfer_used":0,
				   "file_viewer_branding":null,
				   "file_embed_domains":"",
				   "skip_file_viewer":false
				}`
				_, _ = w.Write([]byte(str))
			}

			// ##########################################
			// GET /user/lists
			if r.URL.EscapedPath() == "/user/lists" {
				_ = r.ParseForm()

				w.WriteHeader(http.StatusOK)
				str := `{
				  "lists": [
					{
					  "id": "Cap4T1LP",
					  "title": "Test List",
					  "date_created": "2022-04-04T15:24:06.834Z",
					  "file_count": 2,
					  "files": null,
					  "can_edit": true
					},
					{
					  "id": "fiEm5arj",
					  "title": "Wallpaper",
					  "date_created": "2022-04-20T09:46:42.017Z",
					  "file_count": 2,
					  "files": null,
					  "can_edit": true
					}
				  ]
				}`
				_, _ = w.Write([]byte(str))
			}

			return
		case "DELETE":
			// ##########################################
			// DELETE /file/{id}
			if r.URL.EscapedPath() == "/file/K1dA8U5W" {

				w.WriteHeader(http.StatusOK)
				str := `{
					"success": true,
					"value": "file_deleted",
					"message": "The file has been deleted."
				}`
				_, _ = w.Write([]byte(str))
			}

			return
		default:
			return
		}
	}))
}
