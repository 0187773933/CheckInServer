package routes

import (
	"fmt"
	"io"
	// hex "encoding/hex"
	filepath "path/filepath"
	json "encoding/json"
	sha256 "crypto/sha256"
	hkdf "golang.org/x/crypto/hkdf"
	// uuid "github.com/satori/go.uuid"
	// base64 "encoding/base64"
	bolt "github.com/boltdb/bolt"
	bleve "github.com/blevesearch/bleve/v2"
	secretbox "golang.org/x/crypto/nacl/secretbox"
	fiber "github.com/gofiber/fiber/v2"
	user "github.com/0187773933/CheckInServer/v1/user"
	utils "github.com/0187773933/CheckInServer/v1/utils"
	server "github.com/0187773933/GO_SERVER/v1/server"
	encryption "github.com/0187773933/encryption/v1/encryption"
	// kyberk2so "github.com/symbolicsoft/kyber-k2so"
)

func UserNewForm( s *server.Server ) fiber.Handler {
	return func( c *fiber.Ctx ) error {
		c.Set( "Content-Type" , "text/html" )
		return c.SendFile( "./v1/html/admin/user_new.html" )
	}
}

func UserEditForm( s *server.Server ) fiber.Handler {
	return func( c *fiber.Ctx ) error {
		c.Set( "Content-Type" , "text/html" )
		return c.SendFile( "./v1/html/admin/user_edit.html" )
	}
}

func UserBlank( s *server.Server ) fiber.Handler {
	return func( c *fiber.Ctx ) error {
		temp_key := encryption.SecretBoxGenerateRandomKey()
		temp_user := user.New()
		db_result := s.DB.Update( func( tx *bolt.Tx ) error {
			users_blank_bucket , _ := tx.CreateBucketIfNotExists( []byte( "users-blank" ) )
			users_blank_bucket.Put( []byte( temp_user.UUID ) , temp_key[ : ] )
			return nil
		})
		if db_result != nil {
			fmt.Println( "Error saving blank user:", db_result )
			return c.Status( fiber.StatusInternalServerError ).JSON(fiber.Map{
				"error": "Failed to save blank user",
			})
		}
		return c.JSON( fiber.Map{
			"blank": temp_user ,
			"ecb": temp_key ,
			"pk": X25519PublicB64String ,
		})
	}
}

