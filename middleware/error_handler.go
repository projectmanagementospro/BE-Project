package middleware

import (
	"BE-Project/model/web"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type ValidationError struct {
	Key     string `json:"key,omitempty"`
	Message string `json:"message"`
}

func (e ValidationError) Error(splitedError []string) []ValidationError {
	var errors []ValidationError
	for _, error := range splitedError {
		splittedError := strings.Split(error, "'")
		errors = append(errors, ValidationError{
			Key:     splittedError[3],
			Message: "Error :" + splittedError[4] + splittedError[5] + splittedError[6],
		})
	}
	return errors
}

type ErrorHandlerMiddleware struct {
	Context *gin.Context
}

func NewErrorHandlerMiddleware(ctx *gin.Context) *ErrorHandlerMiddleware {
	return &ErrorHandlerMiddleware{
		Context: ctx,
	}
}

func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		if c.Errors != nil {
			err := c.Errors.Last()
			if err.Meta == "VALIDATION_ERROR" {
				validationErrors(c, err)
				return
			} else if err.Meta == "NOT_FOUND" {
				notFoundError(c, err)
				return
			}
			if err.Meta == "UNAUTHORIZED" {
				authenticationError(c, err)
				return
			}
			internalServerError(c, err)
		}
	}
}

func validationErrors(c *gin.Context, err *gin.Error) {
	splittedError := strings.Split(err.Error(), "\n")
	errorWebResponse := web.NewWebErrorResponse(http.StatusBadRequest, "BAD REQUEST", splittedError)
	c.JSON(http.StatusBadRequest, errorWebResponse)
}

func authenticationError(c *gin.Context, err *gin.Error) {
	errorWebResponse := web.NewWebErrorResponse(http.StatusUnauthorized, "UNAUTHORIZED", err.Error())
	c.JSON(http.StatusUnauthorized, errorWebResponse)
}

func notFoundError(c *gin.Context, err *gin.Error) {
	errorWebResponse := web.NewWebErrorResponse(http.StatusNotFound, "Not Found", err.Error())
	c.JSON(http.StatusNotFound, errorWebResponse)
}

func internalServerError(c *gin.Context, err *gin.Error) {
	errorWebResponse := web.NewWebErrorResponse(http.StatusInternalServerError, "INTERNAL SERVER ERROR", err.Error())
	c.JSON(http.StatusInternalServerError, errorWebResponse)
}
