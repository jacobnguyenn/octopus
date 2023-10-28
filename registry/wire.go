//go:build wireinject
// +build wireinject

package registry

import (
	"github.com/google/wire"

	"ddd-sample/domain/service"
	"ddd-sample/interface/persistence/memory"
	"ddd-sample/interface/transport"
	"ddd-sample/interface/workflow/temporal"
	"ddd-sample/usecase"
)

func InitializeServer() transport.IServer {
	wire.Build(transport.NewServer, usecase.NewTicketUseCase, temporal.NewWorkflowRepository, service.NewTicketService,
		memory.NewTicketRepository)
	return transport.Server{}
}
