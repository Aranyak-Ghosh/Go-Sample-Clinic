package jwt

import "github.com/golang-jwt/jwt"

type jwtManager struct {
	privKey []byte
	pubKey  []byte
}

type JWTManager interface {
	GenerateToken(jwt.MapClaims) ([]byte, error)
	VerifyToken([]byte) (jwt.MapClaims, bool, error)
}
