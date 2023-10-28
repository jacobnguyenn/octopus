package memory

import (
	"ddd-sample/domain/model"
	"ddd-sample/domain/repo"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type ticketRepository struct {
	db *gorm.DB
}

var (
	_ repo.ITicketRepo = (*ticketRepository)(nil)
)

const (
	NonExistTicketId = "000000"
)

func newInmemDB() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		return nil, err
	}
	db.AutoMigrate(
		&model.Ticket{},
		&model.ActiveWindow{},
	)
	return db, nil
}

func NewTicketRepository() repo.ITicketRepo {
	db, err := newInmemDB()
	if err != nil {
		panic(err)
	}
	return &ticketRepository{
		db: db,
	}
}

func (t *ticketRepository) Create(ticket *model.Ticket) (id string, err error) {
	if err = t.db.Create(ticket).Error; err != nil {
		return NonExistTicketId, err
	}
	return ticket.GetId(), err
}

func (t *ticketRepository) Get(id string) (resp *model.Ticket, err error) {
	err = t.db.Where("id = ?", id).First(resp).Error
	return resp, err
}
