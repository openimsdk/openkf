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

import "regexp"

// IsValidEmail check if the email is valid.
func IsValidEmail(email string) bool {
	// Regular expression for basic email validation
	// This regex pattern is a simplified version and may not cover all edge cases.
	// You can use more comprehensive patterns depending on your specific needs.
	emailRegex := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`

	return regexp.MustCompile(emailRegex).MatchString(email)
}
