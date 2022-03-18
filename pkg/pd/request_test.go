package pd_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/ManuelReschke/go-pd/pkg/pd"
)

func TestPD_RequestUpload(t *testing.T) {
	ru := &pd.RequestUpload{
		PathToFile: "/test/path/file.data",
		Anonymous:  true,
		FileName:   "test123",
		URL:        "http://example.url",
		Auth:       pd.Auth{APIKey: "test-key"},
	}

	assert.Equal(t, "/test/path/file.data", ru.PathToFile)
	assert.Equal(t, true, ru.Anonymous)
	assert.Equal(t, "test123", ru.FileName)
	assert.Equal(t, "http://example.url", ru.URL)
	assert.Equal(t, "test-key", ru.Auth.APIKey)
}

func TestPD_RequestUpload_GetFileName(t *testing.T) {
	ru := &pd.RequestUpload{
		PathToFile: "/test/path/file.data",
	}

	assert.Equal(t, "file.data", ru.GetFileName())
}
