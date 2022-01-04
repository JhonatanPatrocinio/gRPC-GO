package services

import (
	"context"
	"fmt"

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
