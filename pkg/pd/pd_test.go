package pd_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/ManuelReschke/go-pd/pkg/pd"
)

func Test_Upload(t *testing.T) {
	req := &pd.RequestUpload{
		PathToFile: "testdata/cat.jpg",
		Anonymous:  true,
	}

	opt := &pd.ClientOptions{
		Debug:             false,
		Timeout:           5 * time.Second,
		EnableInsecureTLS: true,
	}

	c := pd.New(opt, nil)
	rsp, err := c.Upload(req)
	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, 201, rsp.StatusCode)
	assert.NotEmpty(t, rsp.ID)
	fmt.Println(rsp.ID)
}
