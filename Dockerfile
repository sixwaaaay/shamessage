# Copyright (c) 2023 sixwaaaay
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#     http:www.apache.org/licenses/LICENSE-2.
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.
#
FROM golang:1.20-alpine AS builder
WORKDIR /app
ENV CGO_ENABLED 0

#ENV GOPROXY https://goproxy.cn,direct
#sed -i 's/dl-cdn.alpinelinux.org/mirrors.ustc.edu.cn/g' /etc/apk/repositories

RUN apk update --no-cache && apk add --no-cache tzdata
WORKDIR /build

COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
RUN go install github.com/google/wire/cmd/wire@latest && wire

RUN go build -ldflags="-s -w" -o /app/shamessage

FROM scratch
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY --from=builder /usr/share/zoneinfo/Asia/Shanghai /usr/share/zoneinfo/Asia/Shanghai
ENV TZ Asia/Shanghai
WORKDIR /app
COPY --from=builder /app/shamessage /app/shamessage
# copy default config
COPY configs /app/configs

# default port
EXPOSE 8080

CMD ["./shamessage"]