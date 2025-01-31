package printer

import (
	"fmt"
	// _ "embed"
	gofpdf "github.com/jung-kurt/gofpdf"
	// // "bufio"
	// "runtime"
	// "io/ioutil"
	// "os"
	// "os/exec"
	// // "reflect"
	// "path/filepath"
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

func PrintWindows( printer_name string , pdf_file_path string ) {
	sumatra_file_path  , _ := filepath.Abs( "./v1/printer/SumatraPDF.exe" )
	args := []string{ sumatra_file_path , "-print-to" , printer_name , pdf_file_path }
	cmd := exec.Command( args[ 0 ] , args[ 1 : ]... )
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		log.Error( "Error Printing PDF" , err )
	}
}