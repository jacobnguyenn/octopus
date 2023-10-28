package model

import (
	"time"
)

type Ticket struct {
	Id           string
	Content      string
	ActiveWindow ActiveWindow
}

type ActiveWindow struct {
	Start    time.Time
	End      time.Time
	TicketId string
}

func NewTicket(id, content string, start time.Time, end time.Time) *Ticket {
	return &Ticket{
		Id:      id,
		Content: content,
		ActiveWindow: ActiveWindow{
			Start: start,
			End:   end,
		},
	}
}

func (t *Ticket) GetId() string {
	return t.Id
}

func (t *Ticket) GetContent() string {
	return t.Content
}

func (a *ActiveWindow) GetStart() time.Time {
	return a.Start
}

func (a *ActiveWindow) GetEnd() time.Time {
	return a.End
}

func (t *Ticket) GetActiveWindow() *ActiveWindow {
	return &t.ActiveWindow
}

func (t *Ticket) GetStart() time.Time {
	return t.ActiveWindow.Start
}

func (t *Ticket) GetEnd() time.Time {
	return t.ActiveWindow.End
}
