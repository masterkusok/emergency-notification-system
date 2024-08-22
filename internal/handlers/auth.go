package handlers

import (
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
	"math/rand"
	"net/http"
	"time"
)

const saltLen int = 6

type AuthContext struct {
	echo.Context
	IsAuthenticated bool
	Id              uint
}

const charSet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func generateSalt() string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	b := make([]byte, saltLen)
	for i := range b {
		b[i] = charSet[r.Intn(len(charSet))]
	}
	return string(b)
}

func (h *AuthHandler) SignUp(c echo.Context) error {
	request := new(signUpRequest)
	if err := c.Bind(request); err != nil {
		return c.JSON(http.StatusBadRequest, nil)
	}

	salt := generateSalt()
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password+salt), bcrypt.DefaultCost)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, nil)
	}

	_, err = h.provider.CreateUser(request.Name, salt, string(hashedPassword))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, nil)
	}
	return c.JSON(http.StatusOK, nil)
}

func (h *AuthHandler) SignIn(c echo.Context) error {
	request := new(signInRequest)
	if err := c.Bind(request); err != nil {
		return c.JSON(http.StatusBadRequest, nil)
	}

	response := new(signInResponse)
	user, err := h.provider.GetUserByName(request.Name)
	if err != nil {
		return c.JSON(http.StatusBadRequest, nil)
	}

	if user == nil {
		response.Seed(false, "", "invalid username or password")
		return c.JSON(http.StatusOK, response)
	}

	hashed := user.PasswordHash
	err = bcrypt.CompareHashAndPassword([]byte(hashed), []byte(request.Password+user.Salt))
	if err != nil {
		response.Seed(false, "", "invalid username or password")
		return c.JSON(http.StatusOK, response)
	}

	claims := jwt.MapClaims{
		"iss": "issuer",
		"exp": time.Now().Add(7 * 24 * time.Hour).Unix(),
		"id":  user.ID,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenStr, err := token.SignedString([]byte("secret-string"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, nil)
	}
	response.Seed(true, tokenStr, "successful")
	return c.JSON(http.StatusOK, response)
}

func (h *AuthHandler) CurrentUser(c echo.Context) error {
	cont := c.(*AuthContext)
	if cont.IsAuthenticated {
		response := new(userResponse)
		user, err := h.provider.GetUserById(cont.Id)
		if err != nil {
			return err
		}
		response.Seed(user)
		return c.JSON(http.StatusOK, response)
	}
	return c.JSON(http.StatusForbidden, nil)
}
