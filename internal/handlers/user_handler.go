package handlers

import (
	"Bookify/internal/models"
	"Bookify/internal/services"

	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

type AuthHandler struct {
	base     BaseHandler
	services services.AuthService
}

func NewAuthHandler(services services.AuthService) *AuthHandler {
	return &AuthHandler{base: *NewBaseHandler(), services: services}
}

// Register
func (h *AuthHandler) Register(c echo.Context) error {
	user := new(models.User)
	if err := c.Bind(&user); err != nil {
		return h.base.JSONBadRequest(c, err)
	}
	//hash password
	hashedpassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	user.Password = string(hashedpassword) // byte => string

	if err := h.services.Register(user); err != nil {
		return h.base.JSONInternalServer(c, err)
	}

	return h.base.JSONCreated(c, "Register Successfully")
}

func (h *AuthHandler) Login(c echo.Context) error {
	var loginReq struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := c.Bind(&loginReq); err != nil {
		return h.base.JSONBadRequest(c, err)
	}

	token, err := h.services.Login(loginReq.Email, loginReq.Password)
	if err != nil {
		return h.base.JSONUnAuth(c, err)
	}

	return h.base.JSONSuccess(c, map[string]string{"accesstoken": token})

}
