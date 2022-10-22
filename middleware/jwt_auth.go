package middleware

import (
	"BE-Project/model/web"
	"BE-Project/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthorizeJWTMiddleware interface {
	AuthorizeJWT() gin.HandlerFunc
}

type authorizeJWTMiddleware struct {
	jwtService service.JWTService
}

func NewAuthorizeJWTMiddleware(jwtService  service.JWTService) AuthorizeJWTMiddleware {
	return &authorizeJWTMiddleware{
		jwtService: jwtService,
	}
}

func (m *authorizeJWTMiddleware) AuthorizeJWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			webResponse := web.NewWebErrorResponse(http.StatusUnauthorized, "UNAUTHORIZED", "No token found")
			c.JSON(http.StatusUnauthorized, webResponse)
			c.Abort()
			return
		}
		_, err := m.jwtService.ValidateToken(authHeader)
		if err != nil {
			webResponse := web.NewWebErrorResponse(http.StatusUnauthorized, "UNAUTHORIZED", "Invalid token")
			c.JSON(http.StatusUnauthorized, webResponse)
			c.Abort()
			return
		}
	}
}
