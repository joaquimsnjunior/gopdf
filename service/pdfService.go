package service

import "os"
import "github.com/SebastiaanKlippert/go-wkhtmltopdf"
import "github.com/google/uuid"

type wk struct {
	rootPath string
}

func NewHtmlToPdf(rootPath string) *wk {
	return &wk{rootPath: rootPath}
}

func (w *wk) CreatePDF(htmlFile string) (string, error) {
	// Abre o arquivo HTML e converte para PDF
	f, err := os.Open(htmlFile)
	if err != nil {
		return "", err
	}
	pdfg, err := wkhtmltopdf.NewPDFGenerator()

	pdfg.AddPage(wkhtmltopdf.NewPageReader(f))

	//cria o PDF
	if err := pdfg.Create(); err != nil {
		return "", err
	}
	fileName := "file/" + uuid.New().String() + ".pdf"

	// Salva o PDF com nome aleatório
	if err := pdfg.WriteFile(fileName); err != nil {
		return "", err
	}
	return fileName, nil
}
