package controller

import (
	"context"
	"user_service/model"
	"user_service/proto"
	"user_service/repository"
)

type UserController struct {
	r repository.UserRepository
	proto.UnimplementedUserServiceServer
}

func (uc *UserController) CreateUser(ctx context.Context, req *proto.CreateUserRequest) (*proto.CreateUserResponse, error) {
	// Logic to create a user

	user := model.NewUser(req.Name, req.Email)
	err := uc.r.CreateUser(user)
	if err != nil {
		return nil, err
	}

	res := &proto.CreateUserResponse{
		Message: "User created successfully",
	}
	return res, nil
}
