package model

import (
	"time"

	"ddd-sample/domain/model"
)

type Ticket struct {
	Id           string
	Content      string
	ActiveWindow ActiveWindow `gorm:"constraint:OnDelete:CASCADE;"`
}

type ActiveWindow struct {
	start    time.Time
	end      time.Time
	TicketId string
}

func ToPersistenceModelTicket(in *model.Ticket) (out *Ticket) {
	return &Ticket{
		Id:      in.GetId(),
		Content: in.GetContent(),
		ActiveWindow: ActiveWindow{
			start:    in.GetStart(),
			end:      in.GetEnd(),
			TicketId: in.GetId(),
		},
	}
}

func ToDomainModelTicket(in *Ticket) (out *model.Ticket) {
	return model.NewTicket(in.Id, in.Content, in.ActiveWindow.start, in.ActiveWindow.end)
}
