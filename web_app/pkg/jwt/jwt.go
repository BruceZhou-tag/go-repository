package jwt

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

var salt = []byte(viper.GetString("auth.jwt_salt"))

type MyClaims struct {
	UserID   int64  `json:"user_id"`
	Username string `json:"username"`
	jwt.RegisteredClaims
}

func GenToken(userID int64, Username string) (string, error) {
	c := MyClaims{
		userID,
		Username,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(viper.GetInt64("auth.expire_hour")) * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    viper.GetString("auth.issuer"),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	return token.SignedString(salt)
}

func ParseToken(tokenString string) (*MyClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &MyClaims{}, func(token *jwt.Token) (any, error) {
		return salt, nil
	})
	if err != nil {
		zap.L().Error("ParseToken error", zap.Error(err))
		return nil, err
	} else if claims, ok := token.Claims.(*MyClaims); ok {
		return claims, nil
	} else {
		zap.L().Error("ParseToken with unknow claims type,cannot processed", zap.Error(err))
		return nil, errors.New("invalid token")
	}
}
