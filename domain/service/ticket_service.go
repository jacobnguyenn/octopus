package service

import (
	"ddd-sample/domain/model"
	"ddd-sample/domain/repo"
)

type TicketService struct {
	repo repo.ITicketRepo
}

func NewTicketService(repo repo.ITicketRepo) *TicketService {
	return &TicketService{
		repo: repo,
	}
}

func (t *TicketService) InvalidActiveWindow(ticket *model.Ticket) bool {
	return ticket.GetStart().After(ticket.GetEnd())
}

func (t *TicketService) EmptyContent(ticket *model.Ticket) bool {
	return ticket.GetContent() == ""
}

func (t *TicketService) EmptyId(ticket *model.Ticket) bool {
	return ticket.GetId() == ""
}
