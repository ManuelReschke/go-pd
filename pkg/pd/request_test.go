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
	}

	assert.Equal(t, "/test/path/file.data", ru.PathToFile)
	assert.Equal(t, true, ru.Anonymous)
	assert.Equal(t, "test123", ru.FileName)
}

func TestPD_RequestUpload_GetFileName(t *testing.T) {
	ru := &pd.RequestUpload{
		PathToFile: "/test/path/file.data",
	}

	assert.Equal(t, "file.data", ru.GetFileName())
}
