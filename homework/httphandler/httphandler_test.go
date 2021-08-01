package httphandler

import (
	"encoding/base64"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
	"github.com/stretchr/testify/require"
)

func TestTryHTTPHandler(t *testing.T) {
	app := fiber.New(fiber.Config{
		Immutable:     true,
		CaseSensitive: true,
		AppName:       "Materi Rakamin v1.0.0",
	})
	const secretKey = "dabest!"
	const authHeader = "Authorization"
	TryHTTPHandler(app, secretKey)

	rakaminURL := "https://rakamin.com"
	req := httptest.NewRequest("GET", rakaminURL+"/internal", nil)
	resp, err := app.Test(req)
	require.Nil(t, err)
	require.Equal(t, http.StatusUnauthorized, resp.StatusCode)

	basicAuth := base64.StdEncoding.EncodeToString([]byte("admin:rakamin"))
	req.Header.Set(authHeader, "Basic "+basicAuth)
	resp, err = app.Test(req)
	require.Nil(t, err)
	require.Equal(t, http.StatusOK, resp.StatusCode)

	var sb = new(strings.Builder)
	_, err = io.Copy(sb, resp.Body)
	require.Nil(t, err)
	require.Equal(t, "Welcome to the internal of Rakamin!", sb.String())

	req = httptest.NewRequest("GET", rakaminURL+"/admin", nil)
	resp, err = app.Test(req)
	require.Nil(t, err)
	require.Equal(t, http.StatusUnauthorized, resp.StatusCode)

	signJwt := jwt.New(jwt.SigningMethodHS256)
	claims := signJwt.Claims.(jwt.MapClaims)
	claims["name"] = "Admin Rakamin"
	claims["admin"] = true
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()
	token, err := signJwt.SignedString([]byte(secretKey))
	require.Nil(t, err)

	req.Header.Set(authHeader, "Bearer "+token)
	resp, err = app.Test(req)
	require.Nil(t, err)
	require.Equal(t, http.StatusOK, resp.StatusCode)

	sb.Reset()
	_, err = io.Copy(sb, resp.Body)
	require.Nil(t, err)
	require.Equal(t, "Welcome back, admin!", sb.String())
}
