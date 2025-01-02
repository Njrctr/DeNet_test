package handlers

import (
	"errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

const (
	autorizationHeader = "Authorization"
	userCtx            = "userId"
)

func (h *Handler) userIdentify(c *gin.Context) {
	authHeader := c.GetHeader(autorizationHeader)
	if authHeader == "" {
		newErrorResponse(c, http.StatusUnauthorized, "Empty auth header")
		return
	}

	headerParts := strings.Split(authHeader, " ")
	if len(headerParts) != 2 {
		newErrorResponse(c, http.StatusUnauthorized, "Invalid auth header")
		return
	}
	if headerParts[0] != "Bearer" {
		newErrorResponse(c, http.StatusUnauthorized, "Invalid auth header. Should be Bearer!")
		return
	}
	if headerParts[1] == "" {
		newErrorResponse(c, http.StatusUnauthorized, "Token is empty")
		return
	}

	userId, err := h.services.ParseJWTToken(headerParts[1])
	if err != nil {
		newErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}

	c.Set(userCtx, userId)
}

func getUserId(c *gin.Context) (int, error) {
	id, ok := c.Get(userCtx)
	if !ok {
		return 0, errors.New("user id not found")
	}

	idInt, ok := id.(int)
	if !ok {
		return 0, errors.New("user Id is of invalid type")
	}

	return idInt, nil
}
