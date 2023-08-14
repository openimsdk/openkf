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
	customerroles "github.com/OpenIMSDK/OpenKF/server/internal/models/customer_roles"
)

// SLACK_PERFIX do not use separator.
const SLACK_PERFIX = "SLACK:"

// CustomerSlackDao customerslack dao.
type CustomerSlackDao struct {
	Dao
}

// NewCustomerSlackDao return a customerslack dao.
func NewCustomerSlackDao() *CustomerSlackDao {
	query := gen.Use(db.GetMysqlDB())
	cache := cache.Use(db.GetRedis())

	return &CustomerSlackDao{
		Dao: Dao{
			ctx:   context.Background(),
			query: query,
			cache: &cache,
		},
	}
}

// Create create one or multi models.
func (d *CustomerSlackDao) Create(m ...*customerroles.CustomerSlack) error {
	return d.query.WithContext(d.ctx).CustomerSlack.Create(m...)
}

// First get first matched result.
func (d *CustomerSlackDao) First() (*customerroles.CustomerSlack, error) {
	return d.query.WithContext(d.ctx).CustomerSlack.First()
}

// FindAll get all matched results.
func (d *CustomerSlackDao) FindAll() ([]*customerroles.CustomerSlack, error) {
	return d.query.WithContext(d.ctx).CustomerSlack.Find()
}

// FindFirstById get first matched result by id.
func (d *CustomerSlackDao) FindFirstById(id uint) (*customerroles.CustomerSlack, error) {
	m := d.query.CustomerSlack

	return m.WithContext(d.ctx).Where(m.Id.Eq(id)).First()
}

// FindByIdPage get page by Id.
func (d *CustomerSlackDao) FindByIdPage(id uint, offset int, limit int) ([]*customerroles.CustomerSlack, int64, error) {
	m := d.query.CustomerSlack

	result, count, err := m.WithContext(d.ctx).Where(m.Id.Eq(id)).FindByPage(offset, limit)

	return result, count, err
}

// FindFirstByCreatedAt get first matched result by createdat.
func (d *CustomerSlackDao) FindFirstByCreatedAt(createdat time.Time) (*customerroles.CustomerSlack, error) {
	m := d.query.CustomerSlack

	return m.WithContext(d.ctx).Where(m.CreatedAt.Eq(createdat)).First()
}

// FindByCreatedAtPage get page by CreatedAt.
func (d *CustomerSlackDao) FindByCreatedAtPage(createdat time.Time, offset int, limit int) ([]*customerroles.CustomerSlack, int64, error) {
	m := d.query.CustomerSlack

	result, count, err := m.WithContext(d.ctx).Where(m.CreatedAt.Eq(createdat)).FindByPage(offset, limit)

	return result, count, err
}

// FindFirstByUpdatedAt get first matched result by updatedat.
func (d *CustomerSlackDao) FindFirstByUpdatedAt(updatedat time.Time) (*customerroles.CustomerSlack, error) {
	m := d.query.CustomerSlack

	return m.WithContext(d.ctx).Where(m.UpdatedAt.Eq(updatedat)).First()
}

// FindByUpdatedAtPage get page by UpdatedAt.
func (d *CustomerSlackDao) FindByUpdatedAtPage(updatedat time.Time, offset int, limit int) ([]*customerroles.CustomerSlack, int64, error) {
	m := d.query.CustomerSlack

	result, count, err := m.WithContext(d.ctx).Where(m.UpdatedAt.Eq(updatedat)).FindByPage(offset, limit)

	return result, count, err
}

// FindFirstByDeletedAt get first matched result by deletedat.
func (d *CustomerSlackDao) FindFirstByDeletedAt(deletedat time.Time) (*customerroles.CustomerSlack, error) {
	m := d.query.CustomerSlack

	return m.WithContext(d.ctx).Where(m.DeletedAt.Eq(deletedat)).First()
}

