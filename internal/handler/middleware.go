package handler

import (
	"errors"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

const (
	authorizationHeader = "Authorization"
	userCtx             = "userId"
)

func (h *Handler) userIdentity(c *gin.Context) {
	header := c.GetHeader(authorizationHeader)
	if header == "" {
		NewErrorResponse(c, http.StatusUnauthorized, "empy auth header")
		return
	}
	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 {
		NewErrorResponse(c, http.StatusUnauthorized, "invalid auth header")
		return
	}
	userId, err := h.service.ParseToken(headerParts[1])
	if err != nil {
		NewErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}
	c.Set(userCtx, userId)
}

func (h *Handler) GetUserID(c *gin.Context) (int, error) {
	userId, exists := c.Get(userCtx)
	if !exists {
		NewErrorResponse(c, http.StatusUnauthorized, "User id does not exists")
		return 0, errors.New("user id does not exists")
	}
	intUserId, err := strconv.Atoi(userId.(string))
	if err != nil {
		NewErrorResponse(c, http.StatusUnauthorized, err.Error())
		return 0, errors.New("user id must be int")
	}
	return intUserId, nil
}
