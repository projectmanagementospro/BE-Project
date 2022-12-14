package exception

import (
	"BE-Project/helper"
	"net/http"

	"github.com/gin-gonic/gin"
)

type TokenError struct {
	context *gin.Context
	log     helper.Log
}

func NewTokenError(context *gin.Context, log helper.Log) TokenError {
	return TokenError{
		context: context,
		log:     log,
	}
}

// this code is used to set the meta of the error
func (tokenError TokenError) SetMeta(message error) bool {
	tokenError.context.Error(message).SetMeta("TOKEN_ERROR")
	tokenError.context.Status(http.StatusUnauthorized)
	return true
}

// this code is used to log the error
func (tokenError TokenError) Logf(message error, args ...interface{}) {
	tokenError.log.Errorf("Token Error : " + message.Error(), args...)
}