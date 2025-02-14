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

// //go:embed SumatraPDF.exe
// var sumatra_binary []byte

// // Define the known location to store SumatraPDF.exe
// func get_persistent_sumatra_path() string {
// 	appData := os.Getenv( "APPDATA" ) // Example: C:\Users\Username\AppData\Roaming
// 	if appData == "" {
// 		appData = "C:\\ProgramData" // Fallback if APPDATA is not set
// 	}
// 	dirPath := filepath.Join( appData , "CheckInServer" ) // Custom folder for the application
// 	exePath := filepath.Join( dirPath , "SumatraPDF.exe" )
// 	return exePath
// }

// // Ensure SumatraPDF.exe is present in the known location
// func EnsureSumatraPDF() (string, error) {
// 	exePath := get_persistent_sumatra_path()

// 	// Check if the file already exists
// 	if _, err := os.Stat(exePath); err == nil {
// 		return exePath, nil // File already exists, no need to extract again
// 	}

// 	// Create directory if it doesn't exist
// 	if err := os.MkdirAll( filepath.Dir(exePath), 0755); err != nil {
// 		return "", err
// 	}

// 	// Write the embedded binary to the persistent location
// 	err := os.WriteFile(exePath, sumatraBinary, 0755)
// 	if err != nil {
// 		return "", err
// 	}

// 	return exePath, nil
// }

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
	sumatra_file_path  , _ := filepath.Abs( "./v1/embed/misc/SumatraPDF.exe" )
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

func Print( printer_name string , name string ) {

	page_width := 2.25
	page_heigth := 4.0
	page_center_x := ( page_width / 2 )
	page_center_y := ( page_heigth / 2 )
	fmt.Printf( "Page Center X === %v || Page Center Y === %v\n" , page_center_x , page_center_y )

	icon_scale_factor := 0.70
	icon_size_x := ( ( 107.0 / 72.0 ) * icon_scale_factor )
	icon_size_y := ( ( 98.0 / 72.0 ) * icon_scale_factor )
	icon_center_x := ( icon_size_x / 2 )
	icon_center_y := ( icon_size_y / 2 )
	fmt.Printf( "Icon Center X === %v || Icon Center Y === %v\n" , icon_center_x , icon_center_y )

	logo_scale_factor := 0.6
	logo_size_x := ( ( 300 / 96 ) * logo_scale_factor )
	logo_size_y := ( ( 156 / 96 ) * logo_scale_factor )
	logo_center_x := ( logo_size_x / 2 )
	logo_center_y := ( logo_size_y / 2 )
	fmt.Printf( "Logo Center X === %v || Logo Center Y === %v\n" , logo_center_x , logo_center_y )

	pdf := gofpdf.NewCustom( &gofpdf.InitType{
		OrientationStr: "P" , // L = landscape , P = portrait
		UnitStr: "in" ,
		Size: gofpdf.SizeType{ Wd: 2.25 , Ht: 4.0 } ,
	})
	pdf.SetMargins( 0.1 , 0.1 , 0.1 ) // left , top , right
	pdf.AddPage()
	// font_abs_path , _ := filepath.Abs( font_path )


	pdf.AddUTF8Font( "ComicNeue" , "" , "./v1/embed/fonts/ComicNeue-Regular.ttf" )
	// pdf.AddUTF8FontFromBytes( "ComicNeue" , "" , font_bytes )


	add_rotated_centered_text( pdf , name , "ComicNeue" , 40 , ( ( page_center_x - ( page_center_x / 2 ) ) + ( page_center_x / 4 ) ) )
	// add_rotated_image( pdf , "icon.png" , ( page_center_x - icon_center_x - 0.1 ) , ( page_center_y + icon_center_y ) , icon_size_x , icon_size_y , 90 )
	// add_rotated_image( pdf , "./v1/printer/icon.png" , ( page_center_x - icon_center_x - 0.1 ) , ( page_heigth - 1.1 ) , icon_size_x , icon_size_y , 90 )
	add_rotated_image( pdf , "./v1/embed/images/logo.png" , ( page_center_x - ( logo_size_y ) ) , ( ( page_center_y - logo_center_y ) - 0.3 ) , logo_size_x , logo_size_y , 90 )

	pdf_temp_file , _ := ioutil.TempFile( "" , "ticket-*.pdf" )
	defer pdf_temp_file.Close()
	pdf_temp_file_path := pdf_temp_file.Name()
	defer func() {
		os.Remove( pdf_temp_file_path )
	}()
	err := pdf.OutputFileAndClose( pdf_temp_file_path )
	fmt.Println( "Printing" , pdf_temp_file_path )
	if err != nil {
		fmt.Println( err )
		return
	}

	if runtime.GOOS == "windows" {
		// clear_printer_que_windows( config.PrinterName )
		PrintWindows( printer_name , pdf_temp_file_path )
	} else if runtime.GOOS == "darwin" {
		ClearPrintQueMacOSX( printer_name )
		PrintMacOSX( printer_name , pdf_temp_file_path , 2 )
	} else if runtime.GOOS == "linux" {
		ClearPrintQueLinux( printer_name )
		PrintLinux( printer_name , pdf_temp_file_path )
	}

	fmt.Println( "printed ticket for" , name )

}