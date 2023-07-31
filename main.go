package main

import (
	"context"
	"fmt"
	"log"
	"time"

	pb "github.com/koetmongkhon/go-grpc-server-test/tutorial"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
)

func main() {
	conn, err := grpc.Dial("localhost:54321", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Println("Connect fail: ", err)
	}

	// Code removed for brevity

	client := pb.NewTutorialClient(conn)

	now := time.Now()

	log.Println(now)

	ctx, cancel := context.WithDeadline(context.Background(), now.Add(1*time.Second))
	defer cancel()

	// Note how we are calling the GetBookList method on the server
	// This is available to us through the auto-generated code
	bookList, err := client.SayHello(ctx, &pb.HelloRequest{Name: "Wasawat"})

	if err != nil {
		fmt.Println(err.Error())
		fmt.Println(context.DeadlineExceeded)
		if err == context.DeadlineExceeded {
			log.Println("Client cancelled")
			status.New(codes.Canceled, "Client cancelled, abandoning.")
		} else {
			log.Fatal("Fail to Say Hello: ", err.Error())
		}
	}

	log.Printf("book list: %v", bookList)
}
