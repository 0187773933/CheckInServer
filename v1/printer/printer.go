package printer

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"

	"github.com/jung-kurt/gofpdf"
)

// findMaxFontSize iteratively finds the largest font size (starting at maxCandidate)
// so that the rendered width of 'text' does not exceed maxWidth.
func findMaxFontSize(pdf *gofpdf.Fpdf, text, fontName string, maxWidth, maxCandidate float64) float64 {
	fontSize := maxCandidate
	pdf.SetFont(fontName, "", fontSize)
	textWidth := pdf.GetStringWidth(text)
	step := 0.5 // decrease by 0.5 points each iteration
	for textWidth > maxWidth && fontSize > step {
		fontSize -= step
		pdf.SetFont(fontName, "", fontSize)
		textWidth = pdf.GetStringWidth(text)
	}
	return fontSize
}

// centerCoords returns the top-left coordinates (x, y) needed to center an element
// of given dimensions (elementWidth x elementHeight) on a page of size (pageWidth x pageHeight).
func centerCoords(pageWidth, pageHeight, elementWidth, elementHeight float64) (x, y float64) {
	x = (pageWidth - elementWidth) / 2
	y = (pageHeight - elementHeight) / 2
	return
}

// addRotatedCenteredText draws text rotated by 'angle' degrees about the pivot point (pivotX, pivotY)
// and centers its bounding box at that pivot.
// (This version is used by PrintTwo and preserves your original behavior.)
func addRotatedCenteredText(pdf *gofpdf.Fpdf, text, fontName string, fontSize, pivotX, pivotY, angle float64) {
	pdf.SetFont(fontName, "", fontSize)
	textWidth := pdf.GetStringWidth(text)
	// Approximate text height in inches (1pt ≈ 1/72 in)
	textHeight := fontSize / 72.0
	// Original computation: simply center the unrotated bounding box
	x := pivotX - (textWidth / 2)
	y := pivotY - (textHeight / 2)

	pdf.TransformBegin()
	pdf.TransformRotate(angle, pivotX, pivotY)
	pdf.Text(x, y, text)
	pdf.TransformEnd()
}

// addRotatedCenteredTextPrint1 is used in Print.
// For a 90° rotation it adjusts the y-coordinate so that after rotation
// the text's bounding box is centered at (pivotX, pivotY).
func addRotatedCenteredTextPrint1(pdf *gofpdf.Fpdf, text, fontName string, fontSize, pivotX, pivotY, angle float64) {
	pdf.SetFont(fontName, "", fontSize)
	textWidth := pdf.GetStringWidth(text)
	textHeight := fontSize / 72.0

	var x, y float64
	// For a 90° rotation, the transformation rotates the text so that
	// the unrotated (x, y) point maps to:
	//   x' = pivotX - (y - pivotY)
	//   y' = pivotY + (x - pivotX)
	// We want the rotated bounding box (with width=textHeight and height=textWidth)
	// to be centered at (pivotX, pivotY). Solving gives:
	//   x = pivotX - (textWidth/2)
	//   y = pivotY + (textHeight/2)
	if angle == 90 {
		x = pivotX - (textWidth / 2)
		y = pivotY + (textHeight / 2)
	} else {
		// For other angles, default to the original computation.
		x = pivotX - (textWidth / 2)
		y = pivotY - (textHeight / 2)
	}

	pdf.TransformBegin()
	pdf.TransformRotate(angle, pivotX, pivotY)
	pdf.Text(x, y, text)
	pdf.TransformEnd()
}

// addRotatedImage adds an image rotated by 'angle' degrees about its center.
// The x coordinate is adjusted by -0.1 as per your original code.
func addRotatedImage(pdf *gofpdf.Fpdf, imagePath string, x, y, width, height, angle float64) {
	centerX := x + width/2
	centerY := y + height/2

	pdf.TransformBegin()
	pdf.TransformRotate(angle, centerX, centerY)
	pdf.ImageOptions(
		imagePath,
		x-0.1, // restored adjustment
		y,
		width, height,
		false,
		gofpdf.ImageOptions{
			ImageType:             "PNG",
			ReadDpi:               true,
			AllowNegativePosition: false,
		},
		0, "",
	)
	pdf.TransformEnd()
}

