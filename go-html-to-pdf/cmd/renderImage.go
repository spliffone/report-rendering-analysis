package cmd

import (
	"log"
	"os"

	"github.com/spf13/cobra"
	chrome "github.com/spliffone/report-rendering-analysis/go-html-to-pdf/pkg/chrome"
)

// renderImageCmd represents the renderImage command
var renderImageCmd = &cobra.Command{
	Use:   "renderImage",
	Short: "Render image",
	Long:  `Render image`,
	Run: func(cmd *cobra.Command, args []string) {
		log.Printf("Args: %v", args)

		pdf, err := chrome.RenderImage(args[0])
		if err != nil {
			log.Fatalf("Failed to render image: %v", err)
		}

		output, _ := cmd.Flags().GetString("out")
		err = os.WriteFile(output, pdf, 0644)
		if err != nil {
			log.Fatalf("Failed to write PDF: %v", err)
		}
	},
	PreRun: changeDir,
}

func init() {
	rootCmd.AddCommand(renderImageCmd)

	// Here you will define your flags and configuration settings.
	renderImageCmd.Flags().StringP("cwd", "c", "", "Change working dir")
	renderImageCmd.Flags().StringP("out", "o", "output.png", "Output for image")
}
