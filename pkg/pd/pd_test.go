package pd_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/ManuelReschke/go-pd/pkg/pd"
)

func TestPD_UploadPOST(t *testing.T) {
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

	assert.Equal(t, 201, rsp.StatusCode)
	assert.NotEmpty(t, rsp.ID)
	fmt.Println("POST Req: " + rsp.GetFileURL())
}

func TestPD_UploadPUT(t *testing.T) {
	req := &pd.RequestUpload{
		PathToFile: "testdata/cat.jpg",
		Anonymous:  true,
		FileName:   "test_put_cat.jpg",
	}

	opt := &pd.ClientOptions{
		Debug:             false,
		Timeout:           5 * time.Second,
		EnableInsecureTLS: true,
	}

	c := pd.New(opt, nil)
	rsp, err := c.UploadPUT(req)
	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, 201, rsp.StatusCode)
	assert.NotEmpty(t, rsp.ID)
	fmt.Println("PUT Req: " + rsp.GetFileURL())
}
