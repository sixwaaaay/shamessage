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

// Package logic is the implementations of each method
package logic

import (
	"context"

	"github.com/sixwaaaay/shamessage/internal/config"
	"github.com/sixwaaaay/shamessage/message"
)

// ListLogic is the logic for List method implementation
type ListLogic struct {
	conf *config.Config
}

// ListLogicOption is the parameters for ListLogic
type ListLogicOption struct {
	Config *config.Config
}

// NewListLogic creates a new ListLogic
func NewListLogic(opt ListLogicOption) *ListLogic {
	return &ListLogic{
		conf: opt.Config,
	}
}

// List is the logic for List method
func (l *ListLogic) List(ctx context.Context, in *message.MessageListRequest) (*message.MessageListResponse, error) {

	return &message.MessageListResponse{}, nil
}
