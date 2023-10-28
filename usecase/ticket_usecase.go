package usecase

import (
	"context"
	"errors"
	"time"

	"ddd-sample/domain/model"
	"ddd-sample/domain/repo"
	"ddd-sample/domain/service"

	"github.com/google/uuid"
)

type (
	ITicketUsecase interface {
		Create(content string, start time.Time, end time.Time) (id string, err error)
		Get(id string) (resp Ticket, err error)
	}
	ActiveWindow struct {
		Start time.Time
		End   time.Time
	}
	Ticket struct {
		Id           string
		Content      string
		ActiveWindow ActiveWindow
	}
	ticketUseCase struct {
		workflowRepo  repo.IWorkflowRepo
		ticketRepo    repo.ITicketRepo
		ticketService service.TicketService
	}
)

var (
	_                  ITicketUsecase = (*ticketUseCase)(nil)
	ErrInValidArgument                = errors.New("invalid argument")
)

func NewTicketUseCase(workflow repo.IWorkflowRepo, ticketService service.TicketService, ticketRepo repo.ITicketRepo) ITicketUsecase {
	return &ticketUseCase{
		workflowRepo:  workflow,
		ticketRepo:    ticketRepo,
		ticketService: ticketService,
	}
}

func (t *ticketUseCase) Create(content string, start time.Time, end time.Time) (id string, err error) {
	ctx := context.TODO()
	ticketId := uuid.New().String()
	ticket := model.NewTicket(ticketId, content, start, end)
	if t.ticketService.EmptyContent(ticket) || t.ticketService.InvalidActiveWindow(ticket) {
		return "", ErrInValidArgument
	}
	id, err = t.ticketRepo.Create(ticket)
	if err != nil {
		return "", err
	}
	return id, t.workflowRepo.Start(ctx, id)
}

func (t *ticketUseCase) Get(id string) (Ticket, error) {
	resp, err := t.ticketRepo.Get(id)
	if err != nil {
		return Ticket{}, err
	}
	return *toUseCaseTicket(resp), nil
}

func toUseCaseTicket(in *model.Ticket) (out *Ticket) {
	return &Ticket{
		Id:           in.GetId(),
		Content:      in.GetContent(),
		ActiveWindow: *toUseCaseActiveWindow(in.GetActiveWindow()),
	}
}

func toUseCaseActiveWindow(in *model.ActiveWindow) (out *ActiveWindow) {
	return &ActiveWindow{
		Start: in.GetStart(),
		End:   in.GetEnd(),
	}
}
