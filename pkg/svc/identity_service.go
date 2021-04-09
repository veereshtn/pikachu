package svc

import (
	"context"
	"github.com/kutty-kumar/db_commons/model"
	"github.com/kutty-kumar/ho_oh/pkg/pikachu_v1"
	"pikachu/pkg/domain"
	"pikachu/pkg/repository"
)

type IdentityService struct {
	db_commons.BaseSvc
	IdentityRepository repository.IdentityRepository
}

func NewIdentityService(baseSvc db_commons.BaseSvc, identityRepository repository.IdentityRepository) IdentityService {
	return IdentityService{
		baseSvc,
		identityRepository,
	}
}

func (is *IdentityService) CreateUserIdentity(ctx context.Context, req *pikachu_v1.CreateUserIdentityRequest) (*pikachu_v1.CreateUserIdentityResponse, error) {
	uIdentity := domain.Identity{}
	uIdentity = *interface{}(uIdentity.FillProperties(*req.Payload)).(*domain.Identity)
	err, identity := is.IdentityRepository.CreateIdentity(req.UserId, &uIdentity)
	if err != nil {
		return nil, err
	}
	identityDto := identity.ToDto().(pikachu_v1.IdentityDto)
	return &pikachu_v1.CreateUserIdentityResponse{Response: &identityDto}, nil
}

func (is *IdentityService) GetUserIdentities(ctx context.Context, req *pikachu_v1.GetUserIdentitiesRequest) (*pikachu_v1.GetUserIdentitiesResponse, error) {
	var response []*pikachu_v1.IdentityDto
	err, identities := is.IdentityRepository.ListIdentities(req.UserId)
	if err != nil {
		return nil, err
	}
	for _, identity := range identities {
		identityDto := identity.ToDto().(pikachu_v1.IdentityDto)
		response = append(response, &identityDto)
	}
	return &pikachu_v1.GetUserIdentitiesResponse{Response: response}, nil
}

func (is *IdentityService) UpdateUserIdentity(ctx context.Context, req *pikachu_v1.UpdateUserIdentityRequest) (*pikachu_v1.UpdateUserIdentityResponse, error) {
	uIdentity := domain.Identity{}
	uIdentity.FillProperties(*req.Payload)
	err, updatedIdentity := is.IdentityRepository.UpdateIdentity(req.UserId, req.UserIdentityId, &uIdentity)
	if err != nil {
		return nil, err
	}
	responseDto := updatedIdentity.ToDto().(pikachu_v1.IdentityDto)
	return &pikachu_v1.UpdateUserIdentityResponse{Response: &responseDto}, nil
}
