/*
 * Copyright (c) 2023 sixwaaaay
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *     http://www.apache.org/licenses/LICENSE-2.
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package logic

import (
	"github.com/gocql/gocql"
	"github.com/sixwaaaay/shamessage/internal/config"
	"github.com/sixwaaaay/shamessage/internal/data"
	"testing"
)

var session *gocql.Session

func TestMain(m *testing.M) {
	c := config.Config{}
	c.Cluster = []string{"localhost:9042"}
	c.KeySpace = "messages"
	var err error
	session, err = data.CreateGoCqlSession(&c)
	if err != nil {
		panic(err)
	}
	defer session.Close()
	m.Run()
}
