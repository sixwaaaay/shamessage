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
	"context"

	"github.com/sixwaaaay/shamessage/internal/config"
	"github.com/sixwaaaay/shamessage/message"
)

// PutLogic is the logic for Put method implementation
type PutLogic struct {
	conf *config.Config
}

// PutLogicOption is the parameters for PutLogic
type PutLogicOption struct {
	Config *config.Config
}

// NewPutLogic creates a new PutLogic
func NewPutLogic(opt PutLogicOption) *PutLogic {
	return &PutLogic{
		conf: opt.Config,
	}
}

// Put is the logic for Put method
func (l *PutLogic) Put(ctx context.Context, in *message.MessageActionRequest) (*message.MessageActionResponse, error) {

	return &message.MessageActionResponse{}, nil
}
