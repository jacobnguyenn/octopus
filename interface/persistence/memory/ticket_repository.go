package memory

import (
	domainModel "ddd-sample/domain/model"
	"ddd-sample/domain/repo"
	"ddd-sample/interface/persistence/memory/model"

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
		Logger: logger.Default,
	})
	if err != nil {
		return nil, err
	}
	if err := db.AutoMigrate(
		&domainModel.Ticket{},
		&domainModel.ActiveWindow{},
	); err != nil {
		return nil, err
	}
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

func (t *ticketRepository) Create(ticket *domainModel.Ticket) (id string, err error) {
	if err = t.db.Create(model.ToPersistenceModelTicket(ticket)).Error; err != nil {
		return NonExistTicketId, err
	}
	return ticket.GetId(), err
}

func (t *ticketRepository) Get(id string) (resp *domainModel.Ticket, err error) {
	var persistenceModelTicket model.Ticket
	err = t.db.Where("id = ?", id).First(&persistenceModelTicket).Error
	resp = model.ToDomainModelTicket(&persistenceModelTicket)
	return resp, err
}
