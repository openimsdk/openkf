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
	"encoding/json"

	"github.com/go-resty/resty/v2"
)

// Client client.
type Client interface {
	GET(operationID string, params interface{}) (map[string]interface{}, error)
	POST(operationID string, params interface{}) (map[string]interface{}, error)
}

// httpClient http client.
type httpClient struct {
	url    string
	client *resty.Client
}

// NewClient new client.
func NewClient(url string) Client {
	return &httpClient{
		url:    url,
		client: resty.New(),
	}
}

// GET get unimplemented.
func (c *httpClient) GET(operationID string, params interface{}) (map[string]interface{}, error) {
	return nil, nil
}

// POST post.
func (c *httpClient) POST(operationID string, params interface{}) (map[string]interface{}, error) {
	resp, err := c.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("operationID", operationID).
		SetBody(params).
		Post(c.url)
	if err != nil {
		return nil, err
	}

	var responseData map[string]interface{}
	err = json.Unmarshal(resp.Body(), &responseData)
	if err != nil {
		return nil, err
	}

	return responseData, nil
}
