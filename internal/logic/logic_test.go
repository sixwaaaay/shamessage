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
	"testing"
)

func TestMessageService_Put(t *testing.T) {
	// create a PutLogic service
	service := NewPutLogic(PutLogicOption{
		Session: session,
	})

	t.Run("put message invalid argument", func(t *testing.T) {
		var requests = []*message.MessageActionRequest{
			{
				UserId:     0,
				ToUserId:   1,
				ActionType: 1,
			},
			{
				UserId:     1,
				ToUserId:   0,
				ActionType: 1,
			},
			{
				UserId:     1,
				ToUserId:   2,
				ActionType: 0,
			},
			{
				UserId:     1,
				ToUserId:   1,
				ActionType: 1,
			},
		}
		for _, req := range requests {
			// call Put method
			_, err := service.Put(context.Background(), req)
			if err == nil {
				t.Fatalf("Put not failed: %v", err)
			}
		}
	})

	t.Run("put message", func(t *testing.T) {
		// create a message action request
		req := &message.MessageActionRequest{
			UserId:     13411111111111111,
			ToUserId:   25215555,
			ActionType: 1,
			Content:    "Hello, world!",
		}

		// call Put method
		resp, err := service.Put(context.Background(), req)
		if err != nil {
			t.Fatalf("Put failed: %v", err)
		}

		// check response status code
		if resp.StatusCode != 0 {
			t.Fatalf("Put returned status code %d, expected 0", resp.StatusCode)
		}
	})
	c := config.Config{}
	c.Limit = 30
	listLogic := NewListLogic(ListLogicOption{
		Session: session,
		Config:  &c,
	})

	t.Run("list message", func(t *testing.T) {
		// create a message action request
		req := &message.MessageListRequest{
			UserId:   13411111111111111,
			ToUserId: 25215555,
		}

		// call Put method
		resp, err := listLogic.List(context.Background(), req)
		if err != nil {
			t.Fatalf("List failed: %v", err)
		}

		// check response status code
		if resp.StatusCode != 0 {
			t.Fatalf("List returned status code %d, expected 0", resp.StatusCode)
		}
	})
	t.Run("list message invalid argument", func(t *testing.T) {
		var requests = []*message.MessageListRequest{
			{
				UserId:   0,
				ToUserId: 1,
			},
			{
				UserId:   1,
				ToUserId: 0,
			},
			{
				UserId:   1,
				ToUserId: 1,
			},
		}
		for _, req := range requests {
			// call Put method
			_, err := listLogic.List(context.Background(), req)
			if err == nil {
				t.Fatalf("List not failed: %v", err)
			}
		}
	})
}