// FindByDeletedAtPage get page by DeletedAt.
func (d *CustomerSlackDao) FindByDeletedAtPage(deletedat time.Time, offset int, limit int) ([]*customerroles.CustomerSlack, int64, error) {
	m := d.query.CustomerSlack

	result, count, err := m.WithContext(d.ctx).Where(m.DeletedAt.Eq(deletedat)).FindByPage(offset, limit)

	return result, count, err
}

// FindFirstByUUID get first matched result by uuid.
func (d *CustomerSlackDao) FindFirstByUUID(uuid string) (*customerroles.CustomerSlack, error) {
	m := d.query.CustomerSlack

	return m.WithContext(d.ctx).Where(m.UUID.Eq(uuid)).First()
}

// FindByUUIDPage get page by UUID.
func (d *CustomerSlackDao) FindByUUIDPage(uuid string, offset int, limit int) ([]*customerroles.CustomerSlack, int64, error) {
	m := d.query.CustomerSlack

	result, count, err := m.WithContext(d.ctx).Where(m.UUID.Eq(uuid)).FindByPage(offset, limit)

	return result, count, err
}

// FindFirstByEmail get first matched result by email.
func (d *CustomerSlackDao) FindFirstByEmail(email string) (*customerroles.CustomerSlack, error) {
	m := d.query.CustomerSlack

	return m.WithContext(d.ctx).Where(m.Email.Eq(email)).First()
}

// FindByEmailPage get page by Email.
func (d *CustomerSlackDao) FindByEmailPage(email string, offset int, limit int) ([]*customerroles.CustomerSlack, int64, error) {
	m := d.query.CustomerSlack

	result, count, err := m.WithContext(d.ctx).Where(m.Email.Eq(email)).FindByPage(offset, limit)

	return result, count, err
}

// FindFirstByNickname get first matched result by nickname.
func (d *CustomerSlackDao) FindFirstByNickname(nickname string) (*customerroles.CustomerSlack, error) {
	m := d.query.CustomerSlack

	return m.WithContext(d.ctx).Where(m.Nickname.Eq(nickname)).First()
}

// FindByNicknamePage get page by Nickname.
func (d *CustomerSlackDao) FindByNicknamePage(nickname string, offset int, limit int) ([]*customerroles.CustomerSlack, int64, error) {
	m := d.query.CustomerSlack

	result, count, err := m.WithContext(d.ctx).Where(m.Nickname.Eq(nickname)).FindByPage(offset, limit)

	return result, count, err
}

// FindFirstByAvatar get first matched result by avatar.
func (d *CustomerSlackDao) FindFirstByAvatar(avatar string) (*customerroles.CustomerSlack, error) {
	m := d.query.CustomerSlack

	return m.WithContext(d.ctx).Where(m.Avatar.Eq(avatar)).First()
}

// FindByAvatarPage get page by Avatar.
func (d *CustomerSlackDao) FindByAvatarPage(avatar string, offset int, limit int) ([]*customerroles.CustomerSlack, int64, error) {
	m := d.query.CustomerSlack

	result, count, err := m.WithContext(d.ctx).Where(m.Avatar.Eq(avatar)).FindByPage(offset, limit)

	return result, count, err
}

// FindFirstByDescription get first matched result by description.
func (d *CustomerSlackDao) FindFirstByDescription(description string) (*customerroles.CustomerSlack, error) {
	m := d.query.CustomerSlack

	return m.WithContext(d.ctx).Where(m.Description.Eq(description)).First()
}

// FindByDescriptionPage get page by Description.
func (d *CustomerSlackDao) FindByDescriptionPage(description string, offset int, limit int) ([]*customerroles.CustomerSlack, int64, error) {
	m := d.query.CustomerSlack

	result, count, err := m.WithContext(d.ctx).Where(m.Description.Eq(description)).FindByPage(offset, limit)

	return result, count, err
}

// FindFirstByFirstName get first matched result by firstname.
func (d *CustomerSlackDao) FindFirstByFirstName(firstname string) (*customerroles.CustomerSlack, error) {
	m := d.query.CustomerSlack

	return m.WithContext(d.ctx).Where(m.FirstName.Eq(firstname)).First()
}

