package printer

import (
	"fmt"
	// _ "embed"
	"runtime"
	"os"
	"os/exec"
	"io/ioutil"
	// "path/filepath"
	gofpdf "github.com/jung-kurt/gofpdf"
	// // "bufio"
	// "runtime"
	// "io/ioutil"
	// "os"
	// "os/exec"
	// // "reflect"
	"path/filepath"
	// "image/png"
	// // "github.com/boombuler/barcode"
	// // "github.com/boombuler/barcode/code128"
	// "github.com/ppsleep/barcode"
	// "github.com/ppsleep/barcode/code128"
)

// https://github.com/0187773933/MastersCloset/blob/master/v1/printer/printer.go

// func ClearPrintQueWindows( printer_name string ) {
// 	fmt.Printn( "todo" )
// }

func add_centered_text( pdf *gofpdf.Fpdf , text string , font_name string , font_size float64 , at_page_height float64 ) {
	page_width , _ := pdf.GetPageSize()
	// margin_left , margin_top , margin_right , margin_bottom := pdf.GetMargins()
	// fmt.Printf( "Page Width === %v || Page Height === %v\n" , page_width , page_height )
	// fmt.Printf( "Margin Left === %v || Margin Right === %v\n" , margin_left , margin_right )
	// fmt.Printf( "Margin Top === %v || Margin Bottom === %v\n" , margin_top , margin_bottom )
	pdf.SetFont( font_name , "" , font_size )
	string_width := pdf.GetStringWidth( text )
	// fmt.Printf( "String Width === %v" , string_width )
	page_center_x := ( page_width / 2 )
	starting_x := ( page_center_x - ( string_width / 2 ) )
	pdf.Text( starting_x , at_page_height , text )
}

func add_rotated_centered_text(pdf *gofpdf.Fpdf, text string, font_name string, font_size float64, at_page_width float64) {
	_, page_height := pdf.GetPageSize() // Get portrait page dimensions

	pdf.SetFont(font_name, "", font_size)
	string_width := pdf.GetStringWidth(text)

	// Compute center of the rotated text (pseudo-landscape)
	page_center_y := page_height / 2
	// starting_y := page_center_y + (string_width / 2) // Shift to align center
	// starting_y := page_center_y + (string_width / 2) // Shift to align center
	starting_y := page_center_y
	starting_x := ( at_page_width - (string_width / 2) ) + 0.3
	// Rotate text by 90 degrees (pseudo-landscape)
	pdf.TransformBegin()
	pdf.TransformRotate(90, at_page_width, page_center_y) // Rotate around (at_page_width, center)
	pdf.Text(starting_x, starting_y, text)             // Draw rotated text
	pdf.TransformEnd()
}

func add_rotated_image(pdf *gofpdf.Fpdf, imagePath string, x, y, width, height, angle float64) {
	// Compute the rotation center (middle of the image)
	centerX := x + (width / 2)
	centerY := y + (height / 2)

	fmt.Printf( "Center X === %v || Center Y === %v\n" , centerX , centerY )

	// Begin transformation
	pdf.TransformBegin()
	pdf.TransformRotate(angle, centerX, centerY) // Rotate around image center

	// Add image (unrotated position)
	pdf.ImageOptions(
		imagePath,
		x, y,
		width, height,
		false,
		gofpdf.ImageOptions{
			ImageType: "PNG",
			ReadDpi: true,
			AllowNegativePosition: false,
		},
		0, "",
	)

	// End transformation
	pdf.TransformEnd()
}

func PrintWindows( printer_name string , pdf_file_path string ) {
	sumatra_file_path  , _ := filepath.Abs( "./v1/embed/windows/SumatraPDF.exe" )
	// sumatra_file_path := "SumatraPDF.exe"
	args := []string{ sumatra_file_path , "-print-to" , printer_name , pdf_file_path }
	cmd := exec.Command( args[ 0 ] , args[ 1 : ]... )
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		fmt.Println( "Error Printing PDF" , err )
	}
}

func ClearPrintQueMacOSX( printer_name string ) {
	cmd := exec.Command( "cancel" , "-a" , printer_name )
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		fmt.Println( "Error Clearing Printer Que:" , err )
	}
}

// lpstat -v
// lp -d "_4BARCODE_4B_2054N" -o PrintSpeed=2 -o fit-to-page "/Users/morpheous/WORKSPACE/GO/TMP2/BarcodePrinterTest/output.pdf"
func PrintMacOSX( printer_name string , pdf_file_path string , print_speed int ) {
	if print_speed < 1 { print_speed = 2 }
	args := []string{ "lp" , "-d" , printer_name , "-o" , fmt.Sprintf( "PrintSpeed=%d" , print_speed ) , "-o" , "fit-to-page" , pdf_file_path }
	cmd := exec.Command( args[ 0 ] , args[ 1 : ]... )
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		fmt.Println( "Error Printing PDF" , err )
	}
}

