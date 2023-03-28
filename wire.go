//go:build wireinject

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

package main

import (
	"github.com/google/wire"
	"github.com/sixwaaaay/shamessage/internal/config"
	"github.com/sixwaaaay/shamessage/internal/logic"
	"github.com/sixwaaaay/shamessage/internal/server"
)

// NewServer creates a new server.
func NewServer(c *config.Config) (*server.MessageServiceServer, error) {
	wire.Build(
		server.NewMessageServiceServer,
		wire.Struct(new(server.ServerOption), "*"),
		logic.NewListLogic,
		wire.Struct(new(logic.ListLogicOption), "*"),
		logic.NewPutLogic,
		wire.Struct(new(logic.PutLogicOption), "*"),
	)
	return nil, nil
}
