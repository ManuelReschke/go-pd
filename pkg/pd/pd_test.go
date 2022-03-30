package pd_test

import (
	"fmt"
	"github.com/ManuelReschke/go-pd/pkg/pd"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"os"
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
	}

	req.Auth = setAuthFromEnv()

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
	}

	req.Auth = setAuthFromEnv()

	c := pd.New(nil, nil)
	rsp, err := c.UploadPUT(req)
	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, 201, rsp.StatusCode)
	assert.NotEmpty(t, rsp.ID)
	fmt.Println("PUT Req: " + rsp.GetFileURL())
}

// TestPD_Download is a unit test for the GET "download" method
func TestPD_Download(t *testing.T) {
	server := pd.MockFileUploadServer()
	defer server.Close()
	testURL := server.URL + "/file/K1dA8U5W"

	req := &pd.RequestDownload{
		PathToSave: "testdata/cat_download.jpg",
		ID:         "K1dA8U5W",
		URL:        testURL,
	}

	c := pd.New(nil, nil)
	rsp, err := c.Download(req)
	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, 200, rsp.StatusCode)
	assert.Equal(t, true, rsp.Success)
}

// TestPD_Download_Integration run a real integration test against the service
func TestPD_Download_Integration(t *testing.T) {
	if testing.Short() {
		t.Skip(SkipIntegrationTest)
	}

	req := &pd.RequestDownload{
		PathToSave: "testdata/cat_download.jpg",
		ID:         "K1dA8U5W",
	}

	req.Auth = setAuthFromEnv()

	c := pd.New(nil, nil)
	rsp, err := c.Download(req)
	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, 200, rsp.StatusCode)
	assert.Equal(t, "cat_download.jpg", rsp.FileName)
	assert.Equal(t, int64(37621), rsp.FileSize)
}

// TestPD_GetFileInfo is a unit test for the GET "file info" method
func TestPD_GetFileInfo(t *testing.T) {
	server := pd.MockFileUploadServer()
	defer server.Close()
	testURL := server.URL + "/file/K1dA8U5W/info"

	req := &pd.RequestFileInfo{
		ID:  "K1dA8U5W",
		URL: testURL,
	}

	c := pd.New(nil, nil)
	rsp, err := c.GetFileInfo(req)
	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, 200, rsp.StatusCode)
	assert.Equal(t, true, rsp.Success)
	assert.Equal(t, "K1dA8U5W", rsp.ID)
	assert.Equal(t, 37621, rsp.Size)
	assert.Equal(t, "1af93d68009bdfd52e1da100a019a30b5fe083d2d1130919225ad0fd3d1fed0b", rsp.HashSha256)
}

// TestPD_GetFileInfo_Integration run a real integration test against the service
func TestPD_GetFileInfo_Integration(t *testing.T) {
	if testing.Short() {
		t.Skip(SkipIntegrationTest)
	}

	req := &pd.RequestFileInfo{
		ID: "K1dA8U5W",
	}

	req.Auth = setAuthFromEnv()

	c := pd.New(nil, nil)
	rsp, err := c.GetFileInfo(req)
	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, 200, rsp.StatusCode)
	assert.Equal(t, true, rsp.Success)
	assert.Equal(t, "K1dA8U5W", rsp.ID)
	assert.Equal(t, 37621, rsp.Size)
	assert.Equal(t, "1af93d68009bdfd52e1da100a019a30b5fe083d2d1130919225ad0fd3d1fed0b", rsp.HashSha256)
}

func setAuthFromEnv() pd.Auth {
	// load api key from .env_test file
	currentWorkDirectory, _ := os.Getwd()
	_ = godotenv.Load(currentWorkDirectory + "/.env_test")
	apiKey := os.Getenv("API_KEY")

	return pd.Auth{
		APIKey: apiKey,
	}
}
