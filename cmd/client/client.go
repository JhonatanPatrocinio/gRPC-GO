package main

import (
	"context"
	"fmt"
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
	AddUser(client)
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
