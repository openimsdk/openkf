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

type testHook struct {
	Hook
	priority int64
	url      string
}

func (h testHook) Priority() int64 {
	return h.priority
}

func (h testHook) Pattern() string {
	return h.url
}

func TestUrlTrie(t *testing.T) {
	hooks := []Hook{
		testHook{priority: 1, url: "/gin"},
		testHook{priority: 1, url: "/api/v1/123"},
		testHook{priority: 1, url: "/openkf/*"},
		testHook{priority: 2, url: "/gin"},
	}
	trie := NewTrie()
	for _, h := range hooks {
		trie.Insert(h.Pattern(), h)
	}

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
