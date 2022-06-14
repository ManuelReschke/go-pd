package app

import (
	"errors"
	"fmt"
	"github.com/ManuelReschke/go-pd/pkg/pd"
	"github.com/spf13/cobra"
	"os"
	"path/filepath"
)

func RunUpload(cmd *cobra.Command, args []string) error {
	if len(args) == 0 {
		return errors.New("please add a file to your upload request")
	}

	apiKey, err := cmd.Flags().GetString("api-key")
	if err != nil {
		return errors.New("please add a valid API-Key to your upload request")
	}

	for _, file := range args {
		// check if file exist
		if _, err := os.Stat(filepath.FromSlash(file)); errors.Is(err, os.ErrNotExist) {
			return errors.New("one of the given files does not exist")
		}

		req := &pd.RequestUpload{
			PathToFile: file,
			Anonymous:  true,
		}

		if apiKey != "" {
			req.Anonymous = false
			req.Auth.APIKey = apiKey
		}

		c := pd.New(nil, nil)
		rsp, err := c.UploadPOST(req)
		if err != nil {
			return err
		}

		msg := ""
		if cmd.Flags().Changed("verbose") {
			msg = fmt.Sprintf("Successful! Anonymous upload: %v | ID: %s | URL: %s", req.Anonymous, rsp.ID, rsp.GetFileURL())
		} else {
			msg = fmt.Sprintf("%s", rsp.GetFileURL())
		}

		fmt.Println(msg)
	}

	return nil
}
