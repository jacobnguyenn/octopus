package repo

import "ddd-sample/domain/model"

type ITicketRepo interface {
	Create(*model.Ticket) (id string, err error)
	Get(id string) (*model.Ticket, error)
}
