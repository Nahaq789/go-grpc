package controller

import (
	"context"
	"user_service/proto"
)

type UserController struct {
	proto.UnimplementedUserServiceServer
}

func (uc *UserController) CreateUser(ctx context.Context, req *proto.CreateUserRequest) (*proto.CreateUserResponse, error) {
	// Logic to create a user

	return nil, nil
}
