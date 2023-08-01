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

	"github.com/gofrs/uuid"

	"github.com/OpenIMSDK/OpenKF/server/internal/conn/db"
	"github.com/OpenIMSDK/OpenKF/server/internal/dal/cache"
	"github.com/OpenIMSDK/OpenKF/server/internal/dal/gen"
	systemroles "github.com/OpenIMSDK/OpenKF/server/internal/models/system_roles"
)

// SysCustomerDao syscustomer dao.
type SysCustomerDao struct {
	Dao
}

// NewSysCustomerDao return a syscustomer dao.
func NewSysCustomerDao() *SysCustomerDao {
	query := gen.Use(db.GetMysqlDB())
	cache := cache.Use(db.GetRedis())

	return &SysCustomerDao{
		Dao: Dao{
			ctx:   context.Background(),
			query: query,
			cache: &cache,
		},
	}
}

// Create create one or multi models.
func (d *SysCustomerDao) Create(m ...*systemroles.SysCustomer) error {
	return d.query.WithContext(d.ctx).SysCustomer.Create(m...)
}

// First get first matched result.
func (d *SysCustomerDao) First() (*systemroles.SysCustomer, error) {
	return d.query.WithContext(d.ctx).SysCustomer.First()
}

// FindAll get all matched results.
func (d *SysCustomerDao) FindAll() ([]*systemroles.SysCustomer, error) {
	return d.query.WithContext(d.ctx).SysCustomer.Find()
}

// FindFirstById get first matched result by id.
func (d *SysCustomerDao) FindFirstById(id uint) (*systemroles.SysCustomer, error) {
	m := d.query.SysCustomer

	return m.WithContext(d.ctx).Where(m.Id.Eq(id)).First()
}

// FindByIdPage get page by Id.
func (d *SysCustomerDao) FindByIdPage(id uint, offset int, limit int) ([]*systemroles.SysCustomer, int64, error) {
	m := d.query.SysCustomer

	result, count, err := m.WithContext(d.ctx).Where(m.Id.Eq(id)).FindByPage(offset, limit)

	return result, count, err
}

// FindFirstByCreatedAt get first matched result by createdat.
func (d *SysCustomerDao) FindFirstByCreatedAt(createdat time.Time) (*systemroles.SysCustomer, error) {
	m := d.query.SysCustomer

	return m.WithContext(d.ctx).Where(m.CreatedAt.Eq(createdat)).First()
}

// FindByCreatedAtPage get page by CreatedAt.
func (d *SysCustomerDao) FindByCreatedAtPage(createdat time.Time, offset int, limit int) ([]*systemroles.SysCustomer, int64, error) {
	m := d.query.SysCustomer

	result, count, err := m.WithContext(d.ctx).Where(m.CreatedAt.Eq(createdat)).FindByPage(offset, limit)

	return result, count, err
}

// FindFirstByUpdatedAt get first matched result by updatedat.
func (d *SysCustomerDao) FindFirstByUpdatedAt(updatedat time.Time) (*systemroles.SysCustomer, error) {
	m := d.query.SysCustomer

	return m.WithContext(d.ctx).Where(m.UpdatedAt.Eq(updatedat)).First()
}

// FindByUpdatedAtPage get page by UpdatedAt.
func (d *SysCustomerDao) FindByUpdatedAtPage(updatedat time.Time, offset int, limit int) ([]*systemroles.SysCustomer, int64, error) {
	m := d.query.SysCustomer

	result, count, err := m.WithContext(d.ctx).Where(m.UpdatedAt.Eq(updatedat)).FindByPage(offset, limit)

	return result, count, err
}

// FindFirstByDeletedAt get first matched result by deletedat.
func (d *SysCustomerDao) FindFirstByDeletedAt(deletedat time.Time) (*systemroles.SysCustomer, error) {
	m := d.query.SysCustomer

	return m.WithContext(d.ctx).Where(m.DeletedAt.Eq(deletedat)).First()
}

// FindByDeletedAtPage get page by DeletedAt.
func (d *SysCustomerDao) FindByDeletedAtPage(deletedat time.Time, offset int, limit int) ([]*systemroles.SysCustomer, int64, error) {
	m := d.query.SysCustomer

	result, count, err := m.WithContext(d.ctx).Where(m.DeletedAt.Eq(deletedat)).FindByPage(offset, limit)

	return result, count, err
}

// FindFirstByUUID get first matched result by uuid.
func (d *SysCustomerDao) FindFirstByUUID(uuid uuid.UUID) (*systemroles.SysCustomer, error) {
	m := d.query.SysCustomer

	return m.WithContext(d.ctx).Where(m.UUID.Eq(uuid)).First()
}

// FindByUUIDPage get page by UUID.
func (d *SysCustomerDao) FindByUUIDPage(uuid uuid.UUID, offset int, limit int) ([]*systemroles.SysCustomer, int64, error) {
	m := d.query.SysCustomer

	result, count, err := m.WithContext(d.ctx).Where(m.UUID.Eq(uuid)).FindByPage(offset, limit)

	return result, count, err
}

// FindFirstByEmail get first matched result by email.
func (d *SysCustomerDao) FindFirstByEmail(email string) (*systemroles.SysCustomer, error) {
	m := d.query.SysCustomer

	return m.WithContext(d.ctx).Where(m.Email.Eq(email)).First()
}

// FindByEmailPage get page by Email.
func (d *SysCustomerDao) FindByEmailPage(email string, offset int, limit int) ([]*systemroles.SysCustomer, int64, error) {
	m := d.query.SysCustomer

	result, count, err := m.WithContext(d.ctx).Where(m.Email.Eq(email)).FindByPage(offset, limit)

	return result, count, err
}

