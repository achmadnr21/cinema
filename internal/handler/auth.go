package handler

import (
	"fmt"
	"net/http"

	"github.com/achmadnr21/cinema/internal/domain/dto"
	"github.com/achmadnr21/cinema/internal/usecase"
	"github.com/achmadnr21/cinema/internal/utils"
	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	uc *usecase.AuthUsecase
}

func NewAuthHandler(apiV *gin.RouterGroup, uc *usecase.AuthUsecase) {

	AuthHandler := &AuthHandler{
		uc: uc,
	}

	auth := apiV.Group("/auth")
	{
		auth.POST("/login", AuthHandler.Login)
		auth.POST("/refresh", AuthHandler.RefreshToken)
		auth.POST("/register", AuthHandler.Register)
	}
}

func (h *AuthHandler) Register(c *gin.Context) {
	var req dto.UserCreateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, utils.ResponseError("Invalid request, field is missing or invalid"))
		return
	}
	fmt.Println("Register request:", req)
	user, err := h.uc.Register(&req)
	if err != nil {
		c.JSON(utils.GetHTTPErrorCode(err), utils.ResponseError(err.Error()))
		return
	}
	c.JSON(http.StatusCreated, utils.ResponseSuccess("User registered successfully", user))
}

func (h *AuthHandler) Login(c *gin.Context) {
	var req dto.UserLoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request"})
		return
	}
	fmt.Println("Login request:", req)
	token, refreshToken, err := h.uc.Login(req.Email, req.Password)
	if err != nil {
		c.JSON(utils.GetHTTPErrorCode(err), utils.ResponseError(err.Error()))
		return
	}
	c.SetCookie(
		"refresh_token", // name
		refreshToken,    // value
		7*24*60*60,      // maxAge (7 hari dalam detik)
		"/",             // path
		"",              // domain (kosong = current domain)
		false,           // secure (true = hanya lewat HTTPS)
		true,            // httpOnly (tidak bisa diakses via JS)
	)

	var loginResponse = struct {
		Token string `json:"access_token"`
	}{
		Token: token,
	}
	c.JSON(http.StatusOK, utils.ResponseSuccess("Login successful", loginResponse))
}

func (h *AuthHandler) RefreshToken(c *gin.Context) {
	// Ambil refresh token dari cookie
	refreshToken, err := c.Cookie("refresh_token")
	if err != nil || refreshToken == "" {
		c.JSON(http.StatusUnauthorized, utils.ResponseError("Refresh token missing or invalid"))
		return
	}

	// Proses generate token baru
	newAccessToken, newRefreshToken, err := h.uc.RefreshToken(refreshToken)
	if err != nil {
		c.JSON(utils.GetHTTPErrorCode(err), utils.ResponseError(err.Error()))
		return
	}

	// Set ulang cookie refresh token yang baru
	c.SetCookie(
		"refresh_token",
		newRefreshToken,
		7*24*60*60, // valid 7 hari
		"/",
		"",
		true, // secure: true jika HTTPS
		true, // httpOnly
	)

	// Response: hanya access_token di JSON
	c.JSON(http.StatusOK, utils.ResponseSuccess("Token refreshed", gin.H{
		"access_token": newAccessToken,
	}))
}
