package jwt

import "github.com/golang-jwt/jwt"

type jwtManager struct {
}

type JWTManager interface {
	GenerateToken(claims jwt.MapClaims) ([]byte, error)
	VerifyToken(token []byte) (jwt.MapClaims, bool, error)
}
