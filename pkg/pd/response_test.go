package pd_test

import (
	"testing"

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
