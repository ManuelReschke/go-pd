package app

import (
	"errors"
	"fmt"
	"github.com/ManuelReschke/go-pd/pkg/pd"
	"github.com/spf13/cobra"
	"os"
	"path/filepath"
	"strings"
)

func RunDownload(cmd *cobra.Command, args []string) error {
	if len(args) == 0 {
		return errors.New("please add a pixeldrain URL or file id to your download request")
	}

	path, err := cmd.Flags().GetString("path")
	if err != nil {
		return errors.New("please add a valid path where you want to save the files")
	}
	if path == "" {
		path, _ = os.Getwd()
	}

	apiKey, err := cmd.Flags().GetString("api-key")
	if err != nil {
		return errors.New("please add a valid API-Key to your request")
	}

	// file is here an url or an ID to a file
	for _, file := range args {
		fileID := file
		if strings.ContainsAny(file, pd.BaseURL) {
			fileID = filepath.Base(file)
		}

		req01 := &pd.RequestFileInfo{
			ID: fileID,
		}
		if apiKey != "" {
			req01.Auth.APIKey = apiKey
		}

		c := pd.New(nil, nil)
		rsp, err := c.GetFileInfo(req01)
		if err != nil {
			return err
		}

		req := &pd.RequestDownload{
			ID:         fileID,
			PathToSave: filepath.FromSlash(path + "/" + rsp.Name),
		}
		if apiKey != "" {
			req.Auth.APIKey = apiKey
		}

		rspDL, err := c.Download(req)
		if err != nil {
			return err
		}

		msg := ""
		if cmd.Flags().Changed("verbose") {
			msg = fmt.Sprintf("Successful! Download complete: %s | ID: %s | Stored to: %s", rspDL.FileName, req.ID, req.PathToSave)
		} else {
			msg = fmt.Sprintf("%s", req.PathToSave)
		}

		fmt.Println(msg)
	}

	return nil
}
