package repository

import "pikachu/pkg/domain"

type IdentityRepository interface {
	CreateIdentity(userId string, identity *domain.Identity) (error, *domain.Identity)
	UpdateIdentity(userId string, identityId string, identity *domain.Identity) (error, *domain.Identity)
	ListIdentities(userId string) (error, []domain.Identity)
}
