package main

import (
	"context"
	"fmt"
	"io"
	"log"

	"github.com/JhonatanPatrocinio/gRPC-GO/pb/pb"
	"google.golang.org/grpc"
)

func main() {
	connection, err := grpc.Dial("localhost:8081", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("could not connect to gRPC server: %v", err)
	}

	defer connection.Close()
	client := pb.NewUserServiceClient(connection)
	// AddUser(client)

	AddUserVerbose(client)
}

func AddUser(client pb.UserServiceClient) {
	user := &pb.User{
		Id:    "1",
		Name:  "Jhonatan",
		Email: "jhonatan.patrocinio@gmail.com",
	}

	res, err := client.AddUser(context.Background(), user)
	if err != nil {
		log.Fatalf("Error when calling AddUser: %v", err)
	}

	fmt.Println(res)
}

func AddUserVerbose(client pb.UserServiceClient) {
	req := &pb.User{
		Id:    "1",
		Name:  "Jhonatan",
		Email: "jhonatan.patrocinio@gmail.com",
	}

	respStream, err := client.AddUserVerbose(context.Background(), req)
	if err != nil {
		log.Fatalf("Error when calling AddUser: %v", err)
	}

	for {
		stream, err := respStream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Error when receiving stream: %v", err)
		}
		fmt.Println("Status: ", stream.Status)
		fmt.Println("User: ", stream.User)

	}
}
