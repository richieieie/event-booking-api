package utils

import (
	"fmt"
	"log"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/richieieie/event-booking/internal/model"
)

var (
	secretKey string = "trungnguyenissohandsomesoicannotrefusehim"
	issuer    string = "TrungDepTrai"
)

func GenerateJwtToken(user model.User) (string, error) {
	log.Println(secretKey, issuer)
	claims := jwt.RegisteredClaims{
		IssuedAt:  jwt.NewNumericDate(time.Now().UTC()),
		ExpiresAt: jwt.NewNumericDate(time.Now().UTC().Add(time.Minute * 3)),
		NotBefore: jwt.NewNumericDate(time.Now().UTC()),
		Issuer:    issuer,
		Subject:   user.Email,
		ID:        fmt.Sprintf("%d", user.Id),
		Audience:  []string{user.Email},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(secretKey))

	return tokenString, err
}
