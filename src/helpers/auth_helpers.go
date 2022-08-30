package helpers

import (
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
	user_domain "github.com/yeisonLucio/shopping-cart/src/components/users/domain"
	"github.com/yeisonLucio/shopping-cart/src/config"
	"golang.org/x/crypto/bcrypt"
)

type JwtPayload struct {
	UserID    uint
	Email     string
	FirstName string
	LastName  string
	jwt.StandardClaims
}

func (h *Helpers) GenerateUserTokens(user user_domain.User) (string, string, error) {

	secretKey := config.App.JwtSecretKey

	claims := JwtPayload{
		UserID:    user.ID,
		Email:     user.Email,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(24)).Unix(),
		},
	}

	refreshClaims := JwtPayload{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(168)).Unix(),
		},
	}

	token, err := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		claims,
	).SignedString([]byte(secretKey))

	if err != nil {
		return "", "", err

	}

	refreshToken, err := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		refreshClaims,
	).SignedString([]byte(secretKey))

	if err != nil {
		return "", "", err
	}

	return token, refreshToken, nil
}

func (h *Helpers) VerifyPassword(password string, hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func (h *Helpers) GeneratePassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return string(hashedPassword), err
	}

	return string(hashedPassword), nil
}

func ValidateToken(token string) (JwtPayload, error) {
	secretKey := config.App.JwtSecretKey

	jwtToken, err := jwt.ParseWithClaims(
		token,
		&JwtPayload{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(secretKey), nil
		},
	)

	if err != nil {
		return JwtPayload{}, err
	}

	claims, ok := jwtToken.Claims.(JwtPayload)
	if !ok {
		return claims, errors.New("error obtains payload of token")
	}

	if claims.ExpiresAt < time.Now().Local().Unix() {
		return claims, errors.New("token is expired")
	}

	return claims, nil
}
