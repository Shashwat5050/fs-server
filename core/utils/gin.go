package utils

import (
	"errors"
	"strconv"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func NewGinCors(allowedDomains []string) gin.HandlerFunc {
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = allowedDomains
	corsConfig.AllowCredentials = true
	corsConfig.AllowAllOrigins = false

	return cors.New(corsConfig)
}

// GetUserIDFromContext returns user id from the gin context
func GetUserIDFromContext(ctx *gin.Context) (string, error) {
	claims, ok := ctx.Get("claims")
	if !ok {
		return "", errors.New("claims not found in context")
	}

	claimsMap, ok := claims.(map[string]interface{})
	if !ok {
		return "", errors.New("claims is not a map")
	}

	userId, ok := claimsMap["user"]
	if !ok {
		return "", errors.New("user not found in claims")
	}

	userIdStr, ok := userId.(string)
	if !ok {
		return "", errors.New("user id is not string")
	}

	return userIdStr, nil
}

func ParseInt(s string, defaultVal string) (int, error) {
	if s == "" {
		s = defaultVal
	}

	return strconv.Atoi(s)
}
