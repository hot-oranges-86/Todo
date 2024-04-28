package middlewares

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func ConfigureCORS() echo.MiddlewareFunc {
	return middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:5173"},
		AllowMethods: []string{"GET", "POST", "PUT", "DELETE"},
	})
}

func GetMiddlewares() []echo.MiddlewareFunc {
	loggerMiddleware := middleware.Logger()
	recoverMiddleware := middleware.Recover()
	corsMiddleware := ConfigureCORS()

	return []echo.MiddlewareFunc{
		loggerMiddleware, recoverMiddleware, corsMiddleware,
	}
}
