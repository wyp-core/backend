package app

import (
	repository "github.com/Abhyuday04/wyp/layers/repository"
	"github.com/Abhyuday04/wyp/layers/services"
	"github.com/Abhyuday04/wyp/layers/transport"
)

var Srv Server

type Server struct {
	Transport  transport.ITransport
	Service    services.IService
	Repository repository.IRepository
}
