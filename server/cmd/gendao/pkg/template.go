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

import "html/template"

func checkTemplate(t string) *template.Template {
	return template.Must(template.New("").Parse(t))
}

var hookTemplate = checkTemplate(`
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

package dao

import (
	"context"

	"github.com/OpenIMSDK/OpenKF/server/internal/conn/db"
	"github.com/OpenIMSDK/OpenKF/server/internal/dal/cache"
	"github.com/OpenIMSDK/OpenKF/server/internal/dal/gen"
	{{.NormPkgName}} "github.com/OpenIMSDK/OpenKF/server/internal/models/{{.PkgName}}"
)

// {{.ModelName}}Dao {{.NormModelName}} dao.
type {{.ModelName}}Dao struct {
	Dao
}

// New{{.ModelName}}Dao return a {{.NormModelName}} dao.
func New{{.ModelName}}Dao() *{{.ModelName}}Dao {
	query := gen.Use(db.GetMysqlDB())
	cache := cache.Use(db.GetRedis())

	return &{{.ModelName}}Dao{
		Dao: Dao{
			ctx:   context.Background(),
			query: query,
			cache: &cache,
		},
	}
}

// Create create one or multi models.
func (d *{{.ModelName}}Dao) Create(m ...*{{.NormPkgName}}.{{.ModelName}}) error {
	return d.query.WithContext(d.ctx).{{.ModelName}}.Create(m...)
}

// First get first matched result.
func (d *{{.ModelName}}Dao) First() (*{{.NormPkgName}}.{{.ModelName}}, error) {
	return d.query.WithContext(d.ctx).{{.ModelName}}.First()
}

// FindAll get all matched results.
func (d *{{.ModelName}}Dao) FindAll() ([]*{{.NormPkgName}}.{{.ModelName}}, error) {
	return d.query.WithContext(d.ctx).{{.ModelName}}.Find()
}
{{$modelName := .ModelName}}{{$normPkgName := .NormPkgName}}{{range .Fields}}
// FindFirstBy{{.FieldName}} get first matched result by {{.NormFieldName}}.
func (d *{{$modelName}}Dao) FindFirstBy{{.FieldName}}({{.NormFieldName}} {{.FieldType}}) (*{{$normPkgName}}.{{$modelName}}, error) {
	m := d.query.{{$modelName}}

	return m.WithContext(d.ctx).Where(m.{{.FieldName}}.Eq({{.NormFieldName}})).First()
}

// FindBy{{.FieldName}}Page get page by {{.FieldName}}.
func (d *{{$modelName}}Dao) FindBy{{.FieldName}}Page({{.NormFieldName}} {{.FieldType}}, offset int, limit int) ([]*{{$normPkgName}}.{{$modelName}}, int64, error) {
	m := d.query.{{$modelName}}

	result, count, err := m.WithContext(d.ctx).Where(m.{{.FieldName}}.Eq({{.NormFieldName}})).FindByPage(offset, limit)

	return result, count, err
}
{{end}}
// Update update model.
func (d *{{.ModelName}}Dao) Update(m *{{.NormPkgName}}.{{.ModelName}}) error {
	res, err := d.query.WithContext(d.ctx).{{.ModelName}}.Updates(m)
	if err != nil && res.Error != nil {
		return err
	}

	return nil
}

// Delete delete model.
func (d *{{.ModelName}}Dao) Delete(m ...*{{.NormPkgName}}.{{.ModelName}}) error {
	res, err := d.query.WithContext(d.ctx).{{.ModelName}}.Delete(m...)
	if err != nil && res.Error != nil {
		return err
	}

	return nil
}

// Count count matched records.
func (d *{{.ModelName}}Dao) Count() (int64, error) {
	return d.query.WithContext(d.ctx).{{.ModelName}}.Count()
}

///////////////////////////////////////////////////////////
//              Append your code here.                   //
///////////////////////////////////////////////////////////
`)
