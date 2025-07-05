package repository

import (
	"context"
	"fmt"
	"organization_service/model"
	"organization_service/proto"
)

type UserServiceClient struct {
	client proto.UserServiceClient
}

func NewUserServiceClient(client proto.UserServiceClient) *UserServiceClient {
	return &UserServiceClient{
		client: client,
	}
}

func (r *UserServiceClient) CallUserServiceCreateUser(u model.User) error {
	req := &proto.CreateUserRequest{
		Name:  u.Name,
		Email: u.Email,
	}
	res, err := r.client.CreateUser(context.Background(), req)
	if err != nil {
		return err
	}
	fmt.Printf("User created with ID: %s\n", res.Message)
	return nil
}
