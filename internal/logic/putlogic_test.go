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
	"github.com/sixwaaaay/shamessage/message"
	"testing"
)

func TestMessageService_Put(t *testing.T) {
	// create a PutLogic service
	service := NewPutLogic(PutLogicOption{
		Session: session,
	})

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
}