// -----------------
// OS-specific Printing Functions
// -----------------

func PrintWindows(printerName, pdfFilePath string) {
	sumatra_file_path  , _ := filepath.Abs( "./v1/embed/windows/SumatraPDF.exe" )
	args := []string{sumatra_file_path, "-print-to", printerName, pdfFilePath}
	cmd := exec.Command(args[0], args[1:]...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		fmt.Println("Error printing on Windows:", err)
	}
}

func ClearPrintQueMacOSX(printerName string) {
	cmd := exec.Command("cancel", "-a", printerName)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		fmt.Println("Error clearing printer queue on macOS:", err)
	}
}

func PrintMacOSX(printerName, pdfFilePath string, printSpeed int) {
	if printSpeed < 1 {
		printSpeed = 2
	}
	args := []string{"lp", "-d", printerName, "-o", fmt.Sprintf("PrintSpeed=%d", printSpeed), "-o", "fit-to-page", pdfFilePath}
	cmd := exec.Command(args[0], args[1:]...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		fmt.Println("Error printing on macOS:", err)
	}
}

func ClearPrintQueLinux(printerName string) {
	fmt.Println("ClearPrintQueLinux: TODO")
}

func PrintLinux(printerName, pdfFilePath string) {
	fmt.Println("PrintLinux: TODO")
}

// -----------------
// Main Print Functions
// -----------------

// Print creates a PDF ticket with a rotated text and a rotated logo.
// It uses findMaxFontSize to scale the text so that it is as large as possible
// without overflowing the available width. It uses the adjusted centering function
// (addRotatedCenteredTextPrint1) so that for a 90° rotation the text remains on the page.
func Print(printerName, name string) {
	// Page dimensions in inches.
	const pageWidth = 2.25
	const pageHeight = 4.0

	pdf := gofpdf.NewCustom(&gofpdf.InitType{
		OrientationStr: "P",
		UnitStr:        "in",
		Size:           gofpdf.SizeType{Wd: pageWidth, Ht: pageHeight},
	})
	pdf.SetMargins(0.1, 0.1, 0.1)
	pdf.AddPage()

	pdf.AddUTF8Font("ComicNeue", "", "./v1/embed/fonts/ComicNeue-Regular.ttf")

	// For Print, use a pivot that centers the rotated text.
	// Here we simply use the page center.
	pageCenterX := pageWidth / 2
	pageCenterY := pageHeight / 2
	// pivotX := pageCenterX
	pivotX := (pageCenterX - (pageCenterX / 2)) - 0.1
	pivotY := pageCenterY

	// For rotated text, let the available width be nearly the page height.
	const maxTextWidth = pageHeight - 0.2
	const candidateFontSize = 40.0

	maxFontSize := findMaxFontSize(pdf, name, "ComicNeue", maxTextWidth, candidateFontSize)

	// Draw the rotated text using the adjusted centering function.
	addRotatedCenteredTextPrint1(pdf, name, "ComicNeue", maxFontSize, pivotX, pivotY, 90)

	// --- Add rotated, centered logo ---
	logoScale := 0.6
	logoWidth := (300.0 / 96.0) * logoScale
	logoHeight := (156.0 / 96.0) * logoScale
	effectiveLogoWidth := logoHeight
	effectiveLogoHeight := logoWidth
	logoX, logoY := centerCoords(pageWidth, pageHeight, effectiveLogoWidth, effectiveLogoHeight)
	logoY += 0.3

	addRotatedImage(pdf, "./v1/embed/images/logo.png", logoX, logoY, logoWidth, logoHeight, 90)

	tmpFile, err := ioutil.TempFile("", "ticket-*.pdf")
	defer tmpFile.Close()
	if err != nil {
		fmt.Println("Error creating temporary file:", err)
		return
	}
	tmpFilePath := tmpFile.Name()
	defer func() {
		os.Remove( tmpFilePath )
	}()
	if err := pdf.OutputFileAndClose(tmpFilePath); err != nil {
		fmt.Println("Error outputting PDF:", err)
		return
	}
	fmt.Println("Printing", tmpFilePath)

	switch runtime.GOOS {
		case "windows":
			PrintWindows(printerName, tmpFilePath)
		case "darwin":
			ClearPrintQueMacOSX(printerName)
			PrintMacOSX(printerName, tmpFilePath, 2)
		case "linux":
			ClearPrintQueLinux(printerName)
			PrintLinux(printerName, tmpFilePath)
		default:
			fmt.Println("Unsupported OS")
	}

	fmt.Println("Printed ticket for", name)
}

