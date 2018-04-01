package helper

import (
	"crypto/rand"
	"encoding/base64"
	"rpgo-backend/models"

	"github.com/gin-gonic/gin"
)

// SetAdminSession adds cookies from login form
func SetAdminSession(c *gin.Context, admin models.Administrador) {
	token, _ := generateRandomString(32)

	c.SetCookie("token", token, 28800, "", "", false, true)
	c.SetCookie("nome", admin.Nome, 28800, "", "", false, true)
	c.SetCookie("usuario", admin.Usuario, 28800, "", "", false, true)
	c.SetCookie("logged", "true", 28800, "", "", false, true)
}

// generateRandomString returns a random string by range parameter
func generateRandomString(number int) (string, error) {
	b, err := generateRandomBytes(number)
	return base64.URLEncoding.EncodeToString(b), err
}

// generateRandomBytes
func generateRandomBytes(n int) ([]byte, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)

	if err != nil {
		return nil, err
	}

	return b, nil
}

/**
 * LoggedIn
 */
func LoggedIn(c *gin.Context) bool {
	loggedInInterface, _ := c.Cookie("logged")
	logged := false
	if loggedInInterface == "true" {
		logged = true
	}
	return logged
}

// GetActiveAdminData
func GetActiveAdminData(c *gin.Context) models.Administrador {
	admin := models.Administrador{}
	admin.Nome, _ = c.Cookie("nome")
	admin.Usuario, _ = c.Cookie("usuario")

	return admin
}

// decodeCookie
func decodeCookie(name string, c *gin.Context) string {
	cookie, err := c.Cookie(name)
	if err == nil {
		return Decode(cookie)
	}
	return ""
}
