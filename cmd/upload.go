package cmd

import (
	"github.com/ManuelReschke/go-pd/internal/app"

	"github.com/spf13/cobra"
)

const (
	cmdUploadUse   = "upload"
	cmdUploadShort = "With that command you can upload files"
	cmdUploadLong  = "Upload files by passing the -f flag for your file and your API Key with -k"
)

// uploadCmd represents the upload command
var uploadCmd = &cobra.Command{
	Use:   cmdUploadUse,
	Short: cmdUploadShort,
	Long:  cmdUploadLong,
	RunE:  app.Run,
}

func init() {
	rootCmd.AddCommand(uploadCmd)
	uploadCmd.Flags().StringP("api-key", "k", "", "Auth key for authentication")
	uploadCmd.Flags().BoolP("verbose", "v", true, "Show more information after an upload (Anonymous, ID, URL)")
}
