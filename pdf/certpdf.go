package pdf

import (
	"fmt"
	"os"
	"path"

	"certificate.go/cert"
	"github.com/jung-kurt/gofpdf"
)

type PdfSaver struct {
	OutputDir string
}

func New(outputdir string) (*PdfSaver, error) {
	var p *PdfSaver
	err := os.MkdirAll(outputdir, os.ModePerm)
	if err != nil {
		return p, err
	}

	p = &PdfSaver{
		OutputDir: outputdir,
	}
	return p, nil
}

func (p *PdfSaver) Save(c cert.Cert) error {
	pdf := gofpdf.New(gofpdf.OrientationLandscape, "mm", "A4", "")
	pdf.SetTitle(c.LabelTitle, false)
	pdf.AddPage()

	// Background
	setBackground(pdf)

	// Header
	setHeader(pdf, &c)
	pdf.Ln(15.0)

	// Body
	pdf.SetFont("Helvetica", "I", 20)
	pdf.WriteAligned(0, 105, c.LabelPresented, "C")
	pdf.Ln(20.0)

	pdf.SetFont("Times", "B", 40)
	pdf.WriteAligned(0, 105, c.Name, "C")
	pdf.Ln(20.0)

	pdf.SetFont("Helvetica", "I", 20)
	pdf.WriteAligned(0, 105, c.LabelParticipation, "C")
	pdf.Ln(14.0)

	pdf.SetFont("Helvetica", "I", 15)
	pdf.WriteAligned(0, 105, c.LabelDate, "C")

	// Footer
	setFooter(pdf)

	// save pdf file
	filename := fmt.Sprintf("%v.pdf", c.LabelTitle)
	path := path.Join(p.OutputDir, filename)
	err := pdf.OutputFileAndClose(path)
	if err != nil {
		return err
	}
	fmt.Printf("Saved certificate to '%v'", path)
	return nil
}

func setBackground(pdf *gofpdf.Fpdf) {
	opts := gofpdf.ImageOptions{
		ImageType: "png",
	}
	pW, pH := pdf.GetPageSize()
	pdf.ImageOptions(
		"assets/img/background.png",
		0, 0,
		pW, pH,
		false, opts, 0, "",
	)
}

func setHeader(pdf *gofpdf.Fpdf, c *cert.Cert) {
	opts := gofpdf.ImageOptions{
		ImageType: "png",
	}

	margin := 36.0
	x := 0.0
	imageWidth := 30.0
	filename := "assets/img/gopher.png"
	pdf.ImageOptions(
		filename,
		x+margin, 65,
		imageWidth, 0,
		false, opts,
		0, "",
	)
	pW, _ := pdf.GetPageSize()
	x = pW - imageWidth
	pdf.ImageOptions(
		filename,
		x-margin, 65,
		imageWidth, 0,
		false, opts, 0, "",
	)

	pdf.SetFont("Arial", "", 37)
	pdf.WriteAligned(0, 85, c.LabelCompletion, "C")
}

func setFooter(pdf *gofpdf.Fpdf) {
	opts := gofpdf.ImageOptions{
		ImageType: "png",
	}

	margin := -180.0
	x := 0.0
	imageWidth := 50.0
	filename := "assets/img/stamp.png"
	pdf.ImageOptions(
		filename,
		x-margin, 127,
		imageWidth, 0,
		false, opts, 0, "",
	)
}