// FindByFirstNamePage get page by FirstName.
func (d *CustomerSlackDao) FindByFirstNamePage(firstname string, offset int, limit int) ([]*customerroles.CustomerSlack, int64, error) {
	m := d.query.CustomerSlack

	result, count, err := m.WithContext(d.ctx).Where(m.FirstName.Eq(firstname)).FindByPage(offset, limit)

	return result, count, err
}

// FindFirstByLastName get first matched result by lastname.
func (d *CustomerSlackDao) FindFirstByLastName(lastname string) (*customerroles.CustomerSlack, error) {
	m := d.query.CustomerSlack

	return m.WithContext(d.ctx).Where(m.LastName.Eq(lastname)).First()
}

// FindByLastNamePage get page by LastName.
func (d *CustomerSlackDao) FindByLastNamePage(lastname string, offset int, limit int) ([]*customerroles.CustomerSlack, int64, error) {
	m := d.query.CustomerSlack

	result, count, err := m.WithContext(d.ctx).Where(m.LastName.Eq(lastname)).FindByPage(offset, limit)

	return result, count, err
}

// FindFirstByRealName get first matched result by realname.
func (d *CustomerSlackDao) FindFirstByRealName(realname string) (*customerroles.CustomerSlack, error) {
	m := d.query.CustomerSlack

	return m.WithContext(d.ctx).Where(m.RealName.Eq(realname)).First()
}

// FindByRealNamePage get page by RealName.
func (d *CustomerSlackDao) FindByRealNamePage(realname string, offset int, limit int) ([]*customerroles.CustomerSlack, int64, error) {
	m := d.query.CustomerSlack

	result, count, err := m.WithContext(d.ctx).Where(m.RealName.Eq(realname)).FindByPage(offset, limit)

	return result, count, err
}

// FindFirstByRealNameNormalized get first matched result by realnamenormalized.
func (d *CustomerSlackDao) FindFirstByRealNameNormalized(realnamenormalized string) (*customerroles.CustomerSlack, error) {
	m := d.query.CustomerSlack

	return m.WithContext(d.ctx).Where(m.RealNameNormalized.Eq(realnamenormalized)).First()
}

// FindByRealNameNormalizedPage get page by RealNameNormalized.
func (d *CustomerSlackDao) FindByRealNameNormalizedPage(realnamenormalized string, offset int, limit int) ([]*customerroles.CustomerSlack, int64, error) {
	m := d.query.CustomerSlack

	result, count, err := m.WithContext(d.ctx).Where(m.RealNameNormalized.Eq(realnamenormalized)).FindByPage(offset, limit)

	return result, count, err
}

// FindFirstByDisplayName get first matched result by displayname.
func (d *CustomerSlackDao) FindFirstByDisplayName(displayname string) (*customerroles.CustomerSlack, error) {
	m := d.query.CustomerSlack

	return m.WithContext(d.ctx).Where(m.DisplayName.Eq(displayname)).First()
}

// FindByDisplayNamePage get page by DisplayName.
func (d *CustomerSlackDao) FindByDisplayNamePage(displayname string, offset int, limit int) ([]*customerroles.CustomerSlack, int64, error) {
	m := d.query.CustomerSlack

	result, count, err := m.WithContext(d.ctx).Where(m.DisplayName.Eq(displayname)).FindByPage(offset, limit)

	return result, count, err
}

// FindFirstByDisplayNameNormalized get first matched result by displaynamenormalized.
func (d *CustomerSlackDao) FindFirstByDisplayNameNormalized(displaynamenormalized string) (*customerroles.CustomerSlack, error) {
	m := d.query.CustomerSlack

	return m.WithContext(d.ctx).Where(m.DisplayNameNormalized.Eq(displaynamenormalized)).First()
}

// FindByDisplayNameNormalizedPage get page by DisplayNameNormalized.
func (d *CustomerSlackDao) FindByDisplayNameNormalizedPage(displaynamenormalized string, offset int, limit int) ([]*customerroles.CustomerSlack, int64, error) {
	m := d.query.CustomerSlack

	result, count, err := m.WithContext(d.ctx).Where(m.DisplayNameNormalized.Eq(displaynamenormalized)).FindByPage(offset, limit)

	return result, count, err
}

