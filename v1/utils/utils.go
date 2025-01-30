package utils

import (
	"fmt"
	"io"
	"time"
	"strings"
	"embed"
	b64 "encoding/base64"
	sha256 "crypto/sha256"
	hkdf "golang.org/x/crypto/hkdf"
)

//go:embed zoneinfo
var ZoneInfoFS embed.FS
var location *time.Location

// https://stackoverflow.com/questions/48439363/missing-location-in-call-to-time-in
// thanks bro
func Get_location( name string ) ( *time.Location ) {
	return location
}

func SetLocation( location_string string ) {
	if location_string == "" { location_string = "America/New_York" }
	bs , err := ZoneInfoFS.ReadFile( "zoneinfo/" + location_string )
	if err != nil { panic( err ) }
	loc , err := time.LoadLocationFromTZData( location_string , bs )
	if err != nil { panic( err ) }
	location = loc
	fmt.Println( "location set to: " , location_string )
}

func FormatTime( input_time *time.Time ) ( result string ) {
	// location , _ := time.LoadLocation( "America/New_York" )
	// location := get_location( "America/New_York" )
	time_object := input_time.In( location )
	month_name := strings.ToUpper( time_object.Format( "Jan" ) )
	milliseconds := time_object.Format( ".000" )
	date_part := fmt.Sprintf( "%02d%s%d" , time_object.Day() , month_name , time_object.Year() )
	time_part := fmt.Sprintf( "%02d:%02d:%02d%s" , time_object.Hour() , time_object.Minute() , time_object.Second() , milliseconds )
	result = fmt.Sprintf( "%s === %s" , date_part , time_part )
	return
}

func GetFormattedTimeString() ( result string ) {
	// location , _ := time.LoadLocation( "America/New_York" )
	// location := get_location( "America/New_York" )
	time_object := time.Now().In( location )
	month_name := strings.ToUpper( time_object.Format( "Jan" ) )
	milliseconds := time_object.Format( ".000" )
	date_part := fmt.Sprintf( "%02d%s%d" , time_object.Day() , month_name , time_object.Year() )
	time_part := fmt.Sprintf( "%02d:%02d:%02d%s" , time_object.Hour() , time_object.Minute() , time_object.Second() , milliseconds )
	result = fmt.Sprintf( "%s === %s" , date_part , time_part )
	return
}

func GetNowTimeOBJ() ( result time.Time ) {
	result = time.Now().In( location )
	return
}

func GetNowDateString( now *time.Time ) ( result string ) {
	month_name := strings.ToUpper( now.Format( "Jan" ) )
	result = fmt.Sprintf( "%02d%s%d" , now.Day() , month_name , now.Year() )
	return
}

func GetNowTimeString( now *time.Time ) ( result string ) {
	milliseconds := now.Format( ".000" )
	result = fmt.Sprintf( "%02d:%02d:%02d%s" , now.Hour() , now.Minute() , now.Second() , milliseconds )
	return
}

func GetFormattedTimeStringOBJ() ( result_string string , result_time time.Time ) {
	result_time = time.Now().In( location )
	month_name := strings.ToUpper( result_time.Format( "Jan" ) )
	milliseconds := result_time.Format( ".000" )
	date_part := fmt.Sprintf( "%02d%s%d" , result_time.Day() , month_name , result_time.Year() )
	time_part := fmt.Sprintf( "%02d:%02d:%02d%s" , result_time.Hour() , result_time.Minute() , result_time.Second() , milliseconds )
	result_string = fmt.Sprintf( "%s === %s" , date_part , time_part )
	return
}

func IsStringInArray( target string , array []string ) ( bool ) {
	for _ , value := range array {
		if value == target {
			return true
		}
	}
	return false
}

func ConvertB64StringToBytes( b64_string string ) ( result []byte ) {
	result , _ = b64.StdEncoding.DecodeString( b64_string )
	return
}

func ConvertBytesToB64String( bytes []byte ) ( result string ) {
	result = b64.StdEncoding.EncodeToString( bytes )
	return
}

func DeriveKey( shared_secret []byte ) []byte {
	hkdf_reader := hkdf.New( sha256.New , shared_secret , nil , nil )
	derived_key := make( []byte , 32 )
	io.ReadFull( hkdf_reader , derived_key )
	return derived_key
}