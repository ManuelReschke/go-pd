[![Version](https://img.shields.io/github/v/release/ManuelReschke/go-pd)](https://github.com/ManuelReschke/go-pd/releases)

# go-pd - another pixeldrain.com client

A free pixeldrain.com client written in go. We use the super power from [imroc/req](https://github.com/imroc/req) (v0.3.2) to build a robust and fast pixeldrain client.

## Why?

Because we want a simple, fast and robust upload to pixeldrain.com.

## ToDo:

- [x] implement simple upload method (POST /file)
- [x] write integration test for the upload method
- [ ] implement all other API methods
- [ ] write tests for all other stuff
- [ ] create CLI tool for uploading to pixeldrain.com

## Import

```go
import "github.com/ManuelReschke/go-pd/pkg/pd"
```

## Example

```go
package main

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/ManuelReschke/go-pd/pkg/pd"
)

func main() {
	req := &pd.RequestUpload{
		PathToFile: "testdata/cat.jpg",
		Anonymous:  true,
	}

	opt := &pd.ClientOptions{
		Debug:             false,
		Timeout:           5 * time.Second,
		EnableInsecureTLS: true,
	}

	c := pd.New(opt, nil)
	rsp, err := c.Upload(req)
	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, 201, rsp.StatusCode)
	assert.NotEmpty(t, rsp.ID)
	fmt.Println(rsp.ID)

    // ID = xFNz76Vp
    // URL = https://pixeldrain.com/u/xFNz76Vp
}
```

## License

This software is released under the MIT License, see LICENSE.
