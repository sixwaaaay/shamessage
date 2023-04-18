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
	"github.com/gocql/gocql"
	"github.com/sixwaaaay/shamessage/internal/config"
	"github.com/sixwaaaay/shamessage/message"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// ListLogic is the logic for List method implementation
type ListLogic struct {
	conf    *config.Config
	session *gocql.Session
}

// ListLogicOption is the parameters for ListLogic
type ListLogicOption struct {
	Config  *config.Config
	Session *gocql.Session
}

// NewListLogic creates a new ListLogic
func NewListLogic(opt ListLogicOption) *ListLogic {
	return &ListLogic{
		conf:    opt.Config,
		session: opt.Session,
	}
}

// List is the logic for List method
func (l *ListLogic) List(ctx context.Context, req *message.MessageListRequest) (*message.MessageListResponse, error) {
	if req.UserId == 0 || req.ToUserId == 0 || req.UserId == req.ToUserId {
		return nil, status.Errorf(
			codes.InvalidArgument,
			"invalid argument",
		)
	}

	var partitionKey = genKey(req.UserId, req.ToUserId)
	var (
		messageList []*message.Message
		stmt        = l.session.Query(`
		SELECT id, to_user_id, user_id, content, created_at
		FROM messages
		WHERE channel_id = ? and id > ?
		LIMIT ?
		    `,
			partitionKey, req.Token, l.conf.Limit)
	)

	iterator := stmt.WithContext(ctx).Iter()

	for {
		m := &message.Message{}
		if iterator.Scan(&m.Id, &m.ToUserId, &m.FromUserId, &m.Content, &m.CreateTime) {
			messageList = append(messageList, m)
		} else {
			break
		}
	}
	err := iterator.Close()
	if err != nil {
		return nil, status.Errorf(codes.Internal, "internal error")
	}

	return &message.MessageListResponse{
		StatusCode:  OK,
		StatusMsg:   "OK",
		MessageList: messageList,
	}, nil
}