// PrintTwo creates a PDF ticket with two pieces of rotated text (name and extraString)
// and a rotated logo. The main text's font size is maximized, and the extra text's font
// size is clamped (ensuring it is smaller than the main text) based on a fixed ratio.
func PrintTwo(printerName, name, extraString string) {
	const pageWidth = 2.25
	const pageHeight = 4.0

	pdf := gofpdf.NewCustom(&gofpdf.InitType{
		OrientationStr: "P",
		UnitStr:        "in",
		Size:           gofpdf.SizeType{Wd: pageWidth, Ht: pageHeight},
	})
	pdf.SetMargins(0.1, 0.1, 0.1)
	pdf.AddPage()

	pdf.AddUTF8Font("ComicNeue", "", "./v1/embed/fonts/ComicNeue-Regular.ttf")

	pageCenterX := pageWidth / 2
	pageCenterY := pageHeight / 2
	// Use your original pivot calculations for PrintTwo.
	pivotX := (pageCenterX - (pageCenterX / 2)) + (pageCenterX / 4)
	pivotXExtra := pivotX + 0.25
	pivotY := pageCenterY

	const maxTextWidth = pageHeight - 0.2
	const candidateFontSize = 40.0

	mainFontSize := findMaxFontSize(pdf, name, "ComicNeue", maxTextWidth, candidateFontSize)
	extraFontSizeCandidate := findMaxFontSize(pdf, extraString, "ComicNeue", maxTextWidth, candidateFontSize)
	desiredRatio := 0.4
	var extraFontSize float64
	if extraFontSizeCandidate > mainFontSize*desiredRatio {
		extraFontSize = mainFontSize * desiredRatio
	} else {
		extraFontSize = extraFontSizeCandidate
	}

	addRotatedCenteredText(pdf, name, "ComicNeue", mainFontSize, pivotX, pivotY, 90)
	addRotatedCenteredText(pdf, extraString, "ComicNeue", extraFontSize, pivotXExtra, pivotY, 90)

	logoScale := 0.6
	logoWidth := (300.0 / 96.0) * logoScale
	logoHeight := (156.0 / 96.0) * logoScale
	effectiveLogoWidth := logoHeight
	effectiveLogoHeight := logoWidth
	logoX, logoY := centerCoords(pageWidth, pageHeight, effectiveLogoWidth, effectiveLogoHeight)
	logoY += 0.3

	addRotatedImage(pdf, "./v1/embed/images/logo.png", logoX, logoY, logoWidth, logoHeight, 90)

	tmpFile, err := ioutil.TempFile("", "ticket-*.pdf")
	defer tmpFile.Close()
	if err != nil {
		fmt.Println("Error creating temporary file:", err)
		return
	}
	tmpFilePath := tmpFile.Name()
	defer func() {
		os.Remove( tmpFilePath )
	}()
	if err := pdf.OutputFileAndClose(tmpFilePath); err != nil {
		fmt.Println("Error outputting PDF:", err)
		return
	}
	fmt.Println("Printing", tmpFilePath)

	switch runtime.GOOS {
		case "windows":
			PrintWindows(printerName, tmpFilePath)
		case "darwin":
			ClearPrintQueMacOSX(printerName)
			PrintMacOSX(printerName, tmpFilePath, 2)
		case "linux":
			ClearPrintQueLinux(printerName)
			PrintLinux(printerName, tmpFilePath)
		default:
			fmt.Println("Unsupported OS")
	}

	fmt.Println("Printed ticket for", name)
}