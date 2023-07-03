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

package urltrie

import (
	"fmt"
	"reflect"
	"testing"
)

type Hook1 struct {
	Hook
}

func (h Hook1) Priority() int64 {
	return 1
}

type Hook2 struct {
	Hook
}

func (h Hook2) Priority() int64 {
	return 1
}

type Hook3 struct {
	Hook
}

func (h Hook3) Priority() int64 {
	return 1
}

type Hook4 struct {
	Hook
}

func (h Hook4) Priority() int64 {
	return 2
}

func TestUrlTrie(t *testing.T) {
	trie := NewTrie()
	trie.Insert("/gin", Hook1{})
	trie.Insert("/api/v1/123", Hook2{})
	trie.Insert("/openkf/*", Hook3{})
	trie.Insert("/gin", Hook4{})

	values, matched := trie.Match("/gin/1")
	if matched {
		fmt.Printf("Matched, values: %#v\n", reflect.ValueOf(values)) // Output: Matched, values: [1 2]
	} else {
		fmt.Println("No match found")
	}

	values, matched = trie.Match("/openim")
	if matched {
		fmt.Printf("Matched, values: %#v\n", reflect.ValueOf(values)) // Output: Matched, values: [1 2]
	} else {
		fmt.Println("No match found")
	}

	values, matched = trie.Match("/openkf/v1")
	if matched {
		fmt.Printf("Matched, values: %#v\n", reflect.ValueOf(values)) // Output: Matched, values: [1 2]
	} else {
		fmt.Println("No match found")
	}
}
