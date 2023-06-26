// Copyright © 2023 OpenIMSDK open source community. All rights reserved.
// Licensed under the MIT License (the "License");
// you may not use this file except in compliance with the License.

package client

import (
	"fmt"
	"net/smtp"

	"github.com/OpenIMSDK/OpenKF/server/internal/config"
	"github.com/jordan-wright/email"
)

// todo: use email connection pool to reduce the cost of creating a connection
// link: https://github.com/jordan-wright/email#a-pool-of-reusable-connections

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
