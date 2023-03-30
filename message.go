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
	"flag"
	"github.com/sixwaaaay/shamessage/internal/config"
	"github.com/sixwaaaay/shamessage/internal/data"
	"github.com/sixwaaaay/shamessage/message"
	"google.golang.org/grpc"
	"net"
	"os"
	"os/signal"
	"syscall"
)

var configFile = flag.String("f", "configs/config.yaml", "the config file")

func main() {
	flag.Parse()
	c, err := config.NewConfig(*configFile)
	if err != nil {
		panic(err)
	}

	session, err := data.CreateGoCqlSession(c)
	if err != nil {
		panic(err)
	}
	defer session.Close()
	server, err := NewServer(c, session)
	if err != nil {
		panic(err)
	}
	grpcServer := grpc.NewServer()
	message.RegisterMessageServiceServer(grpcServer, server)
	ln, err := net.Listen("tcp", c.ListenOn)
	if err != nil {
		panic(err)
	}

	// graceful shutdown
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-ch
		grpcServer.GracefulStop()
	}()
	if err := grpcServer.Serve(ln); err != nil {
		panic(err)
	}
}
