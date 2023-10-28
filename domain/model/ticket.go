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
	start    time.Time
	end      time.Time
	TicketId string
}

func NewTicket(id, content string, start time.Time, end time.Time) *Ticket {
	return &Ticket{
		Id:      id,
		Content: content,
		ActiveWindow: ActiveWindow{
			start: start,
			end:   end,
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
	return a.start
}

func (a *ActiveWindow) GetEnd() time.Time {
	return a.end
}

func (t *Ticket) GetActiveWindow() *ActiveWindow {
	return &t.ActiveWindow
}

func (t *Ticket) GetStart() time.Time {
	return t.ActiveWindow.start
}

func (t *Ticket) GetEnd() time.Time {
	return t.ActiveWindow.end
}
