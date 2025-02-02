package utils

import (
	"io"
	b64 "encoding/base64"
	sha256 "crypto/sha256"
	hkdf "golang.org/x/crypto/hkdf"
)

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