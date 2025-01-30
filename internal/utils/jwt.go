package utils

import (
	"fmt"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var (
	secretKey     string                 = "trungnguyenissohandsomesoicannotrefusehim"
	issuer        string                 = "TrungDepTrai"
	signingMethod *jwt.SigningMethodHMAC = jwt.SigningMethodHS256
	expireTime    time.Duration          = time.Minute * 30
)

func GenerateJwtToken(email string, id int64) (string, error) {
	claims := jwt.RegisteredClaims{
		IssuedAt:  jwt.NewNumericDate(time.Now().UTC()),
		ExpiresAt: jwt.NewNumericDate(time.Now().UTC().Add(expireTime)),
		NotBefore: jwt.NewNumericDate(time.Now().UTC()),
		Issuer:    issuer,
		Subject:   email,
		ID:        fmt.Sprintf("%d", id),
		Audience:  []string{email},
	}
	token := jwt.NewWithClaims(signingMethod, claims)
	tokenString, err := token.SignedString([]byte(secretKey))

	return tokenString, err
}

func VerifyJwtToken(token string) (int64, error) {
	parsedToken, err := parseToken(token)
	if err != nil {
		return 0, fmt.Errorf("token parsing failed: %w", err)
	}

	if !parsedToken.Valid {
		return 0, fmt.Errorf("token is invalid")
	}

	userId, err := extractUserId(parsedToken)
	if err != nil {
		return 0, fmt.Errorf("failed to extract user ID: %w", err)
	}

	return userId, nil
}

func parseToken(token string) (*jwt.Token, error) {
	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secretKey), nil
	})
	return parsedToken, err
}

func extractUserId(parsedToken *jwt.Token) (int64, error) {
	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	if !ok {
		return 0, fmt.Errorf("claims are not in expected format")
	}

	userIdStr, ok := claims["jti"].(string)
	if !ok {
		return 0, fmt.Errorf("user ID not found in token")
	}

	userId, err := strconv.ParseInt(userIdStr, 10, 64)
	if err != nil {
		return 0, fmt.Errorf("invalid user ID format: %w", err)
	}

	return userId, nil
}
