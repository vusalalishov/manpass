package service

import (
	"github.com/vusalalishov/manpass/api"
	"github.com/vusalalishov/manpass/internal/model"
	"github.com/vusalalishov/manpass/internal/repository"
	"time"
)

type CredentialService interface {
	Save(*api.Credential) (*api.CredentialId, error)
	GetAll() (*api.CredentialsResponse, error)
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

func (cs *credentialService) GetAll() (*api.CredentialsResponse, error) {
	var all, err = cs.credRepository.GetAll()
	if err != nil {
		return nil, err
	}
	var creds = make([]*api.CredentialResponseItem, 0)
	for _, v := range *all {
		creds = append(creds, &api.CredentialResponseItem{
			Id: v.Id,
			Credential: &api.Credential{
				Title: v.Title,
				Login: v.Login,
				Password: v.Password,
			},
		})
	}
	return &api.CredentialsResponse{Credentials: creds}, nil
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
