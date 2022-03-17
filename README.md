[![Version](https://img.shields.io/github/v/release/ManuelReschke/go-pd)](https://github.com/ManuelReschke/go-pd/releases)
![GitHub](https://img.shields.io/github/license/ManuelReschke/go-pd)
![GitHub code size in bytes](https://img.shields.io/github/languages/code-size/ManuelReschke/go-pd)
![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/ManuelReschke/go-pd)
![GitHub top language](https://img.shields.io/github/languages/top/ManuelReschke/go-pd)

# Go-PD - Another pixeldrain.com client

A free pixeldrain.com client written in go. We use the super power from [imroc/req](https://github.com/imroc/req) (v0.3.2) to build a robust and fast pixeldrain client.

## Why?

Because we want a simple, fast, robust and 100% tested go package to upload to pixeldrain.com.

## Import the pkg / lib

```go
import "github.com/ManuelReschke/go-pd/pkg/pd"
```

## Example 1 - the easiest way to upload a file

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

## Example 2 - advanced way to upload a file

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
		Anonymous:  true,
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
- [ ] add API-KEY auth to requests
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
