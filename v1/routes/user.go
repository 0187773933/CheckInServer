package routes

import (
	"fmt"
	"strings"
	"time"
	"crypto/rand"
	"io"
	// hex "encoding/hex"
	filepath "path/filepath"
	json "encoding/json"
	uuid "github.com/satori/go.uuid"
	// base64 "encoding/base64"
	bolt "github.com/boltdb/bolt"
	bleve "github.com/blevesearch/bleve/v2"
	secretbox "golang.org/x/crypto/nacl/secretbox"
	fiber "github.com/gofiber/fiber/v2"
	user "github.com/0187773933/CheckInServer/v1/user"
	printer "github.com/0187773933/CheckInServer/v1/printer"
	utils "github.com/0187773933/CheckInServer/v1/utils"
	server "github.com/0187773933/GO_SERVER/v1/server"
	encryption "github.com/0187773933/encryption/v1/encryption"
	// kyberk2so "github.com/symbolicsoft/kyber-k2so"
)

func UserNewForm( s *server.Server ) fiber.Handler {
	if s.Config.MiscMap[ "mode" ] == "dev" {
		return func( c *fiber.Ctx ) error {
			c.Set( "Content-Type" , "text/html" )
			return c.SendFile( "./v1/embed/html/admin/user_new.html" )
		}
	} else {
		return func( c *fiber.Ctx ) error {
			c.Set( "Content-Type" , "text/html" )
			file_path := "html/admin/user_new.html"
			file , _ := s.EMBEDED.Open( file_path )
			defer file.Close()
			content , _ := io.ReadAll( file )
			return c.SendString( string( content ) )
		}
	}
}

func UserEditForm( s *server.Server ) fiber.Handler {
	if s.Config.MiscMap[ "mode" ] == "dev" {
		return func( c *fiber.Ctx ) error {
			c.Set( "Content-Type" , "text/html" )
			return c.SendFile( "./v1/embed/html/admin/user_edit.html" )
		}
	} else {
		return func( c *fiber.Ctx ) error {
			c.Set( "Content-Type" , "text/html" )
			file_path := "html/admin/user_edit.html"
			file , _ := s.EMBEDED.Open( file_path )
			defer file.Close()
			content , _ := io.ReadAll( file )
			return c.SendString( string( content ) )
		}
	}
}

func UserCheckInForm( s *server.Server ) fiber.Handler {
	if s.Config.MiscMap[ "mode" ] == "dev" {
		return func( c *fiber.Ctx ) error {
			c.Set( "Content-Type" , "text/html" )
			return c.SendFile( "./v1/embed/html/admin/user_checkin.html" )
		}
	} else {
		return func( c *fiber.Ctx ) error {
			c.Set( "Content-Type" , "text/html" )
			file_path := "html/admin/user_checkin.html"
			file , _ := s.EMBEDED.Open( file_path )
			defer file.Close()
			content , _ := io.ReadAll( file )
			return c.SendString( string( content ) )
		}
	}
}

