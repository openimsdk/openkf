// Copyright © 2023 OpenIM open source community. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v4"

	"github.com/openimsdk/openkf/server/internal/config"
)

type JwtClaims struct {
	UserUUID      string `json:"user_uuid"`
	CommunityUUID string `json:"community_uuid"`
	jwt.RegisteredClaims
}

func GenerateJwtToken(user_uuid string, community_uuid string) (string, uint, error) {
	secret := []byte(config.Config.JWT.Secret)
	issuer := config.Config.JWT.Issuer
	expireDays := config.Config.JWT.ExpireDays

	claims := &JwtClaims{
		user_uuid,
		community_uuid,
		jwt.RegisteredClaims{
			Issuer:    issuer,
			NotBefore: jwt.NewNumericDate(time.Now().Add(-1000)),
			ExpiresAt: jwt.NewNumericDate(time.Now().AddDate(0, 0, expireDays)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token_string, err := token.SignedString(secret)
	expire_time_seconds := uint(time.Now().AddDate(0, 0, expireDays).Unix())

	return token_string, expire_time_seconds, err
}

func ParseJwtToken(tokenString string) (*JwtClaims, error) {
	secret := []byte(config.Config.JWT.Secret)
	token, err := jwt.ParseWithClaims(tokenString, &JwtClaims{}, func(token *jwt.Token) (interface{}, error) {
		return secret, nil
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
