package pd_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/ManuelReschke/go-pd/pkg/pd"
)

func TestPD_ResponseUpload(t *testing.T) {
	rsp := &pd.ResponseUpload{
		StatusCode: 201,
		ID:         "test123",
		Success:    true,
		Value:      "test",
		Message:    "test message",
	}

	assert.Equal(t, 201, rsp.StatusCode)
	assert.Equal(t, "test123", rsp.ID)
	assert.Equal(t, true, rsp.Success)
	assert.Equal(t, "test", rsp.Value)
	assert.Equal(t, "test message", rsp.Message)
}