// FindFirstByNickname get first matched result by nickname.
func (d *SysCustomerDao) FindFirstByNickname(nickname string) (*systemroles.SysCustomer, error) {
	m := d.query.SysCustomer

	return m.WithContext(d.ctx).Where(m.Nickname.Eq(nickname)).First()
}

// FindByNicknamePage get page by Nickname.
func (d *SysCustomerDao) FindByNicknamePage(nickname string, offset int, limit int) ([]*systemroles.SysCustomer, int64, error) {
	m := d.query.SysCustomer

	result, count, err := m.WithContext(d.ctx).Where(m.Nickname.Eq(nickname)).FindByPage(offset, limit)

	return result, count, err
}

// FindFirstByAvatar get first matched result by avatar.
func (d *SysCustomerDao) FindFirstByAvatar(avatar string) (*systemroles.SysCustomer, error) {
	m := d.query.SysCustomer

	return m.WithContext(d.ctx).Where(m.Avatar.Eq(avatar)).First()
}

// FindByAvatarPage get page by Avatar.
func (d *SysCustomerDao) FindByAvatarPage(avatar string, offset int, limit int) ([]*systemroles.SysCustomer, int64, error) {
	m := d.query.SysCustomer

	result, count, err := m.WithContext(d.ctx).Where(m.Avatar.Eq(avatar)).FindByPage(offset, limit)

	return result, count, err
}

// FindFirstByDescription get first matched result by description.
func (d *SysCustomerDao) FindFirstByDescription(description string) (*systemroles.SysCustomer, error) {
	m := d.query.SysCustomer

	return m.WithContext(d.ctx).Where(m.Description.Eq(description)).First()
}

// FindByDescriptionPage get page by Description.
func (d *SysCustomerDao) FindByDescriptionPage(description string, offset int, limit int) ([]*systemroles.SysCustomer, int64, error) {
	m := d.query.SysCustomer

	result, count, err := m.WithContext(d.ctx).Where(m.Description.Eq(description)).FindByPage(offset, limit)

	return result, count, err
}

// FindFirstByDevice get first matched result by device.
func (d *SysCustomerDao) FindFirstByDevice(device string) (*systemroles.SysCustomer, error) {
	m := d.query.SysCustomer

	return m.WithContext(d.ctx).Where(m.Device.Eq(device)).First()
}

// FindByDevicePage get page by Device.
func (d *SysCustomerDao) FindByDevicePage(device string, offset int, limit int) ([]*systemroles.SysCustomer, int64, error) {
	m := d.query.SysCustomer

	result, count, err := m.WithContext(d.ctx).Where(m.Device.Eq(device)).FindByPage(offset, limit)

	return result, count, err
}

// FindFirstByIPAddress get first matched result by ipaddress.
func (d *SysCustomerDao) FindFirstByIPAddress(ipaddress string) (*systemroles.SysCustomer, error) {
	m := d.query.SysCustomer

	return m.WithContext(d.ctx).Where(m.IPAddress.Eq(ipaddress)).First()
}

// FindByIPAddressPage get page by IPAddress.
func (d *SysCustomerDao) FindByIPAddressPage(ipaddress string, offset int, limit int) ([]*systemroles.SysCustomer, int64, error) {
	m := d.query.SysCustomer

	result, count, err := m.WithContext(d.ctx).Where(m.IPAddress.Eq(ipaddress)).FindByPage(offset, limit)

	return result, count, err
}

// FindFirstBySource get first matched result by source.
func (d *SysCustomerDao) FindFirstBySource(source string) (*systemroles.SysCustomer, error) {
	m := d.query.SysCustomer

	return m.WithContext(d.ctx).Where(m.Source.Eq(source)).First()
}

// FindBySourcePage get page by Source.
func (d *SysCustomerDao) FindBySourcePage(source string, offset int, limit int) ([]*systemroles.SysCustomer, int64, error) {
	m := d.query.SysCustomer

	result, count, err := m.WithContext(d.ctx).Where(m.Source.Eq(source)).FindByPage(offset, limit)

	return result, count, err
}

// FindFirstBySourceType get first matched result by sourcetype.
func (d *SysCustomerDao) FindFirstBySourceType(sourcetype int) (*systemroles.SysCustomer, error) {
	m := d.query.SysCustomer

	return m.WithContext(d.ctx).Where(m.SourceType.Eq(sourcetype)).First()
}

// FindBySourceTypePage get page by SourceType.
func (d *SysCustomerDao) FindBySourceTypePage(sourcetype int, offset int, limit int) ([]*systemroles.SysCustomer, int64, error) {
	m := d.query.SysCustomer

	result, count, err := m.WithContext(d.ctx).Where(m.SourceType.Eq(sourcetype)).FindByPage(offset, limit)

	return result, count, err
}

// Update update model.
func (d *SysCustomerDao) Update(m *systemroles.SysCustomer) error {
	res, err := d.query.WithContext(d.ctx).SysCustomer.Updates(m)
	if err != nil && res.Error != nil {
		return err
	}

	return nil
}

// Delete delete model.
func (d *SysCustomerDao) Delete(m ...*systemroles.SysCustomer) error {
	res, err := d.query.WithContext(d.ctx).SysCustomer.Delete(m...)
	if err != nil && res.Error != nil {
		return err
	}

	return nil
}

// Count count matched records.
func (d *SysCustomerDao) Count() (int64, error) {
	return d.query.WithContext(d.ctx).SysCustomer.Count()
}

///////////////////////////////////////////////////////////
//              Append your code here.                   //
///////////////////////////////////////////////////////////
