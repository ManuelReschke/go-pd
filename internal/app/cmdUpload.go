package app

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"github.com/ManuelReschke/go-pd/pkg/pd"
	"github.com/imroc/req/v3"
	"github.com/spf13/cobra"
)

func RunUpload(cmd *cobra.Command, args []string) error {
	if len(args) == 0 {
		return errors.New("please add a file to your upload request")
	}

	verboseFlag := cmd.Flags().Changed("verbose")

	// get key from ENV if available
	apiKeyEnv, ok := os.LookupEnv("PIXELDRAIN_API_KEY")
	if ok && verboseFlag {
		fmt.Println("Using API Key from environment variable: PIXELDRAIN_API_KEY")
	}

	apiKeyParam, err := cmd.Flags().GetString("api-key")
	if err != nil {
		return errors.New("please add a valid API-Key to your upload request")
	}

	for _, file := range args {
		// check if file exist
		if _, err := os.Stat(filepath.FromSlash(file)); errors.Is(err, os.ErrNotExist) {
			return errors.New("one of the given files does not exist")
		}

		r := &pd.RequestUpload{
			PathToFile: file,
			Anonymous:  true,
		}

		if apiKeyEnv != "" {
			r.Anonymous = false
			r.Auth.APIKey = apiKeyEnv
		}
		if apiKeyParam != "" {
			r.Anonymous = false
			r.Auth.APIKey = apiKeyParam
		}

		c := pd.New(nil, nil)
		if verboseFlag {
			c.SetUploadCallback(func(info req.UploadInfo) {
				if info.FileSize > 0 {
					fmt.Printf("%q uploaded %.2f%%\n", info.FileName, float64(info.UploadedSize)/float64(info.FileSize)*100.0)
				} else {
					fmt.Printf("%q uploaded 0%% (file size is zero)\n", info.FileName)
				}
			})
		}
		rsp, err := c.UploadPOST(r)
		if err != nil {
			return err
		}

		msg := ""
		if verboseFlag {
			msg = fmt.Sprintf("Successful! Anonymous upload: %v | ID: %s | URL: %s", r.Anonymous, rsp.ID, rsp.GetFileURL())
		} else {
			msg = fmt.Sprintf("%s", rsp.GetFileURL())
		}

		fmt.Println(msg)
	}

	return nil
}
