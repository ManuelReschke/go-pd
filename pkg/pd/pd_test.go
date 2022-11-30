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

var fileIDPost string
var fileIDPut string
var listID string

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
	fileIDPost = rsp.ID
	fmt.Println("POST Req: " + rsp.GetFileURL())
}

// TestPD_UploadPOST_WithReadCloser_Integration run a real integration test against the service
func TestPD_UploadPOST_WithReadCloser_Integration(t *testing.T) {
	if testing.Short() {
		t.Skip(SkipIntegrationTest)
	}

	// ReadCloser
	file, _ := os.Open("testdata/cat.jpg")

	req := &pd.RequestUpload{
		File:     file,
		FileName: "test_post_cat.jpg",
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
	fileIDPut = rsp.ID
	fmt.Println("PUT Req: " + rsp.GetFileURL())
}

// TestPD_UploadPUT_WithReadCloser_Integration run a real integration test against the service
func TestPD_UploadPUT_WithReadCloser_Integration(t *testing.T) {
	if testing.Short() {
		t.Skip(SkipIntegrationTest)
	}

	// ReadCloser
	file, _ := os.Open("testdata/cat.jpg")

	req := &pd.RequestUpload{
		File:     file,
		FileName: "test_put_cat.jpg",
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
		ID:         fileIDPost,
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
		ID: fileIDPost,
	}

	req.Auth = setAuthFromEnv()

	c := pd.New(nil, nil)
	rsp, err := c.GetFileInfo(req)
	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, 200, rsp.StatusCode)
	assert.Equal(t, true, rsp.Success)
	assert.Equal(t, fileIDPost, rsp.ID)
	assert.Equal(t, 37621, rsp.Size)
	assert.Equal(t, "1af93d68009bdfd52e1da100a019a30b5fe083d2d1130919225ad0fd3d1fed0b", rsp.HashSha256)
}

// TestPD_DownloadThumbnail is a unit test for the GET "download thumbnail" method
func TestPD_DownloadThumbnail(t *testing.T) {
	server := pd.MockFileUploadServer()
	defer server.Close()
	testURL := server.URL + "/file/K1dA8U5W/thumbnail?width=64&height=64"

	req := &pd.RequestThumbnail{
		ID:         "K1dA8U5W",
		Height:     "64",
		Width:      "64",
		PathToSave: "testdata/cat_download_thumbnail.jpg",
		URL:        testURL,
	}

	req.Auth = setAuthFromEnv()

	c := pd.New(nil, nil)
	rsp, err := c.DownloadThumbnail(req)
	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, 200, rsp.StatusCode)
	assert.Equal(t, "cat_download_thumbnail.jpg", rsp.FileName)
	assert.Equal(t, int64(7056), rsp.FileSize)
}

// TestPD_DownloadThumbnail_Integration run a real integration test against the service
func TestPD_DownloadThumbnail_Integration(t *testing.T) {
	if testing.Short() {
		t.Skip(SkipIntegrationTest)
	}

	req := &pd.RequestThumbnail{
		ID:         fileIDPost,
		Height:     "64",
		Width:      "64",
		PathToSave: "testdata/cat_download_thumbnail.jpg",
	}

	req.Auth = setAuthFromEnv()

	c := pd.New(nil, nil)
	rsp, err := c.DownloadThumbnail(req)
	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, 200, rsp.StatusCode)
	assert.Equal(t, "cat_download_thumbnail.jpg", rsp.FileName)
	assert.Equal(t, int64(7056), rsp.FileSize)
}

// TestPD_CreateList is a unit test for the POST "list" method
func TestPD_CreateList(t *testing.T) {
	server := pd.MockFileUploadServer()
	defer server.Close()
	testURL := server.URL + "/list"

	// files to add
	files := []pd.ListFile{
		{ID: "K1dA8U5W", Description: "Hallo Welt"},
		{ID: "bmrc4iyD", Description: "Hallo Welt 2"},
	}

	// create list request
	req := &pd.RequestCreateList{
		Title:     "Test List",
		Anonymous: false,
		Files:     files,
		URL:       testURL,
	}

	req.Auth = setAuthFromEnv()

	c := pd.New(nil, nil)
	rsp, err := c.CreateList(req)
	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, 200, rsp.StatusCode)
	assert.Equal(t, true, rsp.Success)
	assert.NotEmpty(t, rsp.ID)
}

