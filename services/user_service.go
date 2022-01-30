package services

import (
	"context"
	"fmt"
	"io"
	"log"
	"time"

	"github.com/JhonatanPatrocinio/gRPC-GO/pb/pb"
)

type UserService struct {
	pb.UnimplementedUserServiceServer
}

func (s *UserService) AddUser(ctx context.Context, req *pb.User) (*pb.User, error) {

	fmt.Println("AddUser function called with name: ", req.Name)

	return &pb.User{
		Id:    "123",
		Name:  req.GetName(),
		Email: req.GetEmail(),
	}, nil
}

func (*UserService) AddUserVerbose(req *pb.User, stream pb.UserService_AddUserVerboseServer) error {

	fmt.Println("AddUserVerbose function called with name: ", req.Name)

	stream.Send(&pb.UserResultStream{
		Status: "Starting",
		User:   &pb.User{},
	})

	time.Sleep(time.Second * 4)

	stream.Send(&pb.UserResultStream{
		Status: "User added",
		User: &pb.User{
			Id:    "1",
			Name:  req.GetName(),
			Email: req.GetEmail(),
		},
	})

	time.Sleep(time.Second * 4)

	return nil
}

func (*UserService) AddUsers(stream pb.UserService_AddUsersServer) error {
	fmt.Println("AddUsers function called")

	users := []*pb.User{}

	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&pb.Users{
				User: users,
			})
		}
		if err != nil {
			log.Fatalf("Error on receiving stream: %v", err)
		}

		fmt.Println("AddUsers function called with name: ", req.GetName())

		users = append(users, &pb.User{
			Id:    req.GetId(),
			Name:  req.GetName(),
			Email: req.GetEmail(),
		})

		time.Sleep(time.Second * 4)
	}
}

func (*UserService) AddUserStreamBidi(stream pb.UserService_AddUserStreamBidiServer) error {
	fmt.Println("AdduserStreamBidi function called")

	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			log.Fatalf("Error on receiving stream: %v", err)
		}

		fmt.Println("AdduserStreamBidi function called with name: ", req.GetName())

		err = stream.Send(&pb.UserResultStream{
			Status: "User added",
			User:   req,
		})

		if err != nil {
			log.Fatalf("Error on sending stream: %v", err)
		}

		time.Sleep(time.Second * 4)
	}
}
