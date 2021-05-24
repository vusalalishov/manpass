// +build wireinject

package service

import "github.com/google/wire"
import "github.com/vusalalishov/manpass/api"

type CredentialService interface {
	Save(*api.Credential) (*api.CredentialId, error)
}

type credentialService struct {

}

func (cs *credentialService) Save(cred *api.Credential) (*api.CredentialId, error) {
	panic("I don't know what to do !!!")
}

func ProvideCredService() CredentialService {
	return &credentialService{}
}

func InjectCredService() CredentialService {
	panic(wire.Build(ProvideCredService))
}
