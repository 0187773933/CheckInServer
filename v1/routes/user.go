package routes

import (
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
	temp_key := encryption.SecretBoxGenerateRandomKey()
	temp_user := user.New()
	return func( c *fiber.Ctx ) error {
		return c.JSON( fiber.Map{
			"blank": temp_user ,
			"ecb": temp_key ,
		})
	}
}

func UserEdit( s *server.Server ) fiber.Handler {
	return func( c *fiber.Ctx ) error {
		return c.JSON( fiber.Map{
			"blank": user.New() ,
		})
	}
}