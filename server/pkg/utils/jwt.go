// Copyright © 2023 OpenIMSDK open source community. All rights reserved.
// Licensed under the MIT License (the "License");
// you may not use this file except in compliance with the License.

package utils

import (
	"time"

	"github.com/OpenIMSDK/OpenKF/server/internal/config"
	"github.com/golang-jwt/jwt/v4"
)

type JwtClaims struct {
	jwt.StandardClaims
}

var _secret []byte
var _issuer string
var _expireDays int

func init() {
	_secret = []byte(config.GetString("jwt.secret"))
	_issuer = config.GetString("jwt.issuer")
	_expireDays = config.GetInt("jwt.expire_days")
}

func GenerateJwtToken(claims *JwtClaims) (string, error) {
	claims.Issuer = _issuer
	claims.NotBefore = int64(time.Now().Unix())
	claims.ExpiresAt = int64(time.Now().AddDate(0, 0, _expireDays).Unix())

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(_secret)
}

func ParseJwtToken(tokenString string) (*JwtClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &JwtClaims{}, func(token *jwt.Token) (interface{}, error) {
		return _secret, nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*JwtClaims); ok && token.Valid {
		return claims, nil
	} else {
		return nil, err
	}
}
