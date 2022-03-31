package pd_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/ManuelReschke/go-pd/pkg/pd"
)

func TestPD_RequestUpload(t *testing.T) {
	r := &pd.RequestUpload{
		PathToFile: "/test/path/file.data",
		Anonymous:  true,
		FileName:   "test123",
		URL:        "http://example.url",
		Auth:       pd.Auth{APIKey: "test-key"},
	}

	assert.Equal(t, "/test/path/file.data", r.PathToFile)
	assert.Equal(t, true, r.Anonymous)
	assert.Equal(t, "test123", r.FileName)
	assert.Equal(t, "http://example.url", r.URL)
	assert.Equal(t, "test-key", r.Auth.APIKey)
}

func TestPD_RequestUpload_GetFileName(t *testing.T) {
	ru := &pd.RequestUpload{
		PathToFile: "/test/path/file.data",
	}

	assert.Equal(t, "file.data", ru.GetFileName())
}

func TestPD_RequestDownload(t *testing.T) {
	r := &pd.RequestDownload{
		ID:   "123",
		URL:  "http://example.url",
		Auth: pd.Auth{APIKey: "test-key"},
	}

	assert.Equal(t, "123", r.ID)
	assert.Equal(t, "http://example.url", r.URL)
	assert.Equal(t, "test-key", r.Auth.APIKey)
}

func TestPD_RequestThumbnail(t *testing.T) {
	r := &pd.RequestThumbnail{
		ID:     "123",
		Width:  "16",
		Height: "16",
		URL:    "http://example.url",
		Auth:   pd.Auth{APIKey: "test-key"},
	}

	assert.Equal(t, "123", r.ID)
	assert.Equal(t, "16", r.Width)
	assert.Equal(t, "16", r.Height)
	assert.Equal(t, "http://example.url", r.URL)
	assert.Equal(t, "test-key", r.Auth.APIKey)
}
