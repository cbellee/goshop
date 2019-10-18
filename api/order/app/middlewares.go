package app

import (
	"github.com/labstack/echo/middleware"
)

func initMiddlewares(s *server) {
	s.Use(middleware.Logger())
	s.Use(middleware.Recover())
	s.Use(middleware.CORS())
	/* 	s.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost", "https://localhost"},
		AllowHeaders: []string{s.Echo.HeaderOrigin, s.Echo.HeaderContentType, s.Echo.HeaderAccept},
	})) */
	// check list of middleware at https://echo.labstack.com/middleware
}

// Put custom middleware belows
// Example: https://echo.labstack.com/cookbook/middleware
