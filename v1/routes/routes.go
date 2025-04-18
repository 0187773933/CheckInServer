package routes

import (
	"fmt"
	"time"
	// hex "encoding/hex"
	base64 "encoding/base64"
	fiber "github.com/gofiber/fiber/v2"
	server "github.com/0187773933/GO_SERVER/v1/server"
	encryption "github.com/0187773933/encryption/v1/encryption"
	rate_limiter "github.com/gofiber/fiber/v2/middleware/limiter"
)

// var KyberPrivate [3168]byte
// var KyberPublic [1568]byte
// var KyberPrivateString string
// var KyberPublicString string

var X25519Private [32]byte
var X25519Public [32]byte
var X25519PrivateB64String string
var X25519PublicB64String string

func PublicMaxedOut( c *fiber.Ctx ) error {
	ip_address := c.IP()
	log_message := fmt.Sprintf( "%s === %s === %s === PUBLIC RATE LIMIT REACHED !!!" , ip_address , c.Method() , c.Path() );
	fmt.Println( log_message )
	c.Set( "Content-Type" , "text/html" )
	return c.SendString( "<html><h1>loading ...</h1><script>setTimeout(function(){ window.location.reload(1); }, 6000);</script></html>" )
}
var PublicLimter = rate_limiter.New( rate_limiter.Config{
	Max: 3 ,
	Expiration: 1 * time.Second ,
	KeyGenerator: func( c *fiber.Ctx ) string {
		return c.Get( "x-forwarded-for" )
	} ,
	LimitReached: PublicMaxedOut ,
	LimiterMiddleware: rate_limiter.SlidingWindow{} ,
})

func SetupPublicRoutes( s *server.Server ) {
	prefix_string := "/"
	if s.Config.URLS.Prefix != "" {
		prefix_string = s.Config.URLS.Prefix
	}
	prefix := s.FiberApp.Group( prefix_string )

	prefix.Get( "/test/:filename" , func( c *fiber.Ctx ) error {
		fmt.Println( c.Params( "filename" ) )
		file := "./v1/cdn/" + c.Params("filename")
		c.Set( "Content-Type" , "application/javascript" )
		fmt.Println( file )
		return c.SendFile( file )
	})

	prefix.Get( "/" , PublicLimter , func( c *fiber.Ctx ) error {
		return c.JSON( fiber.Map{
			"result": true ,
			"url": "/" ,
		})
	})
	// s.FiberApp.Get( "/twitch" , func( c *fiber.Ctx ) error {
	// 	c.Set( "Content-Type" , "text/html" )
	// 	return c.SendFile( "./v1/html/twitch.html" )
	// })
	// s.FiberApp.Post( "/update_position" , PublicLimter , UpdatePosition( s ) )
	// prefix := s.FiberApp.Group( s.Config.URLS.Prefix )
	// prefix.Get( "/:uuid.:ext" , UUIDFileLimter , ServeFile( s ) )

	// // youtube
	// youtube := prefix.Group( "/youtube" )
	// youtube.Get( "/:library_key/:session_id" , YouTubeSessionHTMLPlayer( s ) )

	// // local library
	// library := prefix.Group( "/library" )
	// library.Get( "/get/entries" , LibraryGetEntries( s ) )
	// // library-session
	// prefix.Use( PublicLimter )
	// prefix.Get( "/:library_key/:session_id/reset" , SessionReset( s ) )
	// prefix.Get( "/:library_key/:session_id/total" , SessionTotal( s ) )
	// prefix.Get( "/:library_key/:session_id/index" , SessionIndex( s ) )
	// prefix.Get( "/:library_key/:session_id/set/index/:index" , SessionSetIndex( s ) )
	// prefix.Get( "/:library_key/:session_id/previous" , SessionPrevious( s ) )
	// prefix.Get( "/:library_key/:session_id/next" , SessionNext( s ) )
	// prefix.Get( "/:library_key/:session_id" , SessionHTMLPlayer( s ) ) // HTML Player
	// prefix.Get( "/:library_key/:session_id/:index" , SessionHTMLPlayerAtIndex( s ) ) // HTML Player at Session Index ?
}

func GetPK( s *server.Server ) fiber.Handler {
	return func( c *fiber.Ctx ) error {
		return c.JSON( fiber.Map{
			"pk": X25519PublicB64String ,
		})
	}
}

func SetupAdminRoutes( s *server.Server ) {
	prefix_string := "/"
	if s.Config.URLS.AdminPrefix != "" {
		prefix_string = s.Config.URLS.AdminPrefix
	}
	admin := s.FiberApp.Group( prefix_string )
	admin.Get( "/" , PublicLimter , func( c *fiber.Ctx ) error {
		return c.JSON( fiber.Map{
			"result": true ,
			"url": "/" ,
		})
	})
	admin.Use( s.ValidateAdminMW )

 	// KyberPublic , KyberPrivate = encryption.KyberGenerateKeyPair()
	// KyberPrivateString = hex.EncodeToString( KyberPrivate[ : ] )
	// KyberPublicString = hex.EncodeToString( KyberPublic[ : ] )

	X25519Public , X25519Private = encryption.CurveX25519GenerateKeyPair()
	X25519PublicB64String = base64.StdEncoding.EncodeToString( X25519Public[ : ] )
	X25519PrivateB64String = base64.StdEncoding.EncodeToString( X25519Private[ : ] )

	admin.Get( "/pk" , GetPK( s ) )
	admin.Get( "/user/new" , UserNewForm( s ) )
	admin.Get( "/user/blank" , UserBlank( s ) )
	admin.Get( "/user/checkin" , UserCheckInForm( s ) )
	admin.Post( "/user/checkin" , UserCheckIn( s ) )
	admin.Post( "/user/edit" , UserEdit( s ) )
	admin.Get( "/user/edit/:uuid" , UserEditForm( s ) ) // TODO : replies with deencrypted
	admin.Get( "/user/edit/" , UserEditForm( s ) ) // TODO : replies with deencrypted
	admin.Get( "/user/get/:uuid" , UserGet( s ) )
	admin.Get( "/user/get/barcode/:barcode" , UserGetByBarcode( s ) )
	admin.Get( "/user/search/:search_term" , UserSearch( s ) )
	// admin.Post( "/user/delete" , UserDelete( s ) )
	admin.Get( "/checkins/:date" , GetCheckInsOnDate( s ) )
}