// FindFirstBySkype get first matched result by skype.
func (d *CustomerSlackDao) FindFirstBySkype(skype string) (*customerroles.CustomerSlack, error) {
	m := d.query.CustomerSlack

	return m.WithContext(d.ctx).Where(m.Skype.Eq(skype)).First()
}

// FindBySkypePage get page by Skype.
func (d *CustomerSlackDao) FindBySkypePage(skype string, offset int, limit int) ([]*customerroles.CustomerSlack, int64, error) {
	m := d.query.CustomerSlack

	result, count, err := m.WithContext(d.ctx).Where(m.Skype.Eq(skype)).FindByPage(offset, limit)

	return result, count, err
}

// FindFirstByPhone get first matched result by phone.
func (d *CustomerSlackDao) FindFirstByPhone(phone string) (*customerroles.CustomerSlack, error) {
	m := d.query.CustomerSlack

	return m.WithContext(d.ctx).Where(m.Phone.Eq(phone)).First()
}

// FindByPhonePage get page by Phone.
func (d *CustomerSlackDao) FindByPhonePage(phone string, offset int, limit int) ([]*customerroles.CustomerSlack, int64, error) {
	m := d.query.CustomerSlack

	result, count, err := m.WithContext(d.ctx).Where(m.Phone.Eq(phone)).FindByPage(offset, limit)

	return result, count, err
}

// FindFirstByImage24 get first matched result by image24.
func (d *CustomerSlackDao) FindFirstByImage24(image24 string) (*customerroles.CustomerSlack, error) {
	m := d.query.CustomerSlack

	return m.WithContext(d.ctx).Where(m.Image24.Eq(image24)).First()
}

// FindByImage24Page get page by Image24.
func (d *CustomerSlackDao) FindByImage24Page(image24 string, offset int, limit int) ([]*customerroles.CustomerSlack, int64, error) {
	m := d.query.CustomerSlack

	result, count, err := m.WithContext(d.ctx).Where(m.Image24.Eq(image24)).FindByPage(offset, limit)

	return result, count, err
}

// FindFirstByImage32 get first matched result by image32.
func (d *CustomerSlackDao) FindFirstByImage32(image32 string) (*customerroles.CustomerSlack, error) {
	m := d.query.CustomerSlack

	return m.WithContext(d.ctx).Where(m.Image32.Eq(image32)).First()
}

// FindByImage32Page get page by Image32.
func (d *CustomerSlackDao) FindByImage32Page(image32 string, offset int, limit int) ([]*customerroles.CustomerSlack, int64, error) {
	m := d.query.CustomerSlack

	result, count, err := m.WithContext(d.ctx).Where(m.Image32.Eq(image32)).FindByPage(offset, limit)

	return result, count, err
}

// FindFirstByImage48 get first matched result by image48.
func (d *CustomerSlackDao) FindFirstByImage48(image48 string) (*customerroles.CustomerSlack, error) {
	m := d.query.CustomerSlack

	return m.WithContext(d.ctx).Where(m.Image48.Eq(image48)).First()
}

// FindByImage48Page get page by Image48.
func (d *CustomerSlackDao) FindByImage48Page(image48 string, offset int, limit int) ([]*customerroles.CustomerSlack, int64, error) {
	m := d.query.CustomerSlack

	result, count, err := m.WithContext(d.ctx).Where(m.Image48.Eq(image48)).FindByPage(offset, limit)

	return result, count, err
}

// FindFirstByImage72 get first matched result by image72.
func (d *CustomerSlackDao) FindFirstByImage72(image72 string) (*customerroles.CustomerSlack, error) {
	m := d.query.CustomerSlack

	return m.WithContext(d.ctx).Where(m.Image72.Eq(image72)).First()
}

// FindByImage72Page get page by Image72.
func (d *CustomerSlackDao) FindByImage72Page(image72 string, offset int, limit int) ([]*customerroles.CustomerSlack, int64, error) {
	m := d.query.CustomerSlack

	result, count, err := m.WithContext(d.ctx).Where(m.Image72.Eq(image72)).FindByPage(offset, limit)

	return result, count, err
}

