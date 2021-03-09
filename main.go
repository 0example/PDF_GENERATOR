package main

// pdf generated with gofpdf
import (
	"github.com/jung-kurt/gofpdf"
)
// Generate the PDf by adding texts to page
func PDF(nameOfFile string) error {

	//set page size
	pdf := gofpdf.New("p","mm","A4","")
	//adding the page, following above setting
	pdf.AddPage()
	// font setting
	pdf.SetFont("Arial","B",18)
    //title for age
	pdf.CellFormat(180, 10, "Hello GO...!", "0", 0, "CM", false, 0, "")
	
	return pdf.OutputFileAndClose(nameOfFile)
}
func main() {

	//generate PDF following given name
	err := PDF("new.pdf")
	if err != nil {
		panic(err)
	}
}
