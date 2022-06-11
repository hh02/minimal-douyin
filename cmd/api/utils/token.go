package utils

import (
	"github.com/golang-jwt/jwt/v4"
	"github.com/hh02/minimal-douyin/cmd/api/middleware"
	"github.com/hh02/minimal-douyin/pkg/constants"
)

func GenerateTokenString(userId int64) (string, error) {
	// Create the token
	token := jwt.New(jwt.GetSigningMethod(middleware.AuthMiddleware.SigningAlgorithm))
	claims := token.Claims.(jwt.MapClaims)
	claims[constants.IdentityKey] = userId
	expire := middleware.AuthMiddleware.TimeFunc().Add(middleware.AuthMiddleware.Timeout)
	claims["exp"] = expire.Unix()
	claims["orig_iat"] = middleware.AuthMiddleware.TimeFunc().Unix()
	tokenString, err := token.SignedString(middleware.AuthMiddleware.Key)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