// FindFirstByImage192 get first matched result by image192.
func (d *CustomerSlackDao) FindFirstByImage192(image192 string) (*customerroles.CustomerSlack, error) {
	m := d.query.CustomerSlack

	return m.WithContext(d.ctx).Where(m.Image192.Eq(image192)).First()
}

// FindByImage192Page get page by Image192.
func (d *CustomerSlackDao) FindByImage192Page(image192 string, offset int, limit int) ([]*customerroles.CustomerSlack, int64, error) {
	m := d.query.CustomerSlack

	result, count, err := m.WithContext(d.ctx).Where(m.Image192.Eq(image192)).FindByPage(offset, limit)

	return result, count, err
}

// FindFirstByImage512 get first matched result by image512.
func (d *CustomerSlackDao) FindFirstByImage512(image512 string) (*customerroles.CustomerSlack, error) {
	m := d.query.CustomerSlack

	return m.WithContext(d.ctx).Where(m.Image512.Eq(image512)).First()
}

// FindByImage512Page get page by Image512.
func (d *CustomerSlackDao) FindByImage512Page(image512 string, offset int, limit int) ([]*customerroles.CustomerSlack, int64, error) {
	m := d.query.CustomerSlack

	result, count, err := m.WithContext(d.ctx).Where(m.Image512.Eq(image512)).FindByPage(offset, limit)

	return result, count, err
}

// FindFirstByImageOriginal get first matched result by imageoriginal.
func (d *CustomerSlackDao) FindFirstByImageOriginal(imageoriginal string) (*customerroles.CustomerSlack, error) {
	m := d.query.CustomerSlack

	return m.WithContext(d.ctx).Where(m.ImageOriginal.Eq(imageoriginal)).First()
}

// FindByImageOriginalPage get page by ImageOriginal.
func (d *CustomerSlackDao) FindByImageOriginalPage(imageoriginal string, offset int, limit int) ([]*customerroles.CustomerSlack, int64, error) {
	m := d.query.CustomerSlack

	result, count, err := m.WithContext(d.ctx).Where(m.ImageOriginal.Eq(imageoriginal)).FindByPage(offset, limit)

	return result, count, err
}

// FindFirstByTitle get first matched result by title.
func (d *CustomerSlackDao) FindFirstByTitle(title string) (*customerroles.CustomerSlack, error) {
	m := d.query.CustomerSlack

	return m.WithContext(d.ctx).Where(m.Title.Eq(title)).First()
}

// FindByTitlePage get page by Title.
func (d *CustomerSlackDao) FindByTitlePage(title string, offset int, limit int) ([]*customerroles.CustomerSlack, int64, error) {
	m := d.query.CustomerSlack

	result, count, err := m.WithContext(d.ctx).Where(m.Title.Eq(title)).FindByPage(offset, limit)

	return result, count, err
}

// FindFirstByBotID get first matched result by botid.
func (d *CustomerSlackDao) FindFirstByBotID(botid string) (*customerroles.CustomerSlack, error) {
	m := d.query.CustomerSlack

	return m.WithContext(d.ctx).Where(m.BotID.Eq(botid)).First()
}

// FindByBotIDPage get page by BotID.
func (d *CustomerSlackDao) FindByBotIDPage(botid string, offset int, limit int) ([]*customerroles.CustomerSlack, int64, error) {
	m := d.query.CustomerSlack

	result, count, err := m.WithContext(d.ctx).Where(m.BotID.Eq(botid)).FindByPage(offset, limit)

	return result, count, err
}

// FindFirstByApiAppID get first matched result by apiappid.
func (d *CustomerSlackDao) FindFirstByApiAppID(apiappid string) (*customerroles.CustomerSlack, error) {
	m := d.query.CustomerSlack

	return m.WithContext(d.ctx).Where(m.ApiAppID.Eq(apiappid)).First()
}

