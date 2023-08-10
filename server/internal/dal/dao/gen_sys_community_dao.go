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
	"time"

	"github.com/OpenIMSDK/OpenKF/server/internal/conn/db"
	"github.com/OpenIMSDK/OpenKF/server/internal/dal/cache"
	"github.com/OpenIMSDK/OpenKF/server/internal/dal/gen"
	systemroles "github.com/OpenIMSDK/OpenKF/server/internal/models/system_roles"
)

// SysCommunityDao syscommunity dao.
type SysCommunityDao struct {
	Dao
}

// NewSysCommunityDao return a syscommunity dao.
func NewSysCommunityDao() *SysCommunityDao {
	query := gen.Use(db.GetMysqlDB())
	cache := cache.Use(db.GetRedis())

	return &SysCommunityDao{
		Dao: Dao{
			ctx:   context.Background(),
			query: query,
			cache: &cache,
		},
	}
}

// Create create one or multi models.
func (d *SysCommunityDao) Create(m ...*systemroles.SysCommunity) error {
	return d.query.WithContext(d.ctx).SysCommunity.Create(m...)
}

// First get first matched result.
func (d *SysCommunityDao) First() (*systemroles.SysCommunity, error) {
	return d.query.WithContext(d.ctx).SysCommunity.First()
}

// FindAll get all matched results.
func (d *SysCommunityDao) FindAll() ([]*systemroles.SysCommunity, error) {
	return d.query.WithContext(d.ctx).SysCommunity.Find()
}

// FindFirstById get first matched result by id.
func (d *SysCommunityDao) FindFirstById(id uint) (*systemroles.SysCommunity, error) {
	m := d.query.SysCommunity

	return m.WithContext(d.ctx).Where(m.Id.Eq(id)).First()
}

// FindByIdPage get page by Id.
func (d *SysCommunityDao) FindByIdPage(id uint, offset int, limit int) ([]*systemroles.SysCommunity, int64, error) {
	m := d.query.SysCommunity

	result, count, err := m.WithContext(d.ctx).Where(m.Id.Eq(id)).FindByPage(offset, limit)

	return result, count, err
}

// FindFirstByCreatedAt get first matched result by createdat.
func (d *SysCommunityDao) FindFirstByCreatedAt(createdat time.Time) (*systemroles.SysCommunity, error) {
	m := d.query.SysCommunity

	return m.WithContext(d.ctx).Where(m.CreatedAt.Eq(createdat)).First()
}

// FindByCreatedAtPage get page by CreatedAt.
func (d *SysCommunityDao) FindByCreatedAtPage(createdat time.Time, offset int, limit int) ([]*systemroles.SysCommunity, int64, error) {
	m := d.query.SysCommunity

	result, count, err := m.WithContext(d.ctx).Where(m.CreatedAt.Eq(createdat)).FindByPage(offset, limit)

	return result, count, err
}

// FindFirstByUpdatedAt get first matched result by updatedat.
func (d *SysCommunityDao) FindFirstByUpdatedAt(updatedat time.Time) (*systemroles.SysCommunity, error) {
	m := d.query.SysCommunity

	return m.WithContext(d.ctx).Where(m.UpdatedAt.Eq(updatedat)).First()
}

// FindByUpdatedAtPage get page by UpdatedAt.
func (d *SysCommunityDao) FindByUpdatedAtPage(updatedat time.Time, offset int, limit int) ([]*systemroles.SysCommunity, int64, error) {
	m := d.query.SysCommunity

	result, count, err := m.WithContext(d.ctx).Where(m.UpdatedAt.Eq(updatedat)).FindByPage(offset, limit)

	return result, count, err
}

// FindFirstByDeletedAt get first matched result by deletedat.
func (d *SysCommunityDao) FindFirstByDeletedAt(deletedat time.Time) (*systemroles.SysCommunity, error) {
	m := d.query.SysCommunity

	return m.WithContext(d.ctx).Where(m.DeletedAt.Eq(deletedat)).First()
}

// FindByDeletedAtPage get page by DeletedAt.
func (d *SysCommunityDao) FindByDeletedAtPage(deletedat time.Time, offset int, limit int) ([]*systemroles.SysCommunity, int64, error) {
	m := d.query.SysCommunity

	result, count, err := m.WithContext(d.ctx).Where(m.DeletedAt.Eq(deletedat)).FindByPage(offset, limit)

	return result, count, err
}

