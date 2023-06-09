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

import { check, sleep } from "k6";
import grpc from "k6/net/grpc";

const client = new grpc.Client();
client.load([".."], "message.proto");
export default () => {
  client.connect("localhost:8080", {
    // plaintext: false
    plaintext: true,
  });

  const data = {
    user_id: randomInt(1, 100000),
    to_user_id: randomInt(1, 100000),
    action_type: 1,
    content: randomString(15),
  };

  const response = client.invoke("message.MessageService/Put", data);

  check(response, {
    "status is OK": (r) => r && r.status === grpc.StatusOK,
  });

  //    console.log(JSON.stringify(response.message));

  client.close();
};

// generate random int
const randomInt = (min, max) => {
  return Math.floor(Math.random() * (max - min + 1) + min);
};

const characters =
  "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789";
// generate random string
const randomString = (length) => {
  let result = "";
  const charactersLength = characters.length;
  for (let i = 0; i < length; i++) {
    result += characters.charAt(Math.floor(Math.random() * charactersLength));
  }
  return result;
};
