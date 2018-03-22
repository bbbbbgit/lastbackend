//
// Last.Backend LLC CONFIDENTIAL
// __________________
//
// [2014] - [2018] Last.Backend LLC
// All Rights Reserved.
//
// NOTICE:  All information contained herein is, and remains
// the property of Last.Backend LLC and its suppliers,
// if any.  The intellectual and technical concepts contained
// herein are proprietary to Last.Backend LLC
// and its suppliers and may be covered by Russian Federation and Foreign Patents,
// patents in process, and are protected by trade secret or copyright law.
// Dissemination of this information or reproduction of this material
// is strictly forbidden unless prior written permission is obtained
// from Last.Backend LLC.
//

package client

import (
	"github.com/lastbackend/lastbackend/pkg/api/client/http"
	"github.com/lastbackend/lastbackend/pkg/api/client/v1"
	"net/url"
)

func New(endpoint string, config *http.Config) (*Client, error) {
	c := new(Client)

	baseURL, err := url.Parse(endpoint)
	if err != nil {
		return nil, err
	}

	if config == nil {
		config = new(http.Config)
		config.SetDefault()
	}

	client, err := http.NewRESTClient(baseURL, config)
	if err != nil {
		return nil, err
	}
	c.client = client
	return c, nil
}

type Client struct {
	client http.Interface
}

func (c Client) V1() IClientV1 {
	return v1.New(c.client)
}
