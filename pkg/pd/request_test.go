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

func TestPD_RequestDelete(t *testing.T) {
	r := &pd.RequestDelete{
		ID:   "123",
		URL:  "http://example.url",
		Auth: pd.Auth{APIKey: "test-key"},
	}

	assert.Equal(t, "123", r.ID)
	assert.Equal(t, "http://example.url", r.URL)
	assert.Equal(t, "test-key", r.Auth.APIKey)
}

func TestPD_RequestCreateList(t *testing.T) {
	r := &pd.RequestCreateList{
		Title:     "test",
		Anonymous: true,
		Files: []pd.ListFile{
			{ID: "123", Description: "Test Description"},
			{ID: "456", Description: "Test Description"},
		},
		URL:  "http://example.url",
		Auth: pd.Auth{APIKey: "test-key"},
	}

	assert.Equal(t, "test", r.Title)
	assert.Equal(t, true, r.Anonymous)
	assert.Equal(t, 2, len(r.Files))
	assert.Equal(t, "123", r.Files[0].ID)
	assert.Equal(t, "Test Description", r.Files[0].Description)
	assert.Equal(t, "http://example.url", r.URL)
	assert.Equal(t, "test-key", r.Auth.APIKey)
}

func TestPD_RequestGetList(t *testing.T) {
	r := &pd.RequestGetList{
		ID:   "123",
		URL:  "http://example.url",
		Auth: pd.Auth{APIKey: "test-key"},
	}

	assert.Equal(t, "123", r.ID)
	assert.Equal(t, "http://example.url", r.URL)
	assert.Equal(t, "test-key", r.Auth.APIKey)
}

func TestPD_RequestGetUserFiles(t *testing.T) {
	r := &pd.RequestGetUserFiles{
		URL:  "http://example.url",
		Auth: pd.Auth{APIKey: "test-key"},
	}

	assert.Equal(t, "http://example.url", r.URL)
	assert.Equal(t, "test-key", r.Auth.APIKey)
}

func TestPD_RequestGetUserLists(t *testing.T) {
	r := &pd.RequestGetUserLists{
		URL:  "http://example.url",
		Auth: pd.Auth{APIKey: "test-key"},
	}

	assert.Equal(t, "http://example.url", r.URL)
	assert.Equal(t, "test-key", r.Auth.APIKey)
}