type BleveUser struct {
	UUID  string `json:"uuid"`
	Username  string `json:"username"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

func NewBleveUserIndex( index_path string ) ( bleve.Index ) {
	// Define mapping for User type
	mapping := bleve.NewIndexMapping()
	docMapping := bleve.NewDocumentMapping()

	// Create field mappings
	textFieldMapping := bleve.NewTextFieldMapping()
	textFieldMapping.Store = true
	textFieldMapping.IncludeInAll = true

	// Add field mappings to document mapping
	docMapping.AddFieldMappingsAt( "username" , textFieldMapping)
	docMapping.AddFieldMappingsAt( "first_name" , textFieldMapping )
	docMapping.AddFieldMappingsAt( "last_name" , textFieldMapping )

	// Add document mapping to index
	mapping.AddDocumentMapping("user", docMapping)

	// Create or open index
	index, err := bleve.New(index_path, mapping)
	if err != nil {
		// If index exists, try to open it
		if err == bleve.ErrorIndexPathExists {
			index, err = bleve.Open( index_path )
			if err != nil {
				return nil
			}
		} else {
			return nil
		}
	}

	return index
}

func deriveKey( sharedSecret []byte ) []byte {
	hkdfReader := hkdf.New( sha256.New , sharedSecret , nil , nil )
	derivedKey := make( []byte , 32 )
	io.ReadFull( hkdfReader , derivedKey )
	return derivedKey
}

func UserEdit( s *server.Server ) fiber.Handler {
	return func( c *fiber.Ctx ) error {

		body := c.Body()
		// fmt.Println( "Raw Request Body:" , string( body ) )
		var response map[string]interface{}
		if err := json.Unmarshal( body , &response ); err != nil {
			fmt.Println( "Error parsing JSON:", err)
			return c.Status( fiber.StatusBadRequest ).JSON( fiber.Map{
				"error": "Failed to parse JSON request",
			})
		}
		encrypted_user_b64_string , ok := response[ "w" ].( string )
		if !ok {
			fmt.Println( "Error parsing JSON: w" , ok )
			return c.Status( fiber.StatusBadRequest ).JSON( fiber.Map{
				"error": "Failed to parse JSON request",
			})
		}
		user_pk_b64_string , ok := response[ "pk" ].( string )
		if !ok {
			fmt.Println( "Error parsing JSON: pk" , ok )
			return c.Status( fiber.StatusBadRequest ).JSON( fiber.Map{
				"error": "Failed to parse JSON request",
			})
		}
		user_pk_bytes := utils.ConvertB64StringToBytes( user_pk_b64_string )
		var user_pk_bytes_32 [32]byte
		copy( user_pk_bytes_32[ : ] , user_pk_bytes )

		// kyber version
		// var cipher_text_bytes [ 1568 ]byte
		// copy( cipher_text_bytes[ : ] , cipher_text_hex )
		// // shared_secret := encryption.KyberDecrypt( cipher_text_bytes , KyberPrivate )
		// shared_secret , err := kyberk2so.KemDecrypt1024( cipher_text_bytes , KyberPrivate )
		// fmt.Println( err )
		// fmt.Println( shared_secret )

		// x25519 version
		// var cipher_text_bytes [ 32 ]byte
		shared_secret := encryption.CurveX25519KeyExchange( X25519Private , user_pk_bytes_32 )
		derived_key := deriveKey( shared_secret[ : ] )
		var derived_key_32 [ 32 ]byte
		copy( derived_key_32[ : ] , derived_key )

		// Decrypt User
		encrypted_user_bytes := utils.ConvertB64StringToBytes( encrypted_user_b64_string )
		var nonce [ 24 ]byte
		copy( nonce[ : ] , encrypted_user_bytes[ 0 : 24 ] )
		decrypted_user_json , ok := secretbox.Open( nil , encrypted_user_bytes[ 24 : ] , &nonce , &derived_key_32 )
		if !ok {
			fmt.Println( "Error decrypting user:" )
			return c.Status( fiber.StatusInternalServerError ).JSON( fiber.Map{
				"error": "Failed to decrypt user",
			})
		}
		var decrypted_user user.User
		json.Unmarshal( decrypted_user_json , &decrypted_user )
		fmt.Println( decrypted_user )

		// update bleve index
		bleve_path := filepath.Join( s.Config.SaveFilesPath , s.Config.MiscMap[ "bleve_path" ] )
		bleve_index := NewBleveUserIndex( bleve_path )
		defer bleve_index.Close()
		user1 := BleveUser{
			UUID: decrypted_user.UUID ,
			Username: decrypted_user.Username ,
		}
		if len( decrypted_user.FamilyMembers ) > 0 {
			user1.FirstName = decrypted_user.FamilyMembers[ 0 ].FirstName
			user1.LastName = decrypted_user.FamilyMembers[ 0 ].LastName
		}
		fmt.Println( "updating bleve index with" )
		fmt.Println( user1 )

		db_result := s.DB.Update( func( tx *bolt.Tx ) error {
			users_blank := tx.Bucket( []byte( "users-blank" ) )
			derived_key_encrypted := encryption.ChaChaEncryptBytes( s.Config.Creds.EncryptionKey , derived_key )
			users_blank.Put( []byte( decrypted_user.UUID ) , derived_key_encrypted )
			users_bucket , _ := tx.CreateBucketIfNotExists( []byte( "users" ) )
			users_bucket.Put( []byte( decrypted_user.UUID ) , []byte( encrypted_user_b64_string ) )
			return nil
		})
		fmt.Println( "Decrypted User:" , decrypted_user )
		if db_result != nil {
			fmt.Println( "Error saving blank user:", db_result )
			return c.Status( fiber.StatusInternalServerError ).JSON( fiber.Map{
				"error": "Failed to save blank user",
			})
		}
		return c.JSON( fiber.Map{
			"result": true ,
		})
	}
}

func UserGet( s *server.Server ) fiber.Handler {
	return func( c *fiber.Ctx ) error {
		uuid := c.Params( "uuid" )
		var decrypted_user user.User
		db_result := s.DB.View( func( tx *bolt.Tx ) error {
			users_blank := tx.Bucket( []byte( "users-blank" ) )
			users_key_encrypted := users_blank.Get( []byte( uuid ) )
			users_key_decrypted := encryption.ChaChaDecryptBytes( s.Config.Creds.EncryptionKey , users_key_encrypted )
			var users_key_32 [ 32 ]byte
			copy( users_key_32[ : ] , users_key_decrypted )
			users_bucket := tx.Bucket( []byte( "users" ) )
			encrypted_user_b64 := users_bucket.Get( []byte( uuid ) )
			encrypted_user_bytes := utils.ConvertB64StringToBytes( string( encrypted_user_b64 ) )
			var nonce [ 24 ]byte
			copy( nonce[ : ] , encrypted_user_bytes[ 0 : 24 ] )
			decrypted_user_json , _ := secretbox.Open( nil , encrypted_user_bytes[ 24 : ] , &nonce , &users_key_32 )
			json.Unmarshal( decrypted_user_json , &decrypted_user )
			return nil
		})
		fmt.Println( "Decrypted User:" , decrypted_user )
		if db_result != nil {
			fmt.Println( "Error saving blank user:", db_result )
			return c.Status( fiber.StatusInternalServerError ).JSON( fiber.Map{
				"error": "Failed to save blank user",
			})
		}
		return c.JSON( fiber.Map{
			"user": decrypted_user ,
		})
	}
}
