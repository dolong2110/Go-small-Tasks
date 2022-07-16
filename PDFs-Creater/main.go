package main

import (
	"fmt"
	"github.com/jung-kurt/gofpdf"
	"log"
)

func main() {
	pdf := gofpdf.New(gofpdf.OrientationLandscape, gofpdf.UnitPoint, gofpdf.PageSizeLetter, "")
	w, h := pdf.GetPageSize()
	fmt.Printf("width=%v, height=%v\n", w, h)
	pdf.AddPage()

	// Basic Text Stuff
	pdf.MoveTo(0, 0)
	pdf.SetFont("Arial", "B", 30)
	_, lineHt := pdf.GetFontSize()
	pdf.SetTextColor(255, 0, 0)
	pdf.Text(0, lineHt, "Hello, world!")
	pdf.MoveTo(0, lineHt * 2.0)

	pdf.SetFont("times", "", 18)
	pdf.SetTextColor(100, 100, 100)
	_, lineHt = pdf.GetFontSize()
	pdf.MultiCell(0, lineHt * 1.5, "Here is some text. If it is too long, it will be word wrapped automatically. If there is a new line, it will be\nwrapped as well (unlike other ways of writting text in gofpdf).", gofpdf.BorderNone, gofpdf.AlignRight, false)

	// Basic Shapes
	pdf.SetFillColor(0, 255, 0)
	pdf.SetDrawColor(0, 0, 255)
	pdf.Rect(10, 100, 100, 100, "FD")
	pdf.SetFillColor(100, 200, 200)
	pdf.Polygon([]gofpdf.PointType{
		{110, 250},
		{160, 300},
		{110, 350},
		{60, 300},
	}, "F")

	//pdf.ImageOptions("", )

	// Grid
	drawGrid(pdf)

	err := pdf.OutputFileAndClose("p1.pdf")
	if err != nil {
		log.Fatalln(err)
	}
}

func drawGrid(pdf *gofpdf.Fpdf) {
	pdf.SetFont("courier", "", 12)
	pdf.SetTextColor(80, 80, 80)
	pdf.SetDrawColor(200, 200, 200)

	w, h := pdf.GetPageSize()
	for x := 0.0; x < w; x += w / 20.0 {
		pdf.Line(x, 0, x, h)
		_, lineHt := pdf.GetFontSize()
		pdf.Text(x, lineHt, fmt.Sprintf("%d", int(x)))
	}

	for y := 0.0; y < h; y += h / 20.0 {
		pdf.Line(0, y, w, y)
		pdf.Text(0, y, fmt.Sprintf("%d", int(y)))
	}
}