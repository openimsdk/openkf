// Copyright © 2023 OpenIMSDK open source community. All rights reserved.
// Licensed under the MIT License (the "License");
// you may not use this file except in compliance with the License.

package utils

import "strconv"

func IntToString(i interface{}) string {
	return strconv.FormatInt(int64(i.(int)), 10)
}

func StringToInt(i string) int {
	j, _ := strconv.Atoi(i)
	return j
}
