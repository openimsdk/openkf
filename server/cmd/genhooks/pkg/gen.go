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

package pkg

import (
	"bytes"
	"fmt"
	"go/format"
	"io/fs"
	"io/ioutil"
	"strings"
)

// HookGenerator is a generator for hooks.
type HookGenerator struct {
	buf      *bytes.Buffer
	config   *config
	savePath string
}

type config struct {
	HookName   string
	UrlPattern string
	Priority   int64
}

// NewHookGenerator returns a new HookGenerator.
func NewHookGenerator(hookName, urlPattern, savePath string, priority int64) *HookGenerator {
	return &HookGenerator{
		buf: bytes.NewBuffer(nil),
		config: &config{
			HookName:   hookName,
			UrlPattern: urlPattern,
			Priority:   priority,
		},
		savePath: savePath,
	}
}

// Generate init the hook.
func (g *HookGenerator) Generate() *HookGenerator {
	if err := hookTemplate.Execute(g.buf, g.config); err != nil {
		panic(err)
	}

	return g
}

// Format format the generated code.
func (g *HookGenerator) Format() *HookGenerator {
	formatOut, err := format.Source(g.buf.Bytes())
	if err != nil {
		panic(err)
	}
	g.buf = bytes.NewBuffer(formatOut)

	return g
}

// Flush write the generated code to file.
func (g *HookGenerator) Flush() {
	filename := fmt.Sprintf("gen_%s_hook.go", strings.ToLower(g.config.HookName))
	if err := ioutil.WriteFile(
		fmt.Sprintf("%s/%s", g.savePath, filename),
		g.buf.Bytes(),
		fs.ModePerm); err != nil {
		panic(err)
	}
	fmt.Println("[OpenKF] gen file ok: ", fmt.Sprintf("%s/%s", g.savePath, filename))
}
