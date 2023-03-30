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

// Package config is project config mapping
package config

import (
	"github.com/spf13/viper"
)

// Config is the config struct
type Config struct {
	ListenOn string   `yaml:"listen_on"` // Listen on address
	Cluster  []string `yaml:"cluster"`   // Cassandra cluster addresses
	KeySpace string   `yaml:"keyspace"`  // Cassandra keyspace
	Limit    int64    // query limit
}

// NewConfig parses the config file and returns a Config struct
func NewConfig(path string) (*Config, error) {
	viper.SetConfigFile(path)
	viper.SetConfigType("yaml")
	viper.AutomaticEnv()
	var config Config
	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}
	if err := viper.Unmarshal(&config); err != nil {
		return nil, err
	}
	return &config, nil
}
