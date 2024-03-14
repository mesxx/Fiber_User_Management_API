package helpers

import (
	"os"
	"time"

	"github.com/mesxx/Fiber_User_Management_API/models"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateJWT(user *models.User) (string, error) {
	secretKey := os.Getenv("SECRET_KEY")

	// START create claims
	claims := models.CustomClaims{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(30 * time.Minute)),
		},
	}
	// END create claims

	// START create token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// END create token

	// START generate encoded token
	tokenEncoded, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}
	// END generate encoded token

	return tokenEncoded, nil
}
