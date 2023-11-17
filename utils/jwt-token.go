package utils

import (
	"belajar-api-goleng/models"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
)

type UserToken struct {
	Id       *uint32 `json:"id"`
	Username *string `json:"username"`
	Email    *string `json:"email"`
	Role     *string `json:"role"`
	jwt.StandardClaims
}

var secretKey = []byte("kieir9c")

func SignToken(user *models.UsersTable) (string, error) {
	fmt.Println("secretKey =", secretKey)
	claems := &UserToken{
		Id:       user.Id,
		Username: user.Username,
		Email:    user.Email,
		Role:     user.Role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claems)
	signedToken, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

func VerifyToken(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method")
		}
		return secretKey, nil
	})

	if err != nil {
		return nil, err
	}

	return token, nil
}
