package app

import (
	"errors"
	"fmt"
	"github.com/ManuelReschke/go-pd/pkg/pd"
	"github.com/spf13/cobra"
)

func RunDownload(cmd *cobra.Command, args []string) error {
	if len(args) == 0 {
		return errors.New("please add a pixeldrain URL or file id to your download request")
	}

	apiKey, err := cmd.Flags().GetString("api-key")
	if err != nil {
		return errors.New("please add a valid API-Key to your upload request")
	}

	// file is here an url or an ID to a file
	for _, file := range args {
		//@todo get file id from URL if a URL

		//@todo build req
		req := &pd.RequestDownload{
			ID:         file,
			PathToSave: "",
		}

		if apiKey != "" {
			req.Auth.APIKey = apiKey
		}

		c := pd.New(nil, nil)
		rsp, err := c.Download(req)
		if err != nil {
			return err
		}

		msg := ""
		if cmd.Flags().Changed("verbose") {
			msg = fmt.Sprintf("Successful! Download complete: %s | ID: %s | StoredTo: %s", rsp.FileName, req.ID, req.PathToSave)
		} else {
			msg = fmt.Sprintf("%s", req.PathToSave)
		}

		fmt.Println(msg)
	}

	return nil
}
