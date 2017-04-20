//
// Last.Backend LLC CONFIDENTIAL
// __________________
//
// [2014] - [2017] Last.Backend LLC
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

package namespace

import (
	"fmt"
	c "github.com/lastbackend/lastbackend/pkg/client/context"
	n "github.com/lastbackend/lastbackend/pkg/daemon/namespace/views/v1"
	"github.com/lastbackend/lastbackend/pkg/errors"
)

func ListNamespaceCmd() {

	namspaceList, err := List()
	if err != nil {
		fmt.Print(err)
		return
	}

	if namspaceList != nil {
		namspaceList.DrawTable()
	}
}

func List() (*n.NamespaceList, error) {

	var (
		err        error
		http       = c.Get().GetHttpClient()
		er         = new(errors.Http)
		nspaceList = new(n.NamespaceList)
	)

	_, _, err = http.
		GET("/namespace").
		Request(nspaceList, er)
	if err != nil {
		return nil, errors.New(er.Message)
	}

	if er.Code == 401 {
		return nil, errors.NotLoggedMessage
	}

	if er.Code != 0 {
		return nil, errors.New(er.Message)
	}

	if len(*nspaceList) == 0 {
		return nil, errors.New("You don't have any namespace")
	}

	return nspaceList, nil
}
