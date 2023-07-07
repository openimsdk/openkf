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
	"reflect"
	"strings"
	"unicode"
)

// DaoGenerator is a generator for dao.
type DaoGenerator struct {
	buf      *bytes.Buffer
	config   *config
	savePath string
}

// field struct.
type field struct {
	FieldName     string
	NormFieldName string
	FieldType     string
}

// config struct.
type config struct {
	PkgName       string
	NormPkgName   string
	ModelName     string
	NormModelName string
	Fields        []field
}

// getFields get fields from struct.
func getFields(m interface{}) []field {
	return getFieldsRecursively(reflect.TypeOf(m))
}

// getFieldsRecursively get fields recursively.
func getFieldsRecursively(value reflect.Type) []field {
	var fields []field

	for i := 0; i < value.NumField(); i++ {
		v := value.Field(i)

		// Get the tag value
		if v.Tag != "" && v.Type.Kind() != reflect.Struct &&
			v.Type.Kind() != reflect.Bool {
			f := field{
				FieldName:     v.Name,
				NormFieldName: strings.ToLower(v.Name),
				FieldType:     v.Type.String(),
			}
			fields = append(fields, f)

			continue
		}

		// Get the embedded struct fields if the tag is empty
		if v.Tag == "" && v.Type.Kind() == reflect.Struct {
			fields = append(fields, getFieldsRecursively(v.Type)...)
		}
	}

	return fields
}

// NewDaoGenerator returns a new DaoGenerator.
func NewDaoGenerator(m interface{}, savePath string) *DaoGenerator {
	pkgPath := reflect.TypeOf(m).PkgPath()
	paths := strings.Split(pkgPath, "/")
	pkgName := paths[len(paths)-1]

	return &DaoGenerator{
		buf: bytes.NewBuffer(nil),
		config: &config{
			PkgName:       pkgName,
			NormPkgName:   strings.ReplaceAll(pkgName, "_", ""),
			ModelName:     reflect.TypeOf(m).Name(),
			NormModelName: strings.ToLower(reflect.TypeOf(m).Name()),
			Fields:        getFields(m),
		},
		savePath: savePath,
	}
}

// Generate init the hook.
func (g *DaoGenerator) Generate() *DaoGenerator {
	if err := hookTemplate.Execute(g.buf, g.config); err != nil {
		panic(err)
	}

	return g
}

// Format format the generated code.
func (g *DaoGenerator) Format() *DaoGenerator {
	formatOut, err := format.Source(g.buf.Bytes())
	if err != nil {
		panic(err)
	}
	g.buf = bytes.NewBuffer(formatOut)

	return g
}

// Flush write the generated code to file.
func (g *DaoGenerator) Flush() {
	filename := fmt.Sprintf("gen_%s_dao.go", camelToUnderscore(g.config.ModelName))
	if err := ioutil.WriteFile(
		fmt.Sprintf("%s/%s", g.savePath, filename),
		g.buf.Bytes(),
		fs.ModePerm); err != nil {
		panic(err)
	}
	fmt.Println("[OpenKF] gen file ok: ", fmt.Sprintf("%s/%s", g.savePath, filename))
}

// camelToUnderscore convert camel to underscore.
func camelToUnderscore(s string) string {
	var result []rune

	for i, r := range s {
		if unicode.IsUpper(r) {
			if i > 0 && (unicode.IsLower(rune(s[i-1])) || (i+1 < len(s) && unicode.IsLower(rune(s[i+1])))) {
				result = append(result, '_')
			}
			result = append(result, unicode.ToLower(r))
		} else {
			result = append(result, r)
		}
	}

	return strings.ToLower(string(result))
}