// FindByApiAppIDPage get page by ApiAppID.
func (d *CustomerSlackDao) FindByApiAppIDPage(apiappid string, offset int, limit int) ([]*customerroles.CustomerSlack, int64, error) {
	m := d.query.CustomerSlack

	result, count, err := m.WithContext(d.ctx).Where(m.ApiAppID.Eq(apiappid)).FindByPage(offset, limit)

	return result, count, err
}

// FindFirstByStatusText get first matched result by statustext.
func (d *CustomerSlackDao) FindFirstByStatusText(statustext string) (*customerroles.CustomerSlack, error) {
	m := d.query.CustomerSlack

	return m.WithContext(d.ctx).Where(m.StatusText.Eq(statustext)).First()
}

// FindByStatusTextPage get page by StatusText.
func (d *CustomerSlackDao) FindByStatusTextPage(statustext string, offset int, limit int) ([]*customerroles.CustomerSlack, int64, error) {
	m := d.query.CustomerSlack

	result, count, err := m.WithContext(d.ctx).Where(m.StatusText.Eq(statustext)).FindByPage(offset, limit)

	return result, count, err
}

// FindFirstByStatusEmoji get first matched result by statusemoji.
func (d *CustomerSlackDao) FindFirstByStatusEmoji(statusemoji string) (*customerroles.CustomerSlack, error) {
	m := d.query.CustomerSlack

	return m.WithContext(d.ctx).Where(m.StatusEmoji.Eq(statusemoji)).First()
}

// FindByStatusEmojiPage get page by StatusEmoji.
func (d *CustomerSlackDao) FindByStatusEmojiPage(statusemoji string, offset int, limit int) ([]*customerroles.CustomerSlack, int64, error) {
	m := d.query.CustomerSlack

	result, count, err := m.WithContext(d.ctx).Where(m.StatusEmoji.Eq(statusemoji)).FindByPage(offset, limit)

	return result, count, err
}

// FindFirstByStatusExpiration get first matched result by statusexpiration.
func (d *CustomerSlackDao) FindFirstByStatusExpiration(statusexpiration int) (*customerroles.CustomerSlack, error) {
	m := d.query.CustomerSlack

	return m.WithContext(d.ctx).Where(m.StatusExpiration.Eq(statusexpiration)).First()
}

// FindByStatusExpirationPage get page by StatusExpiration.
func (d *CustomerSlackDao) FindByStatusExpirationPage(statusexpiration int, offset int, limit int) ([]*customerroles.CustomerSlack, int64, error) {
	m := d.query.CustomerSlack

	result, count, err := m.WithContext(d.ctx).Where(m.StatusExpiration.Eq(statusexpiration)).FindByPage(offset, limit)

	return result, count, err
}

// FindFirstByTeam get first matched result by team.
func (d *CustomerSlackDao) FindFirstByTeam(team string) (*customerroles.CustomerSlack, error) {
	m := d.query.CustomerSlack

	return m.WithContext(d.ctx).Where(m.Team.Eq(team)).First()
}

// FindByTeamPage get page by Team.
func (d *CustomerSlackDao) FindByTeamPage(team string, offset int, limit int) ([]*customerroles.CustomerSlack, int64, error) {
	m := d.query.CustomerSlack

	result, count, err := m.WithContext(d.ctx).Where(m.Team.Eq(team)).FindByPage(offset, limit)

	return result, count, err
}

// Update update model.
func (d *CustomerSlackDao) Update(m *customerroles.CustomerSlack) error {
	res, err := d.query.WithContext(d.ctx).CustomerSlack.Updates(m)
	if err != nil && res.Error != nil {
		return err
	}

	return nil
}

// Delete delete model.
func (d *CustomerSlackDao) Delete(m ...*customerroles.CustomerSlack) error {
	res, err := d.query.WithContext(d.ctx).CustomerSlack.Delete(m...)
	if err != nil && res.Error != nil {
		return err
	}

	return nil
}

// Count count matched records.
func (d *CustomerSlackDao) Count() (int64, error) {
	return d.query.WithContext(d.ctx).CustomerSlack.Count()
}

///////////////////////////////////////////////////////////
//              Append your code here.                   //
///////////////////////////////////////////////////////////
