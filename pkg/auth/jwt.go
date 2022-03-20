package auth

import (
	"goapp/global"

	"github.com/goer-project/goer/auth"
)

func NewJWT() *auth.JWT {
	return &auth.JWT{
		JwtSecret: []byte(global.Config.JWT.SecretKey),
		JwtTtl:    global.Config.JWT.TTL,
	}
}

func NewJWTGuard(guard string) *auth.JWT {
	return &auth.JWT{
		JwtSecret: []byte(global.Config.JWT.SecretKey),
		JwtTtl:    global.Config.JWT.TTL,
		Guard:     guard,
	}
}
