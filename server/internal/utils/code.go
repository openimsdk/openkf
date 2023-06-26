// Copyright © 2023 OpenIMSDK open source community. All rights reserved.
// Licensed under the MIT License (the "License");
// you may not use this file except in compliance with the License.

package utils

import (
	"math/rand"
	"time"
)

// Generate a random code with length 6
func GenerateCode() string {
	dataset := "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz!@#$%^&*()"
	rand.Seed(time.Now().UnixNano())

	code := make([]byte, 6)
	for i := 0; i < 6; i++ {
		code[i] = dataset[rand.Intn(len(dataset))]
	}

	return string(code)
}
