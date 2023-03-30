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

import "encoding/binary"

func genKey(userID, toUserID int64) []byte {
	var first, second int64
	if userID < toUserID {
		first, second = userID, toUserID
	} else {
		first, second = toUserID, userID
	}
	// the partition key is ordered
	// first always less than second
	partitionKey := make([]byte, 16)
	binary.BigEndian.PutUint64(partitionKey[:8], uint64(first))
	binary.BigEndian.PutUint64(partitionKey[8:], uint64(second))
	return partitionKey
}
