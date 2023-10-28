package model

import "time"

type Ticket struct {
	id           string
	content      string
	activeWindow ActiveWindow `gorm:"constraint:OnDelete:CASCADE;"`
}

type ActiveWindow struct {
	start    time.Time
	end      time.Time
	TicketId string
}

func NewTicket(id, content string, start time.Time, end time.Time) *Ticket {
	return &Ticket{
		id:      id,
		content: content,
		activeWindow: ActiveWindow{
			start: start,
			end:   end,
		},
	}
}

func (t *Ticket) GetId() string {
	return t.id
}

func (t *Ticket) GetContent() string {
	return t.content
}

func (a *ActiveWindow) GetStart() time.Time {
	return a.start
}

func (a *ActiveWindow) GetEnd() time.Time {
	return a.end
}

func (t *Ticket) GetActiveWindow() *ActiveWindow {
	return &t.activeWindow
}

func (t *Ticket) GetStart() time.Time {
	return t.activeWindow.start
}

func (t *Ticket) GetEnd() time.Time {
	return t.activeWindow.end
}