func ClearPrintQueLinux( printer_name string ) {
	fmt.Println( "todo" )
}

func PrintLinux( printer_name string , pdf_file_path string ) {
	fmt.Println( "todo" )
}


// centerCoords returns the top-left coordinate (x,y) needed to center an element of
// given dimensions (elementWidth x elementHeight) on a page (pageWidth x pageHeight).
func center_coords(pageWidth, pageHeight, elementWidth, elementHeight float64) (x, y float64) {
	return (pageWidth - elementWidth) / 2, (pageHeight - elementHeight) / 2
}

// addRotatedCenteredText draws text rotated by 'angle' degrees about a pivot point.
// It computes the text’s bounding box (using GetStringWidth and approximating height as fontSize/72)
// and then calculates a top-left coordinate so that the bounding box is centered on (pivotX, pivotY).
func add_rotated_centered_text_two(pdf *gofpdf.Fpdf, text, fontName string, fontSize, pivotX, pivotY, angle float64) {
	pdf.SetFont(fontName, "", fontSize)
	textWidth := pdf.GetStringWidth(text)
	// Approximate text height in inches (1pt = 1/72 in)
	textHeight := fontSize / 72.0

	// Compute top-left coordinates so that the text bounding box is centered at (pivotX, pivotY)
	x := pivotX - (textWidth / 2)
	y := pivotY - (textHeight / 2)

	// Rotate about the pivot and then draw the text at the computed (x, y)
	pdf.TransformBegin()
	pdf.TransformRotate(angle, pivotX, pivotY)
	pdf.Text(x, y, text)
	pdf.TransformEnd()
}

