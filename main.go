package main

import (
	"os"
	"os/signal"
	"syscall"
	"time"
	"strings"
	bolt "github.com/boltdb/bolt"
	fiber "github.com/gofiber/fiber/v2"
	cors "github.com/gofiber/fiber/v2/middleware/cors"
	logger "github.com/0187773933/Logger/v1/logger"
	utils "github.com/0187773933/GO_SERVER/v1/utils"
	server "github.com/0187773933/GO_SERVER/v1/server"
	routes "github.com/0187773933/CheckInServer/v1/routes"
)

var s server.Server
var DB *bolt.DB

func SetupCloseHandler() {
	c := make( chan os.Signal )
	signal.Notify( c , os.Interrupt , syscall.SIGTERM , syscall.SIGINT )
	go func() {
		<-c
		logger.Log.Println( "\r- Ctrl+C pressed in Terminal" )
		DB.Close()
		logger.Log.Printf( "Shutting Down %s Server" , s.Config.Name )
		s.FiberApp.Shutdown()
		logger.CloseDB()
		os.Exit( 0 )
	}()
}

func main() {
	config := utils.GetConfig()
	// utils.GenerateNewKeysWrite( &config )
	defer utils.SetupStackTraceReport()
	logger.New( &config.Log )
	DB , _ = bolt.Open( config.Bolt.Path , 0600 , &bolt.Options{ Timeout: ( 3 * time.Second ) } )
	s = server.New( &config , logger.Log , DB )
	allow_origins_string := strings.Join( config.AllowOrigins , "," )
	s.FiberApp.Use( cors.New( cors.Config{
		AllowOrigins: allow_origins_string ,
		AllowMethods: "GET, POST, PUT, DELETE, OPTIONS" ,
		AllowHeaders: "Origin, Content-Type, Accept, Authorization, k" ,
	}))
	s.FiberApp.Use( func( c *fiber.Ctx ) error {
		c.Set( "Cache-Control" , "no-store, no-cache, must-revalidate, proxy-revalidate" )
		c.Set( "Pragma" , "no-cache" )
		c.Set( "Expires" , "0" )
		return c.Next()
	})
	routes.SetupPublicRoutes( &s )
	routes.SetupAdminRoutes( &s )
	SetupCloseHandler()
	s.Start()
}