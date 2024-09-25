package handler

import (
	"net/http"

	"github.com/fleeper2133/tasks-app/internal/domain"
	"github.com/gin-gonic/gin"
)

// SignIn godoc
// @Summary      signIn
// @Description  login
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        input body domain.SignIn true "sign-in info"
// @Success      200  {object}  pkg.TokenJWT
// @Failure      400  {object}  ErrorResponse
// @Failure      404  {object}  ErrorResponse
// @Failure      500  {object}  ErrorResponse
// @Router       /auth/sign-in [post]
func (h *Handler) SignIn(c *gin.Context) {
	var input domain.SignIn

	if err := c.BindJSON(&input); err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	tokens, err := h.service.Authorization.GenerateTokens(input)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, tokens)

}

// SignUp godoc
// @Summary      signUp
// @Description  registration
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        input body domain.SignUp true "sign-up info"
// @Success      201  {integer}  integer 1
// @Failure      400  {object}  ErrorResponse
// @Failure      404  {object}  ErrorResponse
// @Failure      500  {object}  ErrorResponse
// @Router       /auth/sign-up [post]
func (h *Handler) SignUp(c *gin.Context) {
	var input domain.SignUp

	if err := c.BindJSON(&input); err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.service.Authorization.CreateUser(input)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusCreated, map[string]interface{}{
		"id": id,
	})

}

type Refresh struct {
	RefreshToken string `json:"refresh_token" binding:"required"`
}

// Refresh godoc
// @Summary      Refresh
// @Description  registration
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        refresh body Refresh true "refresh info"
// @Success      200  {object}  pkg.TokenJWT
// @Failure      400  {object}  ErrorResponse
// @Failure      404  {object}  ErrorResponse
// @Failure      500  {object}  ErrorResponse
// @Router       /auth/refresh [post]
func (h *Handler) RefreshToken(c *gin.Context) {
	var input Refresh
	if err := c.BindJSON(&input); err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	jwtTokens, err := h.service.RefreshToken(input.RefreshToken)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, jwtTokens)
}

type EmailSending struct {
	To      []string `json:"to" binding:"required"`
	Message string   `json:"message" binding:"required"`
}

// @Summary      Mail
// @Description  Send Mail
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        email body EmailSending true "email info"
// @Success      200  {object} 	map[string]interface{}
// @Failure      400  {object}  ErrorResponse
// @Failure      404  {object}  ErrorResponse
// @Failure      500  {object}  ErrorResponse
// @Router       /auth/send-mail [post]
func (h *Handler) SendMail(c *gin.Context) {
	var input EmailSending
	if err := c.BindJSON(&input); err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.service.Authorization.SendMail(input.To, input.Message); err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"status": "ok",
	})
}