// FindFirstByUUID get first matched result by uuid.
func (d *SysCommunityDao) FindFirstByUUID(uuid string) (*systemroles.SysCommunity, error) {
	m := d.query.SysCommunity

	return m.WithContext(d.ctx).Where(m.UUID.Eq(uuid)).First()
}

// FindByUUIDPage get page by UUID.
func (d *SysCommunityDao) FindByUUIDPage(uuid string, offset int, limit int) ([]*systemroles.SysCommunity, int64, error) {
	m := d.query.SysCommunity

	result, count, err := m.WithContext(d.ctx).Where(m.UUID.Eq(uuid)).FindByPage(offset, limit)

	return result, count, err
}

// FindFirstByName get first matched result by name.
func (d *SysCommunityDao) FindFirstByName(name string) (*systemroles.SysCommunity, error) {
	m := d.query.SysCommunity

	return m.WithContext(d.ctx).Where(m.Name.Eq(name)).First()
}

// FindByNamePage get page by Name.
func (d *SysCommunityDao) FindByNamePage(name string, offset int, limit int) ([]*systemroles.SysCommunity, int64, error) {
	m := d.query.SysCommunity

	result, count, err := m.WithContext(d.ctx).Where(m.Name.Eq(name)).FindByPage(offset, limit)

	return result, count, err
}

// FindFirstByEmail get first matched result by email.
func (d *SysCommunityDao) FindFirstByEmail(email string) (*systemroles.SysCommunity, error) {
	m := d.query.SysCommunity

	return m.WithContext(d.ctx).Where(m.Email.Eq(email)).First()
}

// FindByEmailPage get page by Email.
func (d *SysCommunityDao) FindByEmailPage(email string, offset int, limit int) ([]*systemroles.SysCommunity, int64, error) {
	m := d.query.SysCommunity

	result, count, err := m.WithContext(d.ctx).Where(m.Email.Eq(email)).FindByPage(offset, limit)

	return result, count, err
}

// FindFirstByAvatar get first matched result by avatar.
func (d *SysCommunityDao) FindFirstByAvatar(avatar string) (*systemroles.SysCommunity, error) {
	m := d.query.SysCommunity

	return m.WithContext(d.ctx).Where(m.Avatar.Eq(avatar)).First()
}

// FindByAvatarPage get page by Avatar.
func (d *SysCommunityDao) FindByAvatarPage(avatar string, offset int, limit int) ([]*systemroles.SysCommunity, int64, error) {
	m := d.query.SysCommunity

	result, count, err := m.WithContext(d.ctx).Where(m.Avatar.Eq(avatar)).FindByPage(offset, limit)

	return result, count, err
}

// FindFirstByDescription get first matched result by description.
func (d *SysCommunityDao) FindFirstByDescription(description string) (*systemroles.SysCommunity, error) {
	m := d.query.SysCommunity

	return m.WithContext(d.ctx).Where(m.Description.Eq(description)).First()
}

// FindByDescriptionPage get page by Description.
func (d *SysCommunityDao) FindByDescriptionPage(description string, offset int, limit int) ([]*systemroles.SysCommunity, int64, error) {
	m := d.query.SysCommunity

	result, count, err := m.WithContext(d.ctx).Where(m.Description.Eq(description)).FindByPage(offset, limit)

	return result, count, err
}

// Update update model.
func (d *SysCommunityDao) Update(m *systemroles.SysCommunity) error {
	res, err := d.query.WithContext(d.ctx).SysCommunity.Updates(m)
	if err != nil && res.Error != nil {
		return err
	}

	return nil
}

// Delete delete model.
func (d *SysCommunityDao) Delete(m ...*systemroles.SysCommunity) error {
	res, err := d.query.WithContext(d.ctx).SysCommunity.Delete(m...)
	if err != nil && res.Error != nil {
		return err
	}

	return nil
}

// Count count matched records.
func (d *SysCommunityDao) Count() (int64, error) {
	return d.query.WithContext(d.ctx).SysCommunity.Count()
}

///////////////////////////////////////////////////////////
//              Append your code here.                   //
///////////////////////////////////////////////////////////
