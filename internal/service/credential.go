package service

import (
	"github.com/vusalalishov/manpass/api"
	"github.com/vusalalishov/manpass/internal/model"
	"github.com/vusalalishov/manpass/internal/repository"
	"time"
)

type CredentialService interface {
	Save(*api.Credential) (*api.CredentialId, error)
}

type credentialService struct {
	credRepository repository.CredentialRepository
}

func (cs *credentialService) Save(cred *api.Credential) (*api.CredentialId, error) {
	var credModel = model.Credential{
		Title: cred.Title,
		Login: cred.Login,
		Password: cred.Password,
		UpdatedAt: time.Now(),
	}
	id, err := cs.credRepository.Save(&credModel)
	if err != nil {
		return nil, err
	}
	return &api.CredentialId{Id: id}, nil
}

func ProvideCredService(credentialRepository repository.CredentialRepository) CredentialService {
	return &credentialService{credentialRepository}
}

func InjectCredService() (CredentialService, error) {
	credRepo, err := repository.InjectCredRepository()
	if err != nil {
		return nil, err
	}
	return ProvideCredService(credRepo), nil
}
