package jwt

import "github.com/golang-jwt/jwt"

type jwtManager struct {
}

type JWTManager interface {
	GenerateToken(jwt.MapClaims) ([]byte, error)
	VerifyToken([]byte) (bool, error)
}
