package cmd

import (
	"github.com/ManuelReschke/go-pd/internal/app"

	"github.com/spf13/cobra"
)

const (
	cmdDownloadUse   = "download"
	cmdDownloadShort = "With that command you can download a file"
	cmdDownloadLong  = "Download file by passing the file url or file id and your API Key with -k"
)

// downloadCmd represents the upload command
var downloadCmd = &cobra.Command{
	Use:   cmdDownloadUse,
	Short: cmdDownloadShort,
	Long:  cmdDownloadLong,
	RunE:  app.RunDownload,
}

func init() {
	rootCmd.AddCommand(downloadCmd)
	downloadCmd.Flags().StringP("api-key", "k", "", "Auth key for authentication")
	downloadCmd.Flags().BoolP("verbose", "v", true, "Show more information after an upload (Anonymous, ID, URL)")
}
