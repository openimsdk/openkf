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
	systemroles "github.com/OpenIMSDK/OpenKF/server/internal/models/system_roles"
)

// SysBotDao sysbot dao.
type SysBotDao struct {
	Dao
}

// NewSysBotDao return a sysbot dao.
func NewSysBotDao() *SysBotDao {
	query := gen.Use(db.GetMysqlDB())
	cache := cache.Use(db.GetRedis())

	return &SysBotDao{
		Dao: Dao{
			ctx:   context.Background(),
			query: query,
			cache: &cache,
		},
	}
}

// Create create one or multi models.
func (d *SysBotDao) Create(m ...*systemroles.SysBot) error {
	return d.query.WithContext(d.ctx).SysBot.Create(m...)
}

// First get first matched result.
func (d *SysBotDao) First() (*systemroles.SysBot, error) {
	return d.query.WithContext(d.ctx).SysBot.First()
}

// FindAll get all matched results.
func (d *SysBotDao) FindAll() ([]*systemroles.SysBot, error) {
	return d.query.WithContext(d.ctx).SysBot.Find()
}

// FindFirstById get first matched result by id.
func (d *SysBotDao) FindFirstById(id uint) (*systemroles.SysBot, error) {
	m := d.query.SysBot

	return m.WithContext(d.ctx).Where(m.Id.Eq(id)).First()
}

// FindByIdPage get page by Id.
func (d *SysBotDao) FindByIdPage(id uint, offset int, limit int) ([]*systemroles.SysBot, int64, error) {
	m := d.query.SysBot

	result, count, err := m.WithContext(d.ctx).Where(m.Id.Eq(id)).FindByPage(offset, limit)

	return result, count, err
}

// FindFirstByBotAddr get first matched result by botaddr.
func (d *SysBotDao) FindFirstByBotAddr(botaddr string) (*systemroles.SysBot, error) {
	m := d.query.SysBot

	return m.WithContext(d.ctx).Where(m.BotAddr.Eq(botaddr)).First()
}

// FindByBotAddrPage get page by BotAddr.
func (d *SysBotDao) FindByBotAddrPage(botaddr string, offset int, limit int) ([]*systemroles.SysBot, int64, error) {
	m := d.query.SysBot

	result, count, err := m.WithContext(d.ctx).Where(m.BotAddr.Eq(botaddr)).FindByPage(offset, limit)

	return result, count, err
}

// FindFirstByBotPort get first matched result by botport.
func (d *SysBotDao) FindFirstByBotPort(botport int) (*systemroles.SysBot, error) {
	m := d.query.SysBot

	return m.WithContext(d.ctx).Where(m.BotPort.Eq(botport)).First()
}

// FindByBotPortPage get page by BotPort.
func (d *SysBotDao) FindByBotPortPage(botport int, offset int, limit int) ([]*systemroles.SysBot, int64, error) {
	m := d.query.SysBot

	result, count, err := m.WithContext(d.ctx).Where(m.BotPort.Eq(botport)).FindByPage(offset, limit)

	return result, count, err
}

// FindFirstByBotToken get first matched result by bottoken.
func (d *SysBotDao) FindFirstByBotToken(bottoken string) (*systemroles.SysBot, error) {
	m := d.query.SysBot

	return m.WithContext(d.ctx).Where(m.BotToken.Eq(bottoken)).First()
}

// FindByBotTokenPage get page by BotToken.
func (d *SysBotDao) FindByBotTokenPage(bottoken string, offset int, limit int) ([]*systemroles.SysBot, int64, error) {
	m := d.query.SysBot

	result, count, err := m.WithContext(d.ctx).Where(m.BotToken.Eq(bottoken)).FindByPage(offset, limit)

	return result, count, err
}

// FindFirstByNickname get first matched result by nickname.
func (d *SysBotDao) FindFirstByNickname(nickname string) (*systemroles.SysBot, error) {
	m := d.query.SysBot

	return m.WithContext(d.ctx).Where(m.Nickname.Eq(nickname)).First()
}

// FindByNicknamePage get page by Nickname.
func (d *SysBotDao) FindByNicknamePage(nickname string, offset int, limit int) ([]*systemroles.SysBot, int64, error) {
	m := d.query.SysBot

	result, count, err := m.WithContext(d.ctx).Where(m.Nickname.Eq(nickname)).FindByPage(offset, limit)

	return result, count, err
}

// FindFirstByAvatar get first matched result by avatar.
func (d *SysBotDao) FindFirstByAvatar(avatar string) (*systemroles.SysBot, error) {
	m := d.query.SysBot

	return m.WithContext(d.ctx).Where(m.Avatar.Eq(avatar)).First()
}

// FindByAvatarPage get page by Avatar.
func (d *SysBotDao) FindByAvatarPage(avatar string, offset int, limit int) ([]*systemroles.SysBot, int64, error) {
	m := d.query.SysBot

	result, count, err := m.WithContext(d.ctx).Where(m.Avatar.Eq(avatar)).FindByPage(offset, limit)

	return result, count, err
}

// FindFirstByCommunityId get first matched result by communityid.
func (d *SysBotDao) FindFirstByCommunityId(communityid uint) (*systemroles.SysBot, error) {
	m := d.query.SysBot

	return m.WithContext(d.ctx).Where(m.CommunityId.Eq(communityid)).First()
}

// FindByCommunityIdPage get page by CommunityId.
func (d *SysBotDao) FindByCommunityIdPage(communityid uint, offset int, limit int) ([]*systemroles.SysBot, int64, error) {
	m := d.query.SysBot

	result, count, err := m.WithContext(d.ctx).Where(m.CommunityId.Eq(communityid)).FindByPage(offset, limit)

	return result, count, err
}

// Update update model.
func (d *SysBotDao) Update(m *systemroles.SysBot) error {
	res, err := d.query.WithContext(d.ctx).SysBot.Updates(m)
	if err != nil && res.Error != nil {
		return err
	}

	return nil
}

// Delete delete model.
func (d *SysBotDao) Delete(m ...*systemroles.SysBot) error {
	res, err := d.query.WithContext(d.ctx).SysBot.Delete(m...)
	if err != nil && res.Error != nil {
		return err
	}

	return nil
}

// Count count matched records.
func (d *SysBotDao) Count() (int64, error) {
	return d.query.WithContext(d.ctx).SysBot.Count()
}

///////////////////////////////////////////////////////////
//              Append your code here.                   //
///////////////////////////////////////////////////////////
