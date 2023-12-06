package midware

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/incompetent-developer/restgo-struct/header"
)

// AllowRefererSkipper is skip when referer is not allowed
func AllowRefererSkipper(referers ...string) echo.MiddlewareFunc {

	return allowSkipper(header.Referer, referers...)

}

// AllowOriginSkipper is skip when origin is not allowed
func AllowOriginSkipper(origins ...string) echo.MiddlewareFunc {

	return allowSkipper(header.Origin, origins...)

}

func allowSkipper(src string, allowers ...string) echo.MiddlewareFunc {

	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			req := c.Request()
			source := req.Header.Get(src)

			var (
				isAllowed bool
			)

			// Check allowed origins
			for _, allower := range allowers {
				if allower == "*" || allower == source {
					isAllowed = true
					break
				}
			}

			if !isAllowed {
				return c.NoContent(http.StatusForbidden)
			}

			return next(c)
		}
	}

}
