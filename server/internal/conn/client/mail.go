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

package client

import (
	"fmt"
	"net/smtp"

	"github.com/jordan-wright/email"

	"github.com/OpenIMSDK/OpenKF/server/internal/config"
	"github.com/OpenIMSDK/OpenKF/server/pkg/log"
)

// todo: use email connection pool to reduce the cost of creating a connection
// link: https://github.com/jordan-wright/email#a-pool-of-reusable-connections
var (
// _mailPool *email.Pool
)

// InitMail init mail.
func InitMail() {
	// Unused now
	// _mailPool, _ = email.NewPool(
	// 	fmt.Sprintf("%s:%d", config.Config.Email.Host, config.Config.Email.Port),
	// 	config.Config.Email.PoolSize, // Max number of connections to open
	// 	smtp.PlainAuth(
	// 		"", config.Config.Email.From, config.Config.Email.Password, config.Config.Email.Host,
	// 	),
	// )

	// Send test email
	if err := SendEmail(config.Config.Email.From, "OpenKF Init Email", "OpenKF Email Init Test..."); err != nil {
		log.Panicf("Email", "Init Email failed:%s", err.Error())
	}
}

// SendEmail send email.
func SendEmail(to string, subject string, body string) error {
	email := email.NewEmail()

	email.From = fmt.Sprintf("%s<%s>", config.Config.Email.Nickname, config.Config.Email.From)
	email.To = []string{to}
	email.Subject = subject
	email.Text = []byte(body)

	if err := email.Send(
		fmt.Sprintf("%s:%d", config.Config.Email.Host, config.Config.Email.Port),
		smtp.PlainAuth(
			"", config.Config.Email.From, config.Config.Email.Password, config.Config.Email.Host,
		),
	); err != nil {
		return err
	}

	return nil
}

// SendHtmlEmail send html email.
func SendHtmlEmail(to string, subject string, html string) error {
	email := email.NewEmail()

	email.From = config.Config.Email.From
	email.To = []string{to}
	email.Subject = subject
	email.HTML = []byte(html)

	if err := email.Send(
		fmt.Sprintf("%s:%d", config.Config.Email.Host, config.Config.Email.Port),
		smtp.PlainAuth(
			"", config.Config.Email.From, config.Config.Email.Password, config.Config.Email.Host,
		),
	); err != nil {
		return err
	}

	return nil
}
