package auth

import (
	"errors"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// JWTSecretKey chave secreta para assinar os tokens
var JWTSecretKey = []byte(os.Getenv("JWT_SECRET"))

// Claims define os dados contidos no token JWT
type Claims struct {
	Username string `json:"username"`
	Role     string `json:"role"`
	jwt.StandardClaims
}

// GenerateJWT gera um novo token JWT
func GenerateJWT(username, role string) (string, error) {
	// Tempo de expiração do token (1 hora)
	expirationTime := time.Now().Add(1 * time.Hour).Unix()

	// Criar as Claims
	claims := &Claims{
		Username: username,
		Role:     role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime,
		},
	}

	// Criar token JWT
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(JWTSecretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// ValidateJWT valida um token JWT
func ValidateJWT(tokenString string) (*Claims, error) {
	claims := &Claims{}

	// Verificar e decodificar o token
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return JWTSecretKey, nil
	})

	if err != nil || !token.Valid {
		return nil, errors.New("token inválido ou expirado")
	}

	return claims, nil
}
