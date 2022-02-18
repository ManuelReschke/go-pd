[![Version](https://img.shields.io/github/v/release/ManuelReschke/go-pd)](https://github.com/ManuelReschke/go-pd/releases)
![GitHub](https://img.shields.io/github/license/ManuelReschke/go-pd)
![GitHub code size in bytes](https://img.shields.io/github/languages/code-size/ManuelReschke/go-pd)
![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/ManuelReschke/go-pd)
![GitHub top language](https://img.shields.io/github/languages/top/ManuelReschke/go-pd)

# Go-PD - Another pixeldrain.com client

A free pixeldrain.com client written in go. We use the super power from [imroc/req](https://github.com/imroc/req) (v0.3.2) to build a robust and fast pixeldrain client.

## Why?

Because we want a simple, fast and robust go package to upload to pixeldrain.com.

## Import

```go
import "github.com/ManuelReschke/go-pd/pkg/pd"
```

## Example - use package to upload a file

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
		FileName:   "test_post_cat.jpg",
	}

	opt := &pd.ClientOptions{
		Debug:             false,
		Timeout:           5 * time.Second,
		EnableInsecureTLS: true,
	}

	c := pd.New(opt, nil)
	rsp, err := c.UploadPOST(req)
	if err != nil {
		t.Error(err)
	}

	fmt.Println(rsp.GetFileURL())

    // ID = xFNz76Vp
    // URL = https://pixeldrain.com/u/xFNz76Vp
}
```
## ToDo's:

- [x] implement simple upload method over POST /file
- [x] implement simple upload over PUT /file/{filename}
- [x] write integration test for the upload method
- [ ] implement all other API methods
- [ ] create CLI tool for uploading to pixeldrain.com

## PixelDrain methods covered by this package

### File Methods
| Status  |  PixelDrain Call |  Package Func |
|---|---|---|
|  [<span style="color:green">DONE</span>] | POST - /file  | UploadPOST(r *RequestUpload) (*ResponseUpload, error)  |
|  [<span style="color:green">DONE</span>] | PUT - /file/{name}  |  UploadPUT(r *RequestUpload) (*ResponseUpload, error) |
|  [<span style="color:red">ToDo</span>] | GET - /file/{id}  | -  |
|  [<span style="color:red">ToDo</span>] | GET - /file/{id}/info |  - |
|  [<span style="color:red">ToDo</span>] | GET - /file/{id}/thumbnail?width=x&height=x | -  |
|  [<span style="color:red">ToDo</span>] | DELETE - /file/{id}  | -  |
### List Methods
| Status  |  PixelDrain Call |  Package Func |
|---|---|---|
|  [<span style="color:red">ToDo</span>] | POST - /list | -  |
|  [<span style="color:red">ToDo</span>] | GET - /list/{id} | -  |
### User Methods
| Status  |  PixelDrain Call |  Package Func |
|---|---|---|
|  [<span style="color:red">ToDo</span>] | POST - /user/files | -  |
|  [<span style="color:red">ToDo</span>] | GET - /user/lists | -  |

## License

This software is released under the MIT License, see LICENSE.
