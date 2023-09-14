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

	"github.com/openimsdk/openkf/server/internal/dal/cache"
	"github.com/openimsdk/openkf/server/internal/dal/gen"
)

// Dao dao top level.
type Dao struct {
	ctx   context.Context
	query *gen.Query
	cache *cache.Cache
}

// GetQuery get query.
func (d *Dao) GetQuery() *gen.Query {
	return d.query
}

// GetCtx get context.
func (d *Dao) GetCtx() context.Context {
	return d.ctx
}

// GetCache get cache.
func (d *Dao) GetCache() *cache.Cache {
	return d.cache
}
