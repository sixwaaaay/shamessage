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
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"time"

	"github.com/gocql/gocql"
	"github.com/sixwaaaay/shamessage/internal/config"
	"github.com/sixwaaaay/shamessage/message"
	"github.com/sony/sonyflake"
)

// PutLogic is the logic for Put method implementation
type PutLogic struct {
	conf    *config.Config
	flake   *sonyflake.Sonyflake
	session *gocql.Session
}

// PutLogicOption is the parameters for PutLogic
type PutLogicOption struct {
	Config  *config.Config
	Session *gocql.Session
}

// NewPutLogic creates a new PutLogic
func NewPutLogic(opt PutLogicOption) *PutLogic {
	return &PutLogic{
		conf:    opt.Config,
		flake:   sonyflake.NewSonyflake(sonyflake.Settings{}),
		session: opt.Session,
	}
}

// Put is the logic for Put method
func (l *PutLogic) Put(ctx context.Context, in *message.MessageActionRequest) (*message.MessageActionResponse, error) {
	if in.UserId == 0 || in.ToUserId == 0 || in.ActionType != 1 || in.Content == "" {
		return nil, status.Errorf(codes.InvalidArgument, "invalid argument")
	}

	id, err := l.flake.NextID()
	if err != nil {
		return nil, err
	}

	channel := genKey(in.UserId, in.ToUserId)

	now := time.Now()

	query := l.session.Query(`
        INSERT INTO messages (channel_id, id, user_id, to_user_id, action_type, content, created_at)
        VALUES (?,?, ?, ?, ?, ?, ?)
    `, channel, id, in.UserId, in.ToUserId, in.ActionType, in.Content, now)

	if err := query.Exec(); err != nil {
		return nil, status.Errorf(codes.Internal, "insert message failed: %v", err)
	}

	resp := &message.MessageActionResponse{
		StatusCode: 0,
		StatusMsg:  "success",
	}
	return resp, nil
}
