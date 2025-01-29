package routes

import (
	"fmt"
	hex "encoding/hex"
	json "encoding/json"
	base64 "encoding/base64"
	bolt "github.com/boltdb/bolt"
	secretbox "golang.org/x/crypto/nacl/secretbox"
	fiber "github.com/gofiber/fiber/v2"
	user "github.com/0187773933/CheckInServer/v1/user"
	server "github.com/0187773933/GO_SERVER/v1/server"
	encryption "github.com/0187773933/encryption/v1/encryption"
)

func UserNewForm( s *server.Server ) fiber.Handler {
	return func( c *fiber.Ctx ) error {
		c.Set( "Content-Type" , "text/html" )
		return c.SendFile( "./v1/html/admin/user_new.html" )
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
			"kp": KyberPublic ,
		})
	}
}

func UserEdit( s *server.Server ) fiber.Handler {
	return func( c *fiber.Ctx ) error {

		body := c.Body()
		// fmt.Println( "Raw Request Body:" , string( body ) )
		var response map[string]interface{}
		if err := json.Unmarshal( body , &response ); err != nil {
			fmt.Println( "Error parsing JSON:", err)
			return c.Status( fiber.StatusBadRequest ).JSON(fiber.Map{
				"error": "Failed to parse JSON request",
			})
		}
		encrypted_user_b64_string , ok := response[ "w" ].( string )
		if !ok {
			fmt.Println( "Error parsing JSON: w" , ok )
			return c.Status( fiber.StatusBadRequest ).JSON(fiber.Map{
				"error": "Failed to parse JSON request",
			})
		}
		user_uuid , ok := response[ "u" ].( string )
		if !ok {
			fmt.Println( "Error parsing JSON: u" , ok )
			return c.Status( fiber.StatusBadRequest ).JSON(fiber.Map{
				"error": "Failed to parse JSON request",
			})
		}

		cipher_text_string , ok := response[ "ct" ].( string )
		if !ok {
			fmt.Println( "Error parsing JSON: ct" , ok )
			return c.Status( fiber.StatusBadRequest ).JSON(fiber.Map{
				"error": "Failed to parse JSON request",
			})
		}
		cipher_text_hex , _ := hex.DecodeString( cipher_text_string )
		var cipher_text_bytes [ 1568 ]byte
		copy( cipher_text_bytes[ : ] , cipher_text_hex )
		shared_secret := encryption.KyberDecrypt( cipher_text_bytes , KyberPrivate )
		fmt.Println( shared_secret )

		encrypted_user , _ := base64.StdEncoding.DecodeString( encrypted_user_b64_string )

		var decrypted_user user.User
		db_result := s.DB.View( func( tx *bolt.Tx ) error {
			users_blank_bucket := tx.Bucket( []byte( "users-blank" ) )
			outer_key := users_blank_bucket.Get( []byte( user_uuid ) )

			var outer_key_array [32]byte
			copy( outer_key_array[ : ] , outer_key )
			var nonce [24]byte
			copy( nonce[ : ] , encrypted_user[ 0 : 24 ] )
			outer_decrypted , ok := secretbox.Open( nil , encrypted_user[ 24 : ] , &nonce , &outer_key_array )
			if ok != true { fmt.Println( "Error decrypting user:" ) }

			var inner_nonce [24]byte
			copy( inner_nonce[ : ] , outer_decrypted[ 0 : 24 ] )
			decrypted , ok := secretbox.Open( nil , outer_decrypted[ 24 : ] , &inner_nonce , &shared_secret )
			if ok != true { fmt.Println( "Error decrypting user:" ) }

			// fmt.Println( string( decrypted ) )

			// decrypted_bucket_value := encrypt.ChaChaDecryptBytes( u.Config.BoltDBEncryptionKey , x_user )
			json.Unmarshal( decrypted , &decrypted_user )
			return nil
		})
		fmt.Println( decrypted_user )
		if db_result != nil {
			fmt.Println( "Error saving blank user:", db_result )
			return c.Status( fiber.StatusInternalServerError ).JSON(fiber.Map{
				"error": "Failed to save blank user",
			})
		}
		return c.JSON( fiber.Map{
			"blank": "" ,
		})
	}
}