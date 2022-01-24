package services

import (
	"context"
	"fmt"
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
