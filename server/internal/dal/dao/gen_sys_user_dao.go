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

	"github.com/gofrs/uuid"

	"github.com/OpenIMSDK/OpenKF/server/internal/conn/db"
	"github.com/OpenIMSDK/OpenKF/server/internal/dal/cache"
	"github.com/OpenIMSDK/OpenKF/server/internal/dal/gen"
	systemroles "github.com/OpenIMSDK/OpenKF/server/internal/models/system_roles"
)

// SysUserDao sysuser dao.
type SysUserDao struct {
	Dao
}

// NewSysUserDao return a sysuser dao.
func NewSysUserDao() *SysUserDao {
	query := gen.Use(db.GetMysqlDB())
	cache := cache.Use(db.GetRedis())

	return &SysUserDao{
		Dao: Dao{
			ctx:   context.Background(),
			query: query,
			cache: &cache,
		},
	}
}

// Create create one or multi models.
func (d *SysUserDao) Create(m ...*systemroles.SysUser) error {
	return d.query.WithContext(d.ctx).SysUser.Create(m...)
}

// First get first matched result.
func (d *SysUserDao) First() (*systemroles.SysUser, error) {
	return d.query.WithContext(d.ctx).SysUser.First()
}

// FindAll get all matched results.
func (d *SysUserDao) FindAll() ([]*systemroles.SysUser, error) {
	return d.query.WithContext(d.ctx).SysUser.Find()
}

// FindFirstById get first matched result by id.
func (d *SysUserDao) FindFirstById(id uint) (*systemroles.SysUser, error) {
	m := d.query.SysUser

	return m.WithContext(d.ctx).Where(m.Id.Eq(id)).First()
}

// FindByIdPage get page by Id.
func (d *SysUserDao) FindByIdPage(id uint, offset int, limit int) ([]*systemroles.SysUser, int64, error) {
	m := d.query.SysUser

	result, count, err := m.WithContext(d.ctx).Where(m.Id.Eq(id)).FindByPage(offset, limit)

	return result, count, err
}

// FindFirstByUUID get first matched result by uuid.
func (d *SysUserDao) FindFirstByUUID(uuid uuid.UUID) (*systemroles.SysUser, error) {
	m := d.query.SysUser

	return m.WithContext(d.ctx).Where(m.UUID.Eq(uuid)).First()
}

// FindByUUIDPage get page by UUID.
func (d *SysUserDao) FindByUUIDPage(uuid uuid.UUID, offset int, limit int) ([]*systemroles.SysUser, int64, error) {
	m := d.query.SysUser

	result, count, err := m.WithContext(d.ctx).Where(m.UUID.Eq(uuid)).FindByPage(offset, limit)

	return result, count, err
}

// FindFirstByEmail get first matched result by email.
func (d *SysUserDao) FindFirstByEmail(email string) (*systemroles.SysUser, error) {
	m := d.query.SysUser

	return m.WithContext(d.ctx).Where(m.Email.Eq(email)).First()
}

// FindByEmailPage get page by Email.
func (d *SysUserDao) FindByEmailPage(email string, offset int, limit int) ([]*systemroles.SysUser, int64, error) {
	m := d.query.SysUser

	result, count, err := m.WithContext(d.ctx).Where(m.Email.Eq(email)).FindByPage(offset, limit)

	return result, count, err
}

// FindFirstByNickname get first matched result by nickname.
func (d *SysUserDao) FindFirstByNickname(nickname string) (*systemroles.SysUser, error) {
	m := d.query.SysUser

	return m.WithContext(d.ctx).Where(m.Nickname.Eq(nickname)).First()
}

// FindByNicknamePage get page by Nickname.
func (d *SysUserDao) FindByNicknamePage(nickname string, offset int, limit int) ([]*systemroles.SysUser, int64, error) {
	m := d.query.SysUser

	result, count, err := m.WithContext(d.ctx).Where(m.Nickname.Eq(nickname)).FindByPage(offset, limit)

	return result, count, err
}

// FindFirstByAvatar get first matched result by avatar.
func (d *SysUserDao) FindFirstByAvatar(avatar string) (*systemroles.SysUser, error) {
	m := d.query.SysUser

	return m.WithContext(d.ctx).Where(m.Avatar.Eq(avatar)).First()
}

// FindByAvatarPage get page by Avatar.
func (d *SysUserDao) FindByAvatarPage(avatar string, offset int, limit int) ([]*systemroles.SysUser, int64, error) {
	m := d.query.SysUser

	result, count, err := m.WithContext(d.ctx).Where(m.Avatar.Eq(avatar)).FindByPage(offset, limit)

	return result, count, err
}

// FindFirstByPassword get first matched result by password.
func (d *SysUserDao) FindFirstByPassword(password string) (*systemroles.SysUser, error) {
	m := d.query.SysUser

	return m.WithContext(d.ctx).Where(m.Password.Eq(password)).First()
}

// FindByPasswordPage get page by Password.
func (d *SysUserDao) FindByPasswordPage(password string, offset int, limit int) ([]*systemroles.SysUser, int64, error) {
	m := d.query.SysUser

	result, count, err := m.WithContext(d.ctx).Where(m.Password.Eq(password)).FindByPage(offset, limit)

	return result, count, err
}

// FindFirstByCommunityId get first matched result by communityid.
func (d *SysUserDao) FindFirstByCommunityId(communityid uint) (*systemroles.SysUser, error) {
	m := d.query.SysUser

	return m.WithContext(d.ctx).Where(m.CommunityId.Eq(communityid)).First()
}

// FindByCommunityIdPage get page by CommunityId.
func (d *SysUserDao) FindByCommunityIdPage(communityid uint, offset int, limit int) ([]*systemroles.SysUser, int64, error) {
	m := d.query.SysUser

	result, count, err := m.WithContext(d.ctx).Where(m.CommunityId.Eq(communityid)).FindByPage(offset, limit)

	return result, count, err
}

// Update update model.
func (d *SysUserDao) Update(m *systemroles.SysUser) error {
	res, err := d.query.WithContext(d.ctx).SysUser.Updates(m)
	if err != nil && res.Error != nil {
		return err
	}

	return nil
}

// Delete delete model.
func (d *SysUserDao) Delete(m ...*systemroles.SysUser) error {
	res, err := d.query.WithContext(d.ctx).SysUser.Delete(m...)
	if err != nil && res.Error != nil {
		return err
	}

	return nil
}

// Count count matched records.
func (d *SysUserDao) Count() (int64, error) {
	return d.query.WithContext(d.ctx).SysUser.Count()
}

///////////////////////////////////////////////////////////
//              Append your code here.                   //
///////////////////////////////////////////////////////////
