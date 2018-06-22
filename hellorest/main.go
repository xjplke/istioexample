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

 package main

 import (
	"log"
	"os"
	"time"

	"github.com/gin-gonic/gin"
 
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	pb "google.golang.org/grpc/examples/helloworld/helloworld"
 )
 
 const (
	defaultName = "world"
 )

func determineListenAddress() (string) {
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatalf("PORT not set")
	}
	log.Println("Listen Address:"+" :"+port)
	return ":" + port
}


func discoveryService() (string) {
	service := os.Getenv("HELLO_SERVICE")
	if service == "" { 
		//service = "localhost"
		log.Fatalf("HELLO_SERVICE not set")
	}
	log.Println("HELLO_SERVICE : "+service)
	port := os.Getenv("HELLO_PORT")
	if port == "" {
		log.Fatalf("HELLO_PORT not set")
	}
	log.Println("HELLO_PORT :"+port)
	return service + ":" + port
}

func main() {
	listen := determineListenAddress()
	log.Println("Listen address = ",listen)
	// Set up a connection to the server.
	address := discoveryService()
	log.Println("server address = ",address)
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	client := pb.NewGreeterClient(conn)

	 // pre test
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	rsp, rex := client.SayHello(ctx, &pb.HelloRequest{Name: "xxxx"})
	if rex != nil {
		log.Fatal("rex = ", rex.Error())
	}else{
		log.Println("rsp.Message = ", rsp.Message)
	}
	cancel()

	r := gin.Default()
	r.GET("/hello/:name", func(gctx *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		name := gctx.Param("name")
		rsp, err := client.SayHello(ctx, &pb.HelloRequest{Name: name})
	 	if err != nil {
			gctx.JSON(401, gin.H{
				"error": err.Error(),
			})
	 	}else{
			log.Printf("Greeting: %s", rsp.Message)
			gctx.JSON(200, gin.H{
				"message": rsp.Message,
			})
		}
		cancel()
	})
	r.Run(listen)
 }