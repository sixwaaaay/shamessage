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

package config

import (
	"github.com/stretchr/testify/assert"
	"os"
	"path/filepath"
	"testing"
)

func TestNewConfig(t *testing.T) {
	fcontent := `
ListenOn: 0.0.0.0:8080
Cluster:
  - localhost:9042
KeySpace: messages
`
	assertions := assert.New(t)
	dir := os.TempDir()
	path := filepath.Join(dir, "test_config.yaml")
	f, err := os.Create(path)
	assertions.NoError(err)

	_, err = f.WriteString(fcontent)
	assertions.NoError(err)

	c, err := NewConfig(path)
	assertions.NoError(err)
	assertions.Equal("0.0.0.0:8080", c.ListenOn)
	assertions.Equal(1, len(c.Cluster))
	assertions.Equal("localhost:9042", c.Cluster[0])
	assertions.Equal("messages", c.KeySpace)
	assertions.NoError(f.Close())
	assertions.NoError(os.Remove(path))
}
