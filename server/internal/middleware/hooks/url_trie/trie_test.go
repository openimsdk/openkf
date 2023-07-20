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
	urls     []string
}

func (h *testHook) Priority() int64 {
	return h.priority
}

func (h *testHook) Patterns() []string {
	return h.urls
}

func TestUrlTrie(t *testing.T) {
	// test case
	testData := []struct {
		priority int64
		urls     []string
	}{
		{1, []string{"/gin"}},
		{1, []string{"/api/v1/123"}},
		{1, []string{"/openkf/*"}},
		{2, []string{"/gin"}},
		{3, []string{"/gin/1"}},
	}

	trie := NewTrie()

	// range test data
	for _, data := range testData {
		trie.InsertBatch(data.urls, &testHook{
			urls:     data.urls,
			priority: data.priority,
		})
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
