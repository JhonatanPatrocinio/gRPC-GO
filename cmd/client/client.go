package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"time"

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
	// AddUserVerbose(client)
	// AddUsers(client)
	AddUserStreamBi(client)
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

func AddUsers(client pb.UserServiceClient) {
	reqs := []*pb.User{
		&pb.User{
			Id:    "1",
			Name:  "Jhonatan",
			Email: "jh@gmail.com",
		},
		&pb.User{
			Id:    "2",
			Name:  "Jhonatan2",
			Email: "jho@gmail.com",
		},
		&pb.User{
			Id:    "3",
			Name:  "Jhonatan3",
			Email: "jhon@gmail.com",
		},
	}

	stream, err := client.AddUsers(context.Background())

	if err != nil {
		log.Fatalf("Error when calling AddUsers: %v", err)
	}

	for _, req := range reqs {
		stream.Send(req)
		time.Sleep(time.Second * 2)
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("Error when receiving stream: %v", err)
	}
	fmt.Println(res)
}

func AddUserStreamBi(client pb.UserServiceClient) {
	stream, err := client.AddUserStreamBidi(context.Background())

	if err != nil {
		log.Fatalf("Error when calling AddUserStreamBidi: %v", err)
	}

	waitc := make(chan int)

	go func() {
		for i := 0; i < 5; i++ {
			fmt.Println("Sending: ", i)
			req := &pb.User{
				Id:    fmt.Sprintf("%d", i),
				Name:  fmt.Sprintf("Jhonatan%d", i),
				Email: fmt.Sprintf("jhonatan%d@gmail.com", i),
			}
			stream.Send(req)
			time.Sleep(time.Second * 2)
		}
		stream.CloseSend()
	}()

	go func() {
		for {
			res, err := stream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatalf("Error on received data %v", err)
			}

			fmt.Printf("Received user: %v with status: %v\n", res.GetUser().GetName(), res.GetStatus())
		}
		close(waitc)
	}()

	<-waitc
}
