package auth

import (
	"errors"
	"time"

	"goer/global"

	"github.com/golang-jwt/jwt"
)

var (
	TokenExpired     = errors.New("token is expired")
	TokenNotValidYet = errors.New("token not active yet")
	TokenMalformed   = errors.New("that's not even a token")
	TokenInvalid     = errors.New("couldn't handle this token")
)

type JWT struct {
	jwtSecret []byte
	jwtTtl    int64
}

func NewJWT() *JWT {
	return &JWT{
		[]byte(global.Config.JWT.SecretKey),
		global.Config.JWT.TTL,
	}
}

type CustomClaims struct {
	ID uint64 `json:"id"`
	jwt.StandardClaims
}

type JwtToken struct {
	AccessToken string    `json:"access_token"`
	ExpiresAt   time.Time `json:"expires_at"`
}

func (j *JWT) CreateToken(id uint64) (JwtToken, error) {
	// Create the claims
	claims := CustomClaims{
		id,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Unix() + j.jwtTtl,
		},
	}

	// Create a new token object, specifying signing method and the claims
	// you would like it to contain.
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString(j.jwtSecret)

	jt := JwtToken{
		AccessToken: tokenString,
		ExpiresAt:   time.Unix(claims.ExpiresAt, 0),
	}

	return jt, err
}

func (j JWT) ParseToken(tokenString string) (*CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return j.jwtSecret, nil
	})
	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, TokenMalformed
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				// Token is expired
				return nil, TokenExpired
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, TokenNotValidYet
			} else {
				return nil, TokenInvalid
			}
		}
	}

	if token == nil {
		return nil, TokenInvalid
	}

	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		return claims, nil
	} else {
		return nil, TokenInvalid
	}
}
