package pd_test

import (
	"fmt"
	"github.com/ManuelReschke/go-pd/pkg/pd"
	"github.com/stretchr/testify/assert"
	"testing"
)

const SkipIntegrationTest = "skipping integration test"

// TestPD_UploadPOST is a unit test for the POST upload method
func TestPD_UploadPOST(t *testing.T) {
	server := pd.MockFileUploadServer()
	defer server.Close()
	testURL := server.URL + "/file"

	req := &pd.RequestUpload{
		PathToFile: "testdata/cat.jpg",
		FileName:   "test_post_cat.jpg",
		Anonymous:  true,
		URL:        testURL,
	}

	c := pd.New(nil, nil)
	rsp, err := c.UploadPOST(req)
	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, 201, rsp.StatusCode)
	assert.NotEmpty(t, rsp.ID)
	assert.Equal(t, "https://pixeldrain.com/u/123456", rsp.GetFileURL())
	fmt.Println("POST Req: " + rsp.GetFileURL())
}

// TestPD_UploadPOST_Integration run a real integration test against the service
func TestPD_UploadPOST_Integration(t *testing.T) {
	if testing.Short() {
		t.Skip(SkipIntegrationTest)
	}

	req := &pd.RequestUpload{
		PathToFile: "testdata/cat.jpg",
		FileName:   "test_post_cat.jpg",
		Anonymous:  true,
	}

	c := pd.New(nil, nil)
	rsp, err := c.UploadPOST(req)
	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, 201, rsp.StatusCode)
	assert.NotEmpty(t, rsp.ID)
	fmt.Println("POST Req: " + rsp.GetFileURL())
}

// TestPD_UploadPUT is a unit test for the PUT upload method
func TestPD_UploadPUT(t *testing.T) {
	server := pd.MockFileUploadServer()
	defer server.Close()
	testURL := server.URL + "/file/"

	req := &pd.RequestUpload{
		PathToFile: "testdata/cat.jpg",
		FileName:   "test_put_cat.jpg",
		Anonymous:  true,
		URL:        testURL + "test_put_cat.jpg",
	}

	c := pd.New(nil, nil)
	rsp, err := c.UploadPUT(req)
	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, 201, rsp.StatusCode)
	assert.NotEmpty(t, rsp.ID)
	assert.Equal(t, "https://pixeldrain.com/u/123456", rsp.GetFileURL())
	fmt.Println("PUT Req: " + rsp.GetFileURL())
}

// TestPD_UploadPUT_Integration run a real integration test against the service
func TestPD_UploadPUT_Integration(t *testing.T) {
	if testing.Short() {
		t.Skip(SkipIntegrationTest)
	}

	req := &pd.RequestUpload{
		PathToFile: "testdata/cat.jpg",
		FileName:   "test_put_cat.jpg",
		Anonymous:  true,
	}

	c := pd.New(nil, nil)
	rsp, err := c.UploadPUT(req)
	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, 201, rsp.StatusCode)
	assert.NotEmpty(t, rsp.ID)
	fmt.Println("PUT Req: " + rsp.GetFileURL())
}
