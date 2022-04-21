package pd_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/ManuelReschke/go-pd/pkg/pd"
)

func TestPD_ResponseDefault(t *testing.T) {
	rsp := &pd.ResponseDefault{}
	rsp.Success = true
	rsp.Value = "test"
	rsp.Message = "test message"

	assert.Equal(t, true, rsp.Success)
	assert.Equal(t, "test", rsp.Value)
	assert.Equal(t, "test message", rsp.Message)
}

func TestPD_ResponseUpload(t *testing.T) {
	rsp := &pd.ResponseUpload{}
	rsp.StatusCode = 201
	rsp.ID = "test123"
	rsp.Success = true
	rsp.Value = "test"
	rsp.Message = "test message"

	assert.Equal(t, 201, rsp.StatusCode)
	assert.Equal(t, "test123", rsp.ID)
	assert.Equal(t, true, rsp.Success)
	assert.Equal(t, "test", rsp.Value)
	assert.Equal(t, "test message", rsp.Message)
}

func TestPD_ResponseDownload(t *testing.T) {
	rsp := &pd.ResponseDownload{}
	rsp.StatusCode = 200
	rsp.Success = true
	rsp.Value = "test"
	rsp.Message = "test message"
	rsp.FileName = "filename"
	rsp.FileSize = 123123
	rsp.FilePath = "/my/path/file.jpg"

	assert.Equal(t, 200, rsp.StatusCode)
	assert.Equal(t, true, rsp.Success)
	assert.Equal(t, "test", rsp.Value)
	assert.Equal(t, "test message", rsp.Message)
	assert.Equal(t, "filename", rsp.FileName)
	assert.Equal(t, int64(123123), rsp.FileSize)
	assert.Equal(t, "/my/path/file.jpg", rsp.FilePath)
}

func TestPD_ResponseThumbnail(t *testing.T) {
	rsp := &pd.ResponseThumbnail{}
	rsp.StatusCode = 200
	rsp.Success = true
	rsp.Value = "test"
	rsp.Message = "test message"
	rsp.FileName = "filename"
	rsp.FileSize = 123123
	rsp.FilePath = "/my/path/thumbnail.jpg"

	assert.Equal(t, 200, rsp.StatusCode)
	assert.Equal(t, true, rsp.Success)
	assert.Equal(t, "test", rsp.Value)
	assert.Equal(t, "test message", rsp.Message)
	assert.Equal(t, "filename", rsp.FileName)
	assert.Equal(t, int64(123123), rsp.FileSize)
	assert.Equal(t, "/my/path/thumbnail.jpg", rsp.FilePath)
}

func TestPD_ResponseDelete(t *testing.T) {
	rsp := &pd.ResponseDelete{}
	rsp.StatusCode = 200
	rsp.Success = true
	rsp.Value = "test"
	rsp.Message = "test message"

	assert.Equal(t, 200, rsp.StatusCode)
	assert.Equal(t, true, rsp.Success)
	assert.Equal(t, "test", rsp.Value)
	assert.Equal(t, "test message", rsp.Message)
}

func TestPD_ResponseCreateList(t *testing.T) {
	rsp := &pd.ResponseCreateList{}
	rsp.StatusCode = 200
	rsp.Success = true
	rsp.Value = "test"
	rsp.Message = "test message"
	rsp.ID = "123"

	assert.Equal(t, 200, rsp.StatusCode)
	assert.Equal(t, true, rsp.Success)
	assert.Equal(t, "test", rsp.Value)
	assert.Equal(t, "test message", rsp.Message)
	assert.Equal(t, "123", rsp.ID)
}

func TestPD_ResponseGetList(t *testing.T) {
	rsp := &pd.ResponseGetList{}
	rsp.StatusCode = 200
	rsp.Success = true
	rsp.Value = "test"
	rsp.Message = "test message"
	rsp.ID = "123"
	rsp.Title = "Test Title"
	layout := "2014-09-12T11:45:26.371Z"
	timeStr := "2020-02-04T18:34:13.466276Z"
	rsp.DateCreated, _ = time.Parse(layout, timeStr)
	//@todo
	rsp.Files = []pd.FileGetList{{
		DetailHref:    "",
		Description:   "",
		Success:       false,
		ID:            "",
		Name:          "",
		Size:          0,
		DateCreated:   time.Time{},
		DateLastView:  time.Time{},
		MimeType:      "",
		Views:         0,
		BandwidthUsed: 0,
		ThumbnailHref: "",
	}}

	assert.Equal(t, 200, rsp.StatusCode)
	assert.Equal(t, true, rsp.Success)
	assert.Equal(t, "test", rsp.Value)
	assert.Equal(t, "test message", rsp.Message)
	assert.Equal(t, "123", rsp.ID)
}

func TestPD_ResponseGetUserFiles(t *testing.T) {
	rsp := &pd.ResponseGetUserFiles{}
	rsp.StatusCode = 200
	rsp.Success = true
	rsp.Value = "test"
	rsp.Message = "test message"

	assert.Equal(t, 200, rsp.StatusCode)
	assert.Equal(t, true, rsp.Success)
	assert.Equal(t, "test", rsp.Value)
	assert.Equal(t, "test message", rsp.Message)
}

func TestPD_ResponseGetUserLists(t *testing.T) {
	rsp := &pd.ResponseGetUserLists{}
	rsp.StatusCode = 200
	rsp.Success = true
	rsp.Value = "test"
	rsp.Message = "test message"

	assert.Equal(t, 200, rsp.StatusCode)
	assert.Equal(t, true, rsp.Success)
	assert.Equal(t, "test", rsp.Value)
	assert.Equal(t, "test message", rsp.Message)
}
