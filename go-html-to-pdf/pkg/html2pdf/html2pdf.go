package html2pdf

import (
	wkhtml "github.com/SebastiaanKlippert/go-wkhtmltopdf"
)

func PrintPDF(htmlStr string, localPath string, orientation string) ([]byte, error) {
	// Create new PDF generator
	pdfg, err := wkhtml.NewPDFGenerator()
	if err != nil {
		return nil, err
	}
	pdfg.LogLevel.Set("info")
	pdfg.Orientation.Set(orientation)
	pdfg.PageSize.Set("A4")
	page := wkhtml.NewPage(htmlStr)
	page.PageOptions.Allow.Set(localPath)
	page.PageOptions.EnableLocalFileAccess.Set(true)
	page.PrintMediaType.Set(true)
	pdfg.AddPage(page)

	// Create PDF document in internal buffer
	err = pdfg.Create()
	if err != nil {
		return nil, err
	}

	return pdfg.Bytes(), nil
}
