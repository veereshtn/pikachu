package svc

import (
	"context"
	"errors"
	"github.com/kutty-kumar/db_commons/model"
	"pikachu/pkg/domain"
	"pikachu/pkg/pb"
	"pikachu/pkg/repository"
)

type UserService struct {
	Repository repository.UserRepository
}

func userOperationResponseMapper(dto *pb.UserDto) *pb.UserOperationResponse {
	return &pb.UserOperationResponse{
		Response: dto,
	}
}

func (u *UserService) handleError(err error, base db_commons.Base,
	responseMapper func(dto *pb.UserDto) *pb.UserOperationResponse) (*pb.UserOperationResponse, error) {
	if err != nil {
		return nil, err
	}
	return responseMapper(base.ToDto().(*pb.UserDto)), err
}

func (u *UserService) getUser(dto pb.UserDto) *domain.User {
	user := domain.User{}
	user.FillProperties(dto)
	return &user
}

func (u *UserService) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.UserOperationResponse, error) {
	user := u.getUser(*req.Payload)
	createdUser, err := u.Repository.Create(user)
	return u.handleError(createdUser, err, userOperationResponseMapper)
}

func (u *UserService) UpdateUser(ctx context.Context, req *pb.UpdateUserRequest) (*pb.UserOperationResponse, error) {
	user := u.getUser(*req.Payload)
	updatedUser, err := u.Repository.Update(req.UserId, user)
	return u.handleError(updatedUser, err, userOperationResponseMapper)
}

func (u *UserService) GetUserByExternalId(ctx context.Context, req *pb.GetUserByExternalIdRequest) (*pb.UserOperationResponse, error) {
	user, err := u.Repository.FindByExternalId(req.UserId)
	return u.handleError(user, err, userOperationResponseMapper)
}

func (u *UserService) MultiGetUsersByExternalId(ctx context.Context, req *pb.MultiGetUsersByExternalIdRequest) (*pb.MultiGetUsersResponse, error) {
	if len(req.UserIds) > 0 {
		err, userSlice := u.Repository.MultiGetByExternalIds(req.UserIds)
		if err != nil {
			return nil, err
		}
		var userDtoSlice []*pb.UserDto
		for _, user := range userSlice {
			userDtoSlice = append(userDtoSlice, user.ToDto().(*pb.UserDto))
		}
		return &pb.MultiGetUsersResponse{Response: userDtoSlice}, nil
	}
	return nil, errors.New("invalid payload")
}