func UserBlank( s *server.Server ) fiber.Handler {
	return func( c *fiber.Ctx ) error {
		temp_key := encryption.SecretBoxGenerateRandomKey()
		temp_user := user.New( s )
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
	ID  string `json:"uuid"`
	Username  string `json:"username"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	ParentID  string `json:"parent_id"`
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
	index , err := bleve.New( index_path , mapping )
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

// func UserEdit( s *server.Server ) fiber.Handler {
// 	return func( c *fiber.Ctx ) error {

// 		body := c.Body()
// 		// fmt.Println( "Raw Request Body:" , string( body ) )
// 		var response map[string]interface{}
// 		if err := json.Unmarshal( body , &response ); err != nil {
// 			fmt.Println( "Error parsing JSON:", err)
// 			return c.Status( fiber.StatusBadRequest ).JSON( fiber.Map{
// 				"error": "Failed to parse JSON request",
// 			})
// 		}
// 		encrypted_user_b64_string , ok := response[ "w" ].( string )
// 		if !ok {
// 			fmt.Println( "Error parsing JSON: w" , ok )
// 			return c.Status( fiber.StatusBadRequest ).JSON( fiber.Map{
// 				"error": "Failed to parse JSON request",
// 			})
// 		}
// 		user_pk_b64_string , ok := response[ "pk" ].( string )
// 		if !ok {
// 			fmt.Println( "Error parsing JSON: pk" , ok )
// 			return c.Status( fiber.StatusBadRequest ).JSON( fiber.Map{
// 				"error": "Failed to parse JSON request",
// 			})
// 		}
// 		user_pk_bytes := utils.ConvertB64StringToBytes( user_pk_b64_string )
// 		var user_pk_bytes_32 [32]byte
// 		copy( user_pk_bytes_32[ : ] , user_pk_bytes )

// 		// kyber version
// 		// var cipher_text_bytes [ 1568 ]byte
// 		// copy( cipher_text_bytes[ : ] , cipher_text_hex )
// 		// // shared_secret := encryption.KyberDecrypt( cipher_text_bytes , KyberPrivate )
// 		// shared_secret , err := kyberk2so.KemDecrypt1024( cipher_text_bytes , KyberPrivate )
// 		// fmt.Println( err )
// 		// fmt.Println( shared_secret )

// 		// x25519 version
// 		// var cipher_text_bytes [ 32 ]byte
// 		shared_secret := encryption.CurveX25519KeyExchange( X25519Private , user_pk_bytes_32 )
// 		derived_key := utils.DeriveKey( shared_secret[ : ] )
// 		var derived_key_32 [ 32 ]byte
// 		copy( derived_key_32[ : ] , derived_key )

// 		// Decrypt User
// 		encrypted_user_bytes := utils.ConvertB64StringToBytes( encrypted_user_b64_string )
// 		var nonce [ 24 ]byte
// 		copy( nonce[ : ] , encrypted_user_bytes[ 0 : 24 ] )
// 		decrypted_user_json , ok := secretbox.Open( nil , encrypted_user_bytes[ 24 : ] , &nonce , &derived_key_32 )
// 		if !ok {
// 			fmt.Println( "Error decrypting user:" )
// 			return c.Status( fiber.StatusInternalServerError ).JSON( fiber.Map{
// 				"error": "Failed to decrypt user",
// 			})
// 		}
// 		var decrypted_user user.User
// 		json.Unmarshal( decrypted_user_json , &decrypted_user )
// 		fmt.Println( "new user with username" , decrypted_user.Username )
// 		fmt.Println( decrypted_user )

// 		// update bleve index
// 		bleve_path := filepath.Join( s.Config.SaveFilesPath , s.Config.MiscMap[ "bleve_path" ] )
// 		bleve_index := NewBleveUserIndex( bleve_path )
// 		defer bleve_index.Close()
// 		user1 := BleveUser{
// 			ID: decrypted_user.UUID ,
// 			Username: decrypted_user.Username ,
// 		}
// 		if len( decrypted_user.FamilyMembers ) > 0 {
// 			user1.FirstName = decrypted_user.FamilyMembers[ 0 ].FirstName
// 			user1.LastName = decrypted_user.FamilyMembers[ 0 ].LastName
// 		}
// 		bleve_index.Index( decrypted_user.UUID , user1 )
// 		fmt.Println( "updating bleve index with" )
// 		fmt.Println( user1 )

// 		db_result := s.DB.Update( func( tx *bolt.Tx ) error {
// 			users_blank := tx.Bucket( []byte( "users-blank" ) )
// 			derived_key_encrypted := encryption.ChaChaEncryptBytes( s.Config.Creds.EncryptionKey , derived_key )
// 			users_blank.Put( []byte( decrypted_user.UUID ) , derived_key_encrypted )
// 			users_bucket , _ := tx.CreateBucketIfNotExists( []byte( "users" ) )
// 			if len( decrypted_user.Barcodes ) > 0 {
// 				barcodes_bucket , _ := tx.CreateBucketIfNotExists( []byte( "users-barcodes" ) )
// 				for _ , barcode := range decrypted_user.Barcodes {
// 					barcodes_bucket.Put( []byte( barcode ) , []byte( decrypted_user.UUID ) )
// 				}
// 			}
// 			users_bucket.Put( []byte( decrypted_user.UUID ) , []byte( encrypted_user_b64_string ) )
// 			return nil
// 		})
// 		fmt.Println( "Decrypted User:" , decrypted_user )
// 		if db_result != nil {
// 			fmt.Println( "Error saving blank user:", db_result )
// 			return c.Status( fiber.StatusInternalServerError ).JSON( fiber.Map{
// 				"error": "Failed to save blank user",
// 			})
// 		}
// 		return c.JSON( fiber.Map{
// 			"result": true ,
// 		})
// 	}
// }

func UserEdit(s *server.Server) fiber.Handler {
	return func(c *fiber.Ctx) error {
		body := c.Body()
		// fmt.Println("Raw Request Body:", string(body))
		var response map[string]interface{}
		if err := json.Unmarshal(body, &response); err != nil {
			fmt.Println("Error parsing JSON:", err)
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Failed to parse JSON request",
			})
		}
		encryptedUserB64, ok := response["w"].(string)
		if !ok {
			fmt.Println("Error parsing JSON: w", ok)
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Failed to parse JSON request",
			})
		}
		userPkB64, ok := response["pk"].(string)
		if !ok {
			fmt.Println("Error parsing JSON: pk", ok)
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Failed to parse JSON request",
			})
		}
		userPkBytes := utils.ConvertB64StringToBytes(userPkB64)
		var userPkBytes32 [32]byte
		copy(userPkBytes32[:], userPkBytes)

		// X25519 Version: perform key exchange using the server's private key and provided public key.
		sharedSecret := encryption.CurveX25519KeyExchange(X25519Private, userPkBytes32)
		derivedKey := utils.DeriveKey(sharedSecret[:])
		var derivedKey32 [32]byte
		copy(derivedKey32[:], derivedKey)

		// Decrypt the user: first 24 bytes are nonce.
		encryptedUserBytes := utils.ConvertB64StringToBytes(encryptedUserB64)
		var nonce [24]byte
		copy(nonce[:], encryptedUserBytes[0:24])
		decryptedUserJSON, ok := secretbox.Open(nil, encryptedUserBytes[24:], &nonce, &derivedKey32)
		if !ok {
			fmt.Println("Error decrypting user:")
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Failed to decrypt user",
			})
		}

		// Unmarshal into the new user structure.
		var decryptedUser user.User
		if err := json.Unmarshal(decryptedUserJSON, &decryptedUser); err != nil {
			fmt.Println("Error unmarshaling decrypted user:", err)
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Failed to parse decrypted user",
			})
		}
		fmt.Println("New user with username:", decryptedUser.Username)
		// fmt.Println(decryptedUser)

		// --- Ensure UUIDs are set on all Person fields ---
		// If Spouse is provided, ensure they have a UUID.
		if (decryptedUser.Spouse.FirstName != "" || decryptedUser.Spouse.LastName != "") && decryptedUser.Spouse.UUID == "" {
			decryptedUser.Spouse.UUID = uuid.NewV4().String()
		}
		// For each child, ensure a UUID.
		for i, child := range decryptedUser.Children {
			if child.FirstName != "" && child.UUID == "" {
				decryptedUser.Children[i].UUID = uuid.NewV4().String()
			}
		}

		fmt.Println("Decrypted User:", decryptedUser)

		// --- Update Bleve Index ---
		blevePath := filepath.Join(s.Config.SaveFilesPath, s.Config.MiscMap["bleve_path"])
		bleveIndex := NewBleveUserIndex(blevePath)
		defer bleveIndex.Close()

		// Create a main index entry for the primary adult.
		// (We use the Identity field as the primary account.)
		mainIndex := BleveUser{
			ID:        decryptedUser.UUID,
			// Username:  decryptedUser.Username,
			FirstName: decryptedUser.Identity.FirstName,
			LastName:  decryptedUser.Identity.LastName,
		}
		bleveIndex.Index(decryptedUser.UUID, mainIndex)
		fmt.Println("Indexing main user:")
		fmt.Println(mainIndex)

		// Create a Bleve index entry for the Spouse (if provided).
		if decryptedUser.Spouse.FirstName != "" || decryptedUser.Spouse.LastName != "" {
			spouseIndex := BleveUser{
				ID:        decryptedUser.Spouse.UUID ,
				// Username:  decryptedUser.Spouse.FirstName + " " + decryptedUser.Spouse.LastName,
				FirstName: decryptedUser.Spouse.FirstName ,
				LastName:  decryptedUser.Spouse.LastName ,
				ParentID:  decryptedUser.UUID ,
			}
			bleveIndex.Index( decryptedUser.Spouse.UUID , spouseIndex )
			fmt.Println( "Indexing spouse:" )
			fmt.Println( spouseIndex )
		}

		// Create Bleve index entries for each child.
		for _, child := range decryptedUser.Children {
			if child.FirstName != "" {
				childIndex := BleveUser{
					ID:        child.UUID,
					// Username:  child.FirstName + " " + child.LastName,
					FirstName: child.FirstName,
					LastName:  child.LastName,
					ParentID:  decryptedUser.UUID,
				}
				bleveIndex.Index(child.UUID, childIndex)
				fmt.Println("Indexing child:")
				fmt.Println(childIndex)
			}
		}

		// --- Update Database ---
		dbResult := s.DB.Update(func(tx *bolt.Tx) error {
			usersBlank := tx.Bucket([]byte("users-blank"))
			derivedKeyEncrypted := encryption.ChaChaEncryptBytes(s.Config.Creds.EncryptionKey, derivedKey)
			usersBlank.Put([]byte(decryptedUser.UUID), derivedKeyEncrypted)
			usersBucket, _ := tx.CreateBucketIfNotExists([]byte("users"))
			if len(decryptedUser.Barcodes) > 0 {
				barcodesBucket, _ := tx.CreateBucketIfNotExists([]byte("users-barcodes"))
				for _, barcode := range decryptedUser.Barcodes {
					barcodesBucket.Put([]byte(barcode), []byte(decryptedUser.UUID))
				}
			}
			usersBucket.Put([]byte(decryptedUser.UUID), []byte(encryptedUserB64))
			return nil
		})
		fmt.Println("Decrypted User:", decryptedUser)
		if dbResult != nil {
			fmt.Println("Error saving blank user:", dbResult)
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Failed to save blank user",
			})
		}
		return c.JSON(fiber.Map{
			"result": true,
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
			"pk": X25519PublicB64String ,
		})
	}
}

// func UserSearch( s *server.Server ) fiber.Handler {
// 	return func( c *fiber.Ctx ) error {
// 		search_term := c.Params( "search_term" )
// 		bleve_path := filepath.Join( s.Config.SaveFilesPath , s.Config.MiscMap[ "bleve_path" ] )
// 		bleve_index := NewBleveUserIndex( bleve_path )
// 		defer bleve_index.Close()

// 		words := strings.Fields( search_term )
// 		fmt.Println( "words" , words )

// 		// Create the main boolean query that will combine all word queries
// 		mainQuery := bleve.NewDisjunctionQuery()

// 		// For each word, create a disjunction query (OR) across all fields
// 		for _ , word := range words {
// 			// Create fuzzy queries for each field
// 			usernameQuery := bleve.NewFuzzyQuery(word)
// 			usernameQuery.SetField("username")
// 			usernameQuery.Fuzziness = 1

// 			firstNameQuery := bleve.NewFuzzyQuery(word)
// 			firstNameQuery.SetField("first_name")
// 			firstNameQuery.Fuzziness = 1

// 			lastNameQuery := bleve.NewFuzzyQuery(word)
// 			lastNameQuery.SetField("last_name")
// 			lastNameQuery.Fuzziness = 1

// 			// Add all field queries directly to main disjunction
// 			mainQuery.AddQuery(usernameQuery)
// 			mainQuery.AddQuery(firstNameQuery)
// 			mainQuery.AddQuery(lastNameQuery)
// 		}

// 		// Create and execute search request
// 		search_request := bleve.NewSearchRequest( mainQuery )
// 		search_request.Fields = []string{ "username" , "first_name" , "last_name" }

// 		search_results , _ := bleve_index.Search( search_request )

// 		var matched_users []user.User
// 		s.DB.View( func( tx *bolt.Tx ) error {
// 			users_blank := tx.Bucket( []byte( "users-blank" ) )
// 			users_bucket := tx.Bucket( []byte( "users" ) )
// 			for _ , hit := range search_results.Hits {
// 				users_key_encrypted := users_blank.Get( []byte( hit.ID ) )
// 				users_key_decrypted := encryption.ChaChaDecryptBytes( s.Config.Creds.EncryptionKey , users_key_encrypted )
// 				var users_key_32 [ 32 ]byte
// 				copy( users_key_32[ : ] , users_key_decrypted )
// 				encrypted_user_b64 := users_bucket.Get( []byte( hit.ID ) )
// 				encrypted_user_bytes := utils.ConvertB64StringToBytes( string( encrypted_user_b64 ) )
// 				var nonce [ 24 ]byte
// 				copy( nonce[ : ] , encrypted_user_bytes[ 0 : 24 ] )
// 				decrypted_user_json , _ := secretbox.Open( nil , encrypted_user_bytes[ 24 : ] , &nonce , &users_key_32 )
// 				var decrypted_user user.User
// 				json.Unmarshal( decrypted_user_json , &decrypted_user )
// 				matched_users = append( matched_users , decrypted_user )
// 			}
// 			return nil
// 		})
// 		return c.JSON( fiber.Map{
// 			"users": matched_users ,
// 		})
// 	}
// }

func UserSearch(s *server.Server) fiber.Handler {
	return func(c *fiber.Ctx) error {
		searchTerm := c.Params("search_term")
		blevePath := filepath.Join(s.Config.SaveFilesPath, s.Config.MiscMap["bleve_path"])
		bleveIndex := NewBleveUserIndex(blevePath)
		defer bleveIndex.Close()

		words := strings.Fields(searchTerm)
		fmt.Println("words", words)

		// Create a main disjunction query that combines queries for each word.
		mainQuery := bleve.NewDisjunctionQuery()

		for _, word := range words {
			// Search in username
			// usernameQuery := bleve.NewFuzzyQuery(word)
			// usernameQuery.SetField("username")
			// usernameQuery.Fuzziness = 1

			// Search in first_name
			firstNameQuery := bleve.NewFuzzyQuery(word)
			firstNameQuery.SetField("first_name")
			firstNameQuery.Fuzziness = 1

			// Search in last_name
			lastNameQuery := bleve.NewFuzzyQuery(word)
			lastNameQuery.SetField("last_name")
			lastNameQuery.Fuzziness = 1

			// Optionally, search in spouse_name (if you index spouse's full name in this field)
			// spouseNameQuery := bleve.NewFuzzyQuery(word)
			// spouseNameQuery.SetField("spouse_name")
			// spouseNameQuery.Fuzziness = 1

			// mainQuery.AddQuery(usernameQuery)
			mainQuery.AddQuery(firstNameQuery)
			mainQuery.AddQuery(lastNameQuery)
			// mainQuery.AddQuery(spouseNameQuery)
		}

		searchRequest := bleve.NewSearchRequest(mainQuery)
		// Request fields that you need to build your user object.
		searchRequest.Fields = []string{"first_name", "last_name", "parent_id"}

		searchResults, err := bleveIndex.Search(searchRequest)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Search error",
			})
		}

		var matched_users []user.User
		// Retrieve the encrypted users from the DB and decrypt them.
		s.DB.View(func(tx *bolt.Tx) error {
			usersBlank := tx.Bucket([]byte("users-blank"))
			usersBucket := tx.Bucket([]byte("users"))
			for _, hit := range searchResults.Hits {
				var parentID string
				// Attempt to retrieve "parent_id" from the stored fields.
				if v, ok := hit.Fields["parent_id"].(string); ok && v != "" {
					parentID = v
				} else {
					parentID = hit.ID
				}

				encryptedKey := usersBlank.Get([]byte(parentID))
				// Make sure encryptedKey is valid before proceeding.
				if encryptedKey == nil {
					continue
				}
				decryptedKey := encryption.ChaChaDecryptBytes(s.Config.Creds.EncryptionKey, encryptedKey)
				var key32 [32]byte
				copy(key32[:], decryptedKey)

				encryptedUserB64 := usersBucket.Get([]byte(parentID))
				encryptedUserBytes := utils.ConvertB64StringToBytes(string(encryptedUserB64))
				var nonce [24]byte
				copy(nonce[:], encryptedUserBytes[0:24])
				decryptedUserJSON, _ := secretbox.Open(nil, encryptedUserBytes[24:], &nonce, &key32)
				var decryptedUser user.User
				json.Unmarshal(decryptedUserJSON, &decryptedUser)
				matched_users = append(matched_users, decryptedUser)
			}
			return nil
		})


		return c.JSON(fiber.Map{
			"users": matched_users,
		})
	}
}


func UserGetByBarcode( s *server.Server ) fiber.Handler {
	return func( c *fiber.Ctx ) error {
		barcode := c.Params( "barcode" )
		fmt.Println( "searching for barcode" , barcode )
		var matched_user user.User
		s.DB.Update( func( tx *bolt.Tx ) error {
			users_barcodes  , _ := tx.CreateBucketIfNotExists( []byte( "users-barcodes" ) )
			user_uuid := users_barcodes.Get( []byte( barcode ) )
			if user_uuid == nil {
				return nil
			}
			fmt.Println( "found user" , string( user_uuid ) )
			users_blank  , _ := tx.CreateBucketIfNotExists( []byte( "users-blank" ) )
			users_bucket , _ := tx.CreateBucketIfNotExists( []byte( "users" ) )
			users_key_encrypted := users_blank.Get( user_uuid )
			users_key_decrypted := encryption.ChaChaDecryptBytes( s.Config.Creds.EncryptionKey , users_key_encrypted )
			var users_key_32 [ 32 ]byte
			copy( users_key_32[ : ] , users_key_decrypted )
			encrypted_user_b64 := users_bucket.Get( user_uuid )
			encrypted_user_bytes := utils.ConvertB64StringToBytes( string( encrypted_user_b64 ) )
			var nonce [ 24 ]byte
			copy( nonce[ : ] , encrypted_user_bytes[ 0 : 24 ] )
			decrypted_user_json , _ := secretbox.Open( nil , encrypted_user_bytes[ 24 : ] , &nonce , &users_key_32 )
			json.Unmarshal( decrypted_user_json , &matched_user )
			return nil
		})
		fmt.Println( matched_user )
		return c.JSON( fiber.Map{
			"user": matched_user ,
		})
	}
}

func UserCheckIn( s *server.Server ) fiber.Handler {
	return func( c *fiber.Ctx ) error {
		// body := c.Body()
		// var response map[string]interface{}
		// if err := json.Unmarshal( body , &response ); err != nil {
		// 	fmt.Println( "Error parsing JSON:", err)
		// 	return c.Status( fiber.StatusBadRequest ).JSON( fiber.Map{
		// 		"error": "Failed to parse JSON request",
		// 	})
		// }

		var check_in user.CheckIn
		if err := c.BodyParser(&check_in); err != nil {
			fmt.Println("Error parsing JSON:", err)
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Failed to parse JSON request",
			})
		}
		fmt.Println( check_in )

		var print_tickets []string
		s.DB.Update( func( tx *bolt.Tx ) error {
			users_blank := tx.Bucket( []byte( "users-blank" ) )
			users_key_encrypted := users_blank.Get( []byte( check_in.UUID ) )
			users_key_decrypted := encryption.ChaChaDecryptBytes( s.Config.Creds.EncryptionKey , users_key_encrypted )
			var users_key_32 [ 32 ]byte
			copy( users_key_32[ : ] , users_key_decrypted )
			users_bucket := tx.Bucket( []byte( "users" ) )
			encrypted_user_b64 := users_bucket.Get( []byte( check_in.UUID ) )
			encrypted_user_bytes := utils.ConvertB64StringToBytes( string( encrypted_user_b64 ) )
			var nonce [ 24 ]byte
			copy( nonce[ : ] , encrypted_user_bytes[ 0 : 24 ] )
			decrypted_user_json , _ := secretbox.Open( nil , encrypted_user_bytes[ 24 : ] , &nonce , &users_key_32 )
			var decrypted_user user.User
			json.Unmarshal( decrypted_user_json , &decrypted_user )
			fmt.Println( decrypted_user )

			// Make Parents Tickets
			print_ticket_parents := decrypted_user.GetFamilyName()
			print_tickets = append( print_tickets , print_ticket_parents )
			// Make Included Child Tickets
			for _ , child := range decrypted_user.Children {
				for _ , included := range check_in.Additional {
					if included == child.UUID {
						var child_ticket string
						if child.LastName == "" {
							child_ticket = child.FirstName + " " + decrypted_user.Identity.LastName
						} else {
							child_ticket = child.FirstName + " " + child.LastName
						}
						print_tickets = append( print_tickets , child_ticket )
					}
				}
			}

			// check-in
			now := s.LOG.GetNowTimeOBJ()
			check_in.Date = s.LOG.FormatDateString( &now )
			check_in.Time = s.LOG.FormatTimeString( &now )
			decrypted_user.CheckIns = append( decrypted_user.CheckIns , check_in )

			// this is our only instance of re-encrytping i think right now
			var new_nonce [ 24 ]byte
			rand.Read( new_nonce[ : ] )
			reenrypted_user_json , _ := json.Marshal( decrypted_user )
			sealed_data := secretbox.Seal( nil , reenrypted_user_json , &new_nonce , &users_key_32 )
			reencrypted_user_bytes := append( new_nonce[ : ] , sealed_data... )
			reencrypted_user_b64 := utils.ConvertBytesToB64String( reencrypted_user_bytes )
			users_bucket.Put( []byte( check_in.UUID ) , []byte( reencrypted_user_b64 ) )
			return nil
		})

		for _ , print_string := range print_tickets {
			fmt.Println( "printing" , print_string )
			printer.Print( s.Config.MiscMap[ "printer_name" ] , s.Config.MiscMap[ "font_path" ] , print_string )
		}

		return c.JSON( fiber.Map{
			"result": true ,
			"check_in": check_in ,
			"print_tickets": print_tickets ,
		})
	}
}

// TODO move , to file
type CheckInResult struct {
	Username string `json:"username"`
	UUID string `json:"uuid"`
	Parents string `json:"parents"`
}
func GetCheckInsOnDate( s *server.Server ) fiber.Handler {
	return func( c *fiber.Ctx ) error {
		date_string := c.Params( "date" )
		date , _ := time.Parse( "02Jan2006" , date_string )

		var check_in_results []CheckInResult
		s.DB.View( func( tx *bolt.Tx ) error {

			users_bucket := tx.Bucket( []byte( "users" ) )
			users_blank := tx.Bucket( []byte( "users-blank" ) )
			users_blank.ForEach( func( k , v []byte ) error {

				users_key_decrypted := encryption.ChaChaDecryptBytes( s.Config.Creds.EncryptionKey , v )
				var users_key_32 [ 32 ]byte
				copy( users_key_32[ : ] , users_key_decrypted )
				encrypted_user_b64 := users_bucket.Get( k )
				encrypted_user_bytes := utils.ConvertB64StringToBytes( string( encrypted_user_b64 ) )
				var nonce [ 24 ]byte
				copy( nonce[ : ] , encrypted_user_bytes[ 0 : 24 ] )
				decrypted_user_json , _ := secretbox.Open( nil , encrypted_user_bytes[ 24 : ] , &nonce , &users_key_32 )
				var decrypted_user user.User
				json.Unmarshal( decrypted_user_json , &decrypted_user )

				if len( decrypted_user.CheckIns ) < 1 { return nil }
				for _ , check_in := range decrypted_user.CheckIns {
					check_in_date , _ := time.Parse( "02Jan2006" , check_in.Date )
					if check_in_date != date { continue }
					var check_in_result CheckInResult
					check_in_result.Username = decrypted_user.Username
					check_in_result.UUID = decrypted_user.UUID
					if len( decrypted_user.FamilyMembers ) > 0 {
						check_in_result.Parents = decrypted_user.FamilyMembers[ 0 ].FirstName + " " + decrypted_user.FamilyMembers[ 0 ].LastName
					}
					check_in_results = append( check_in_results , check_in_result )
				}

				return nil
			})
			return nil
		})

		return c.JSON( fiber.Map{
			"date": date ,
			"checkins": check_in_results ,
		})
	}
}