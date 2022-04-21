[![Version](https://img.shields.io/github/v/release/ManuelReschke/go-pd)](https://github.com/ManuelReschke/go-pd/releases)
![GitHub](https://img.shields.io/github/license/ManuelReschke/go-pd)
![GitHub code size in bytes](https://img.shields.io/github/languages/code-size/ManuelReschke/go-pd)
![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/ManuelReschke/go-pd)
![GitHub top language](https://img.shields.io/github/languages/top/ManuelReschke/go-pd)

# Go-PD - Another pixeldrain.com client

A free pixeldrain.com client written in go. We use the super power from [imroc/req](https://github.com/imroc/req) (v0.3.2) to build a robust and fast pixeldrain client.

## Why?

Because we want a simple, fast, robust and 100% tested go package to upload to pixeldrain.com.

### Import the pkg / lib (Go < v1.7)

```
 go get github.com/ManuelReschke/go-pd/pkg/pd
```

### Import the pkg / lib (Go > v1.7)

```
 go install github.com/ManuelReschke/go-pd/pkg/pd
```

## Example 1 - the easiest way to upload an anonymous file

```go
package main

import (
	"fmt"
	"time"

	"github.com/ManuelReschke/go-pd/pkg/pd"
)

func main() {
	req := &pd.RequestUpload{
		PathToFile: "testdata/cat.jpg",
		Anonymous:  true,
	}

	c := pd.New(nil, nil)
	rsp, err := c.UploadPOST(req)
	if err != nil {
		fmt.Println(err)
	}

	// print the full URL
	fmt.Println(rsp.GetFileURL())

        // example ID = xFNz76Vp
        // example URL = https://pixeldrain.com/u/xFNz76Vp
}
```

## Example 2 - advanced way to upload a file to user account

```go
package main

import (
	"fmt"
	"time"

	"github.com/ManuelReschke/go-pd/pkg/pd"
)

func main() {
	req := &pd.RequestUpload{
		PathToFile: "testdata/cat.jpg",
		FileName:   "test_post_cat.jpg",
		Anonymous:  false,
		Auth: pd.Auth{
			APIKey: "you-api-key-from-pixeldrain-account",
		},
	}

	// set specific request options
	opt := &pd.ClientOptions{
		Debug:             false,
		ProxyURL:          "example.socks5.proxy",
		EnableCookies:     true,
		EnableInsecureTLS: true,
		Timeout:           1 * time.Hour,
	}

	c := pd.New(opt, nil)
	rsp, err := c.UploadPOST(req)
	if err != nil {
		fmt.Println(err)
	}

	// print the full URL
	fmt.Println(rsp.GetFileURL())

        // example ID = xFNz76Vp
        // example URL = https://pixeldrain.com/u/xFNz76Vp
}
```
## ToDo's:

- [x] implement simple upload method over POST /file
- [x] implement simple upload over PUT /file/{filename}
- [x] write unit tests
- [x] write integration tests
- [x] add API-KEY auth to requests
- [ ] implement all other API methods
  - [x] implement GET - /file/{id}
  - [x] implement GET - /file/{id}/info
  - [x] implement GET - /file/{id}/thumbnail?width=x&height=x
  - [x] implement DELETE - /file/{id}
  - [x] implement POST - /list
  - [X] implement GET - /list/{id}
  - [x] implement GET - /user/files
  - [x] implement GET - /user/lists
- [ ] create CLI tool for uploading to pixeldrain.com
- [ ] refactor the hole shit and use nice to have patterns (like Option Pattern)

## PixelDrain methods covered by this package

### File Methods
| PixelDrain Call                                 |  Package Func |
|-------------------------------------------------|---|
| [x] POST - /file                                | UploadPOST(r *RequestUpload) (*ResponseUpload, error) |
| [x] PUT - /file/{name}                          | UploadPUT(r *RequestUpload) (*ResponseUpload, error) |
| [x] GET - /file/{id}                            | Download(r *RequestDownload) (*ResponseDownload, error) |
| [x] GET - /file/{id}/info                       | GetFileInfo(r *RequestFileInfo) (*ResponseFileInfo, error) |
| [x] GET - /file/{id}/thumbnail?width=x&height=x | DownloadThumbnail(r *RequestThumbnail) (*ResponseThumbnail, error)  |
| [x] DELETE - /file/{id}                         | Delete(r *RequestDelete) (*ResponseDelete, error)  |
### List Methods
| PixelDrain Call      |  Package Func |
|----------------------|---|
| [x] POST - /list     | CreateList(r *RequestCreateList) (*ResponseCreateList, error)  |
| [x] GET - /list/{id} | GetList(r *RequestGetList) (*ResponseGetList, error)  |
### User Methods
| PixelDrain Call        |  Package Func |
|------------------------|---|
| [x] POST - /user/files | GetUserFiles(r *RequestGetUserFiles) (*ResponseGetUserFiles, error)  |
| [x] GET - /user/lists  | GetUserLists(r *RequestGetUserLists) (*ResponseGetUserLists, error)  |

## Package CLI commands

### Unit Tests - Run pkg unit tests
Run unit tests against a local emulated server.
```shell
make test
```

### Integration Tests - Run pkg integration tests
Run real integration tests against the real pixeldrain.com website.
```shell
make test-integration
```

### Test Coverage - create test coverage report
Create a coverage report c.out and a coverage.html to view the results in web browser
```shell
make coverage
```

## License

This software is released under the MIT License, see LICENSE.