// TestPD_Delete_Integration run a real integration test against the service
func TestPD_CreateList_Integration(t *testing.T) {
	if testing.Short() {
		t.Skip(SkipIntegrationTest)
	}

	// files to add
	files := []pd.ListFile{
		{ID: fileIDPost, Description: "Hallo Welt"},
		{ID: fileIDPut, Description: "Hallo Welt 2"},
	}

	// create list request
	req := &pd.RequestCreateList{
		Title:     "Test List",
		Anonymous: false,
		Files:     files,
	}

	req.Auth = setAuthFromEnv()

	c := pd.New(nil, nil)
	rsp, err := c.CreateList(req)
	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, 201, rsp.StatusCode)
	assert.Equal(t, true, rsp.Success)
	listID = rsp.ID
}

// TestPD_GetList is a unit test for the GET "list/{id}" method
func TestPD_GetList(t *testing.T) {
	server := pd.MockFileUploadServer()
	defer server.Close()
	testURL := server.URL + "/list/123"

	req := &pd.RequestGetList{
		ID:  "123",
		URL: testURL,
	}

	req.Auth = setAuthFromEnv()

	c := pd.New(nil, nil)
	rsp, err := c.GetList(req)
	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, 200, rsp.StatusCode)
	assert.Equal(t, true, rsp.Success)
	assert.NotEmpty(t, rsp.ID)
	assert.Equal(t, "Rust in Peace", rsp.Title)
	assert.Equal(t, 123456, rsp.Files[0].Size)
}

// TestPD_GetList_Integration run a real integration test against the service
func TestPD_GetList_Integration(t *testing.T) {
	if testing.Short() {
		t.Skip(SkipIntegrationTest)
	}

	req := &pd.RequestGetList{
		ID: listID,
	}

	req.Auth = setAuthFromEnv()

	c := pd.New(nil, nil)
	rsp, err := c.GetList(req)
	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, 200, rsp.StatusCode)
	assert.Equal(t, true, rsp.Success)
	assert.NotEmpty(t, rsp.ID)
	assert.Equal(t, "Test List", rsp.Title)
	assert.Equal(t, 37621, rsp.Files[0].Size)
}

// TestPD_GetUser is a unit test for the GET "/user" method
func TestPD_GetUser(t *testing.T) {
	server := pd.MockFileUploadServer()
	defer server.Close()
	testURL := server.URL + "/user"

	req := &pd.RequestGetUser{
		URL: testURL,
	}

	req.Auth = setAuthFromEnv()

	c := pd.New(nil, nil)
	rsp, err := c.GetUser(req)
	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, 200, rsp.StatusCode)
	assert.Equal(t, true, rsp.Success)
	assert.Equal(t, "TestTest", rsp.Username)
	assert.Equal(t, "Free", rsp.Subscription.Name)
}

// TestPD_GetUser_Integration run a real integration test against the service
func TestPD_GetUser_Integration(t *testing.T) {
	if testing.Short() {
		t.Skip(SkipIntegrationTest)
	}

	req := &pd.RequestGetUser{}

	req.Auth = setAuthFromEnv()

	c := pd.New(nil, nil)
	rsp, err := c.GetUser(req)
	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, 200, rsp.StatusCode)
	assert.Equal(t, true, rsp.Success)
	assert.Equal(t, "ManuelReschke", rsp.Username)
	assert.Equal(t, "Free", rsp.Subscription.Name)
}

