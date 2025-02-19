package middlewares

import (
	"net/http"
	
	"github.com/akithepriest/click/services"
	"github.com/labstack/echo/v4"
)

func ProtectedMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		cookie, err := c.Cookie("access_token")
		if err != nil {
			c.Echo().Logger.Error(err)
			c.Redirect(http.StatusSeeOther, "/auth/google/login")
		}

		_, err = services.VerifyToken(cookie.Value)
		if err != nil {
			c.Echo().Logger.Error(err)
			c.Redirect(http.StatusSeeOther, "/auth/google/login")
		}

		return next(c)
	}
}