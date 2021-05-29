package server

import (
	"context"
	"github.com/vusalalishov/manpass/api"
	"github.com/vusalalishov/manpass/internal/service"
	"google.golang.org/grpc"
	"sync"
)

type grpcServer struct {
	server *grpc.Server
	credSvc service.CredentialService
}

func (s *grpcServer) Save(c context.Context, cred *api.Credential) (*api.CredentialId, error) {
	id, err := s.credSvc.Save(cred)
	if err != nil {
		return nil, err
	}
	return id, nil
}

func ProvideGrpcServer(cs service.CredentialService) *grpc.Server {
	var (
		srv *grpc.Server
		once sync.Once
	)
	once.Do(func() {
		srv = grpc.NewServer()
		api.RegisterCredentialStoreServer(srv, &grpcServer{
			server: srv,
			credSvc: cs,
		})
	})
	return srv
}

func InjectGrpcServer() (*grpc.Server, error) {
	srv, err := service.InjectCredService()
	if err != nil {
		return nil, err
	}
	return ProvideGrpcServer(srv), nil
}