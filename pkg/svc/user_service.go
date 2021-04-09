package svc

import (
	"context"
	"errors"
	"github.com/kutty-kumar/db_commons/model"
	"github.com/kutty-kumar/ho_oh/pkg/pikachu_v1"
	"pikachu/pkg/domain"
)

type UserService struct {
	db_commons.BaseSvc
	IdentityService IdentityService
}

func NewUserService(base db_commons.BaseSvc, identitySvc IdentityService) UserService {
	return UserService{base, identitySvc}
}

func userOperationResponseMapper(dto *pikachu_v1.UserDto) *pikachu_v1.UserOperationResponse {
	return &pikachu_v1.UserOperationResponse{
		Response: dto,
	}
}

func (u *UserService) handleError(err error, base db_commons.Base,
	responseMapper func(dto *pikachu_v1.UserDto) *pikachu_v1.UserOperationResponse) (*pikachu_v1.UserOperationResponse, error) {
	if err != nil {
		return nil, err
	}
	return responseMapper(base.ToDto().(*pikachu_v1.UserDto)), err
}

func (u *UserService) getUser(dto pikachu_v1.UserDto) *domain.User {
	user := domain.User{}
	user.FillProperties(dto)
	return &user
}

func (u *UserService) CreateUser(ctx context.Context, req *pikachu_v1.CreateUserRequest) (*pikachu_v1.UserOperationResponse, error) {
	user := u.getUser(*req.Payload)
	createdUser, err := u.Create(user)
	return u.handleError(createdUser, err, userOperationResponseMapper)
}

func (u *UserService) UpdateUser(ctx context.Context, req *pikachu_v1.UpdateUserRequest) (*pikachu_v1.UserOperationResponse, error) {
	user := u.getUser(*req.Payload)
	updatedUser, err := u.Update(req.UserId, user)
	return u.handleError(updatedUser, err, userOperationResponseMapper)
}

func (u *UserService) GetUserByExternalId(ctx context.Context, req *pikachu_v1.GetUserByExternalIdRequest) (*pikachu_v1.UserOperationResponse, error) {
	user, err := u.FindByExternalId(req.UserId)
	return u.handleError(user, err, userOperationResponseMapper)
}

func (u *UserService) MultiGetUsersByExternalId(ctx context.Context, req *pikachu_v1.MultiGetUsersByExternalIdRequest) (*pikachu_v1.MultiGetUsersResponse, error) {
	if len(req.UserIds) > 0 {
		err, userSlice := u.MultiGetByExternalId(req.UserIds)
		if err != nil {
			return nil, err
		}
		var userDtoSlice []*pikachu_v1.UserDto
		for _, user := range userSlice {
			userDtoSlice = append(userDtoSlice, user.ToDto().(*pikachu_v1.UserDto))
		}
		return &pikachu_v1.MultiGetUsersResponse{Response: userDtoSlice}, nil
	}
	return nil, errors.New("invalid payload")
}

func (u *UserService) CreateUserIdentity(ctx context.Context, req *pikachu_v1.CreateUserIdentityRequest) (*pikachu_v1.CreateUserIdentityResponse, error) {
	uIdentity := domain.Identity{}
	uIdentity = *interface{}(uIdentity.FillProperties(*req.Payload)).(*domain.Identity)
	err, user := u.FindByExternalId(req.UserId)
	if err != nil || user == nil {
		return nil, err
	}
	uIdentity.UserID = user.GetExternalId()
	err, identity := u.IdentityService.Create(&uIdentity)
	if err != nil {
		return nil, err
	}
	identityDto := identity.ToDto().(pikachu_v1.IdentityDto)
	return &pikachu_v1.CreateUserIdentityResponse{Response: &identityDto}, nil
}

func (u *UserService) GetUserIdentities(ctx context.Context, req *pikachu_v1.GetUserIdentitiesRequest) (*pikachu_v1.GetUserIdentitiesResponse, error) {
	return u.IdentityService.GetUserIdentities(ctx, req)
}

func (u *UserService) UpdateUserIdentity(ctx context.Context, req *pikachu_v1.UpdateUserIdentityRequest) (*pikachu_v1.UpdateUserIdentityResponse, error) {
	err, user := u.FindByExternalId(req.UserId)
	if err != nil || user == nil {
		return nil, err
	}
	return u.IdentityService.UpdateUserIdentity(ctx, req)
}
