package cmd

import (
	"log"
	"os"
	"path"

	"github.com/spf13/cobra"
	chrome "github.com/spliffone/report-rendering-analysis/go-html-to-pdf/pkg/chrome"
	html2pdf "github.com/spliffone/report-rendering-analysis/go-html-to-pdf/pkg/html2pdf"
	"github.com/thediveo/enumflag"
)

type GeneratorMode enumflag.Flag

// Define the enumeration values for FooMode.
const (
	WK2HTML GeneratorMode = iota
	CHROME
)

// Map 3rd party enumeration values to their textual representations
var GeneratorIDs = map[GeneratorMode][]string{
	WK2HTML: {"WK2HTML"},
	CHROME:  {"CHROME"},
}

// Now use the GeneratorMode enum flag. If you want a non-zero default, then
// simply set it here, such as in "generatorMode = WK2HTML".
var generatorMode GeneratorMode = WK2HTML

// printPdfCmd represents the printPdf command
var printPdfCmd = &cobra.Command{
	Use:        "printPdf [flags] url",
	Short:      "Print PDF from HTML",
	Long:       `Print PDF from HTML.`,
	Args:       cobra.MinimumNArgs(1),
	ArgAliases: []string{"HTML file"},
	Run: func(cmd *cobra.Command, args []string) {
		log.Printf("Args: %v", args)

		var pdf []byte
		var err error
		switch generatorMode {
		case WK2HTML:
			pdf, err = html2pdf.PrintPDF(args[0], path.Dir(args[0]), "Landscape")
		case CHROME:
			pdf, err = chrome.PrintPdf(args[0])
		}

		if err != nil {
			log.Fatalf("Failed to print PDF: %v", err)
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
	rootCmd.AddCommand(printPdfCmd)

	// Here you will define your flags and configuration settings.
	printPdfCmd.Flags().StringP("cwd", "c", "", "Change working dir")
	printPdfCmd.Flags().StringP("out", "o", "output.pdf", "Output for PDF")
	printPdfCmd.Flags().VarP(
		enumflag.New(&generatorMode, "generator", GeneratorIDs, enumflag.EnumCaseInsensitive),
		"generator", "g",
		"generator can be 'WK2HTML' or 'CHROME'")
}
