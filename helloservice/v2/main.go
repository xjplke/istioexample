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

package main

import (
	"log"
	"net"
	"fmt"
	"os"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	pb "github.com/xjplke/istioexample/helloworld"
	"google.golang.org/grpc/reflection"
)


func determineListenAddress() (string) {
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatalf("PORT not set")
	}
	log.Println("Listen Address : " + ":"+port)
	return ":" + port
}
// server is used to implement helloworld.GreeterServer.
type server struct{}
var count int = 0
// SayHello implements helloworld.GreeterServer
func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	count++
	message := fmt.Sprintf("v2 Helo %s call count:%d", in.Name, count)
	return &pb.HelloReply{Message: message}, nil
}

func main() {
	lis, err := net.Listen("tcp", determineListenAddress())
	if err != nil {
		log.Fatalf("v2 failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &server{})
	// Register reflection service on gRPC server.
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("v2 failed to serve: %v", err)
	}
}