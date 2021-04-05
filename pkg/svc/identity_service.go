package svc

import (
	"context"
	"pikachu/pkg/domain"
	"pikachu/pkg/pb"
	"pikachu/pkg/repository"
)

type IdentityService struct {
	IdentityRepository repository.IdentityRepository
}

func (is *IdentityService) CreateUserIdentity(ctx context.Context, req *pb.CreateUserIdentityRequest) (*pb.CreateUserIdentityResponse, error) {
	uIdentity := domain.Identity{}
	uIdentity = *interface{}(uIdentity.FillProperties(*req.Payload)).(*domain.Identity)
	err, identity := is.IdentityRepository.CreateIdentity(req.UserId, &uIdentity)
	if err != nil {
		return nil, err
	}
	identityDto := identity.ToDto().(pb.IdentityDto)
	return &pb.CreateUserIdentityResponse{Response: &identityDto}, nil
}

func (is *IdentityService) GetUserIdentities(ctx context.Context, req *pb.GetUserIdentitiesRequest) (*pb.GetUserIdentitiesResponse, error) {
	var response []*pb.IdentityDto
	err, identities := is.IdentityRepository.ListIdentities(req.UserId)
	if err != nil {
		return nil, err
	}
	for _, identity := range identities {
		identityDto := identity.ToDto().(pb.IdentityDto)
		response = append(response, &identityDto)
	}
	return &pb.GetUserIdentitiesResponse{Response: response}, nil
}

func (is *IdentityService) UpdateUserIdentity(ctx context.Context, req *pb.UpdateUserIdentityRequest) (*pb.UpdateUserIdentityResponse, error) {
	uIdentity := domain.Identity{}
	uIdentity.FillProperties(*req.Payload)
	err, updatedIdentity := is.IdentityRepository.UpdateIdentity(req.UserId, req.UserIdentityId, &uIdentity)
	if err != nil {
		return nil, err
	}
	responseDto := updatedIdentity.ToDto().(pb.IdentityDto)
	return &pb.UpdateUserIdentityResponse{Response: &responseDto}, nil
}
