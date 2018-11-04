/*
 *
 * Copyright 2015 gRPC authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

//go:generate protoc -I ../helloworld --go_out=plugins=grpc:../helloworld ../helloworld/helloworld.proto

//go:protoc -I=protos/ --go_out=plugins=grpc:protobuf protos/EduData.proto

package main

import (
	"context"
	ed "dk/protoTest/protobuf"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	port = ":8000"
)

// server is used to implement helloworld.GreeterServer.
//type server struct{}

type server struct{}

// SayHello implements helloworld.GreeterServer
// func (s *server) SetSubject(ctx context.Context, in *ed.EduData_Subject) (*ed.Reply, error) {
// 	//	return &ed.Reply{Message: in.GetSubjectName}, nil
// 	return &ed.Reply{Message: in.String()}, nil
// }
func (s *server) SetData(ctx context.Context, in *ed.EduData) (*ed.EduData, error) {
	return in, nil
	//return &ed.Reply{Message: in}, nil
}

// func (s *server) SetHistory(ctx context.Context, in []*ed.EduData) ([]*ed.EduData, error) {
// 	return in, nil
// }

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	ed.RegisterSetDataServer(s, &server{})
	// Register reflection service on gRPC server.
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