// TestPD_GetUserFiles is a unit test for the GET "/user/files" method
func TestPD_GetUserFiles(t *testing.T) {
	server := pd.MockFileUploadServer()
	defer server.Close()
	testURL := server.URL + "/user/files"

	req := &pd.RequestGetUserFiles{
		URL: testURL,
	}

	req.Auth = setAuthFromEnv()

	c := pd.New(nil, nil)
	rsp, err := c.GetUserFiles(req)
	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, 200, rsp.StatusCode)
	assert.Equal(t, true, rsp.Success)
	assert.Equal(t, "tUxgDCoQ", rsp.Files[0].ID)
	assert.Equal(t, "test_post_cat.jpg", rsp.Files[0].Name)
}

// TestPD_GetUserFiles_Integration run a real integration test against the service
func TestPD_GetUserFiles_Integration(t *testing.T) {
	if testing.Short() {
		t.Skip(SkipIntegrationTest)
	}

	req := &pd.RequestGetUserFiles{}

	req.Auth = setAuthFromEnv()

	c := pd.New(nil, nil)
	rsp, err := c.GetUserFiles(req)
	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, 200, rsp.StatusCode)
	assert.Equal(t, true, rsp.Success)
	if rsp.Files[0].ID == fileIDPost {
		assert.Equal(t, fileIDPost, rsp.Files[0].ID)
	} else {
		assert.Equal(t, fileIDPut, rsp.Files[0].ID)
	}
}

// TestPD_GetUserLists is a unit test for the GET "/user/files" method
func TestPD_GetUserLists(t *testing.T) {
	server := pd.MockFileUploadServer()
	defer server.Close()
	testURL := server.URL + "/user/lists"

	req := &pd.RequestGetUserLists{
		URL: testURL,
	}

	req.Auth = setAuthFromEnv()

	c := pd.New(nil, nil)
	rsp, err := c.GetUserLists(req)
	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, 200, rsp.StatusCode)
	assert.Equal(t, true, rsp.Success)
	assert.Equal(t, "Test List", rsp.Lists[0].Title)
}

// TestPD_GetUserLists_Integration run a real integration test against the service
func TestPD_GetUserLists_Integration(t *testing.T) {
	if testing.Short() {
		t.Skip(SkipIntegrationTest)
	}

	req := &pd.RequestGetUserLists{}

	req.Auth = setAuthFromEnv()

	c := pd.New(nil, nil)
	rsp, err := c.GetUserLists(req)
	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, 200, rsp.StatusCode)
	assert.Equal(t, true, rsp.Success)
	assert.Equal(t, "Test List", rsp.Lists[0].Title)
}

// TestPD_Delete is a unit test for the DELETE "delete" method
func TestPD_Delete(t *testing.T) {
	server := pd.MockFileUploadServer()
	defer server.Close()
	testURL := server.URL + "/file/K1dA8U5W"

	req := &pd.RequestDelete{
		ID:  "K1dA8U5W",
		URL: testURL,
	}

	req.Auth = setAuthFromEnv()

	c := pd.New(nil, nil)
	rsp, err := c.Delete(req)
	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, true, rsp.Success)
	assert.Equal(t, "file_deleted", rsp.Value)
	assert.Equal(t, "The file has been deleted.", rsp.Message)
}

// TestPD_Delete_Integration run a real integration test against the service
func TestPD_Delete_Integration(t *testing.T) {
	if testing.Short() {
		t.Skip(SkipIntegrationTest)
	}

	req := &pd.RequestDelete{
		ID: fileIDPost,
	}

	req.Auth = setAuthFromEnv()

	c := pd.New(nil, nil)
	rsp, err := c.Delete(req)
	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, true, rsp.Success)
	assert.Equal(t, "ok", rsp.Value)
	assert.Equal(t, "The requested action was successfully performed", rsp.Message)

	req = &pd.RequestDelete{
		ID: fileIDPut,
	}

	req.Auth = setAuthFromEnv()

	c = pd.New(nil, nil)
	rsp, err = c.Delete(req)
	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, true, rsp.Success)
	assert.Equal(t, "ok", rsp.Value)
	assert.Equal(t, "The requested action was successfully performed", rsp.Message)
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
