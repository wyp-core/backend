package app

import (
	repositoryuser "github.com/Abhyuday04/wyp/layers/repository/users"
	"github.com/Abhyuday04/wyp/layers/services"
	"github.com/Abhyuday04/wyp/layers/transport"
)

var Srv Server

type Server struct {
	Transport      transport.ITransport
	Service        services.IService
	RepositoryUser repositoryuser.IUserRepository
}