// addRotatedImage adds an image rotated by 'angle' degrees about its center.
func add_rotated_image_two(pdf *gofpdf.Fpdf, imagePath string, x, y, width, height, angle float64) {
	// Compute the center of the image.
	centerX := x + width/2
	centerY := y + height/2

	pdf.TransformBegin()
	pdf.TransformRotate(angle, centerX, centerY)
	pdf.ImageOptions(
		imagePath,
		(x-0.1), y,
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

func Print( printer_name , name string) {
	// Page dimensions in inches.
	const pageWidth = 2.25
	const pageHeight = 4.0

	// Create PDF with specified page dimensions.
	pdf := gofpdf.NewCustom(&gofpdf.InitType{
		OrientationStr: "P", // portrait
		UnitStr: "in",
		Size:  gofpdf.SizeType{Wd: pageWidth, Ht: pageHeight},
	})
	pdf.SetMargins(0.1, 0.1, 0.1)
	pdf.AddPage()

	// Add the font (ensure the file exists at fontPath).
	pdf.AddUTF8Font( "ComicNeue" , "" , "./v1/embed/fonts/ComicNeue-Regular.ttf" )

	// --- Add Rotated Centered Text ---
	// We'll mimic your original approach: using a horizontal pivot computed as:
	//   pivotX = (pageCenterX - (pageCenterX/2)) + (pageCenterX/4)
	// and the vertical pivot as the page’s vertical center.
	pageCenterX := pageWidth / 2
	pageCenterY := pageHeight / 2
	pivotX := (pageCenterX - (pageCenterX / 2)) + (pageCenterX / 4) // equivalent to 0.75 * pageCenterX
	pivotY := pageCenterY
	// Rotate the text 90° about (pivotX, pivotY)
	add_rotated_centered_text_two(pdf, name, "ComicNeue", 40, pivotX, pivotY, 90)

	// --- Add Rotated, Centered Logo ---
	// Compute logo dimensions.
	logoScale := 0.6
	logoWidth := (300.0 / 96.0) * logoScale
	logoHeight := (156.0 / 96.0) * logoScale
	// For 90° rotation, the effective dimensions swap.
	effectiveLogoWidth := logoHeight
	effectiveLogoHeight := logoWidth
	// Compute top-left coordinates to center the logo's bounding box.
	logoX, logoY := center_coords(pageWidth, pageHeight, effectiveLogoWidth, effectiveLogoHeight)
	// Adjust vertically to nudge the logo downward.
	logoY += 0.3

	add_rotated_image_two( pdf , "./v1/embed/images/logo.png" , logoX, logoY, logoWidth, logoHeight, 90)

	// --- Output the PDF ---
	pdf_temp_file , _ := ioutil.TempFile( "" , "ticket-*.pdf" )
	defer pdf_temp_file.Close()
	pdf_temp_file_path := pdf_temp_file.Name()
	defer func() {
		os.Remove( pdf_temp_file_path )
	}()
	if err := pdf.OutputFileAndClose(pdf_temp_file_path); err != nil {
		fmt.Println("Error outputting PDF:", err)
		return
	}
	fmt.Println("Printing", pdf_temp_file_path)

	// --- OS-specific printing ---
	switch runtime.GOOS {
		case "windows":
			PrintWindows(printer_name, pdf_temp_file_path)
		case "darwin":
			ClearPrintQueMacOSX(printer_name)
			PrintMacOSX(printer_name, pdf_temp_file_path, 2)
		case "linux":
			ClearPrintQueLinux(printer_name)
			PrintLinux(printer_name, pdf_temp_file_path)
		default:
			fmt.Println("Unsupported OS")
	}

	fmt.Println("Printed ticket for", name)
}

func PrintTwo( printer_name , name string , extra_string string ) {
	// Page dimensions in inches.
	const pageWidth = 2.25
	const pageHeight = 4.0

	// Create PDF with specified page dimensions.
	pdf := gofpdf.NewCustom(&gofpdf.InitType{
		OrientationStr: "P", // portrait
		UnitStr: "in",
		Size:  gofpdf.SizeType{Wd: pageWidth, Ht: pageHeight},
	})
	pdf.SetMargins(0.1, 0.1, 0.1)
	pdf.AddPage()

	// Add the font (ensure the file exists at fontPath).
	pdf.AddUTF8Font( "ComicNeue" , "" , "./v1/embed/fonts/ComicNeue-Regular.ttf" )

	// --- Add Rotated Centered Text ---
	// We'll mimic your original approach: using a horizontal pivot computed as:
	//   pivotX = (pageCenterX - (pageCenterX/2)) + (pageCenterX/4)
	// and the vertical pivot as the page’s vertical center.
	pageCenterX := pageWidth / 2
	pageCenterY := pageHeight / 2
	pivotX := (pageCenterX - (pageCenterX / 2)) + (pageCenterX / 4) // equivalent to 0.75 * pageCenterX
	pivotXExtra := (pageCenterX - (pageCenterX / 2)) + (pageCenterX / 4) + 0.25 // equivalent to 0.75 * pageCenterX
	pivotY := pageCenterY
	// Rotate the text 90° about (pivotX, pivotY)
	add_rotated_centered_text_two(pdf, name, "ComicNeue", 40, pivotX, pivotY, 90)
	add_rotated_centered_text_two(pdf, extra_string, "ComicNeue", 16, pivotXExtra, pivotY, 90)

	// --- Add Rotated, Centered Logo ---
	// Compute logo dimensions.
	logoScale := 0.6
	logoWidth := (300.0 / 96.0) * logoScale
	logoHeight := (156.0 / 96.0) * logoScale
	// For 90° rotation, the effective dimensions swap.
	effectiveLogoWidth := logoHeight
	effectiveLogoHeight := logoWidth
	// Compute top-left coordinates to center the logo's bounding box.
	logoX, logoY := center_coords(pageWidth, pageHeight, effectiveLogoWidth, effectiveLogoHeight)
	// Adjust vertically to nudge the logo downward.
	logoY += 0.3

	add_rotated_image_two( pdf , "./v1/embed/images/logo.png" , logoX, logoY, logoWidth, logoHeight, 90)

	// --- Output the PDF ---
	pdf_temp_file , _ := ioutil.TempFile( "" , "ticket-*.pdf" )
	defer pdf_temp_file.Close()
	pdf_temp_file_path := pdf_temp_file.Name()
	defer func() {
		os.Remove( pdf_temp_file_path )
	}()
	if err := pdf.OutputFileAndClose(pdf_temp_file_path); err != nil {
		fmt.Println("Error outputting PDF:", err)
		return
	}
	fmt.Println("Printing", pdf_temp_file_path)

	// --- OS-specific printing ---
	switch runtime.GOOS {
		case "windows":
			PrintWindows(printer_name, pdf_temp_file_path)
		case "darwin":
			ClearPrintQueMacOSX(printer_name)
			PrintMacOSX(printer_name, pdf_temp_file_path, 2)
		case "linux":
			ClearPrintQueLinux(printer_name)
			PrintLinux(printer_name, pdf_temp_file_path)
		default:
			fmt.Println("Unsupported OS")
	}

	fmt.Println("Printed ticket for", name)
}