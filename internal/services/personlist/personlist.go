package personlist

import (
	"context"
	"log/slog"

	"github.com/a-nizam/persons-service/internal/domain/models"
)

type PersonList struct {
	log            *slog.Logger
	personProvider PersonProvider
}

type PersonProvider interface {
	AddPerson(ctx context.Context, person *models.Person) (int64, error)
	GetPerson(ctx context.Context, id int64) (*models.Person, error)
	EditPerson(ctx context.Context, person *models.Person) error
	RemovePerson(ctx context.Context, id int64) error
	GetList() ([]models.Person, error)
}

func New(log *slog.Logger, personProvider PersonProvider) *PersonList {
	return &PersonList{
		log:            log,
		personProvider: personProvider,
	}
}

func (pl *PersonList) AddPerson(ctx context.Context, person *models.Person) (id int64, err error) {
	pl.log.Info("Attempt to add person", slog.Any("person", person))
	id, err = pl.personProvider.AddPerson(ctx, person)
	if err != nil {
		pl.log.Error("Failed to add person", slog.Any("err", err))
	}
	return
}

func (pl *PersonList) GetPerson(ctx context.Context, id int64) (person *models.Person, err error) {
	pl.log.Info("Attempt to get person", slog.Int64("id", id))
	person, err = pl.personProvider.GetPerson(ctx, id)
	if err != nil {
		pl.log.Error("Failed to get person", slog.Any("err", err))
	}
	return
}

func (pl *PersonList) EditPerson(ctx context.Context, person *models.Person) (err error) {
	err = pl.personProvider.EditPerson(ctx, person)
	if err != nil {
		pl.log.Error("Failed to edit person", slog.Any("err", err))
	}
	return
}

func (pl *PersonList) RemovePerson(ctx context.Context, id int64) (err error) {
	err = pl.personProvider.RemovePerson(ctx, id)
	if err != nil {
		pl.log.Error("Failed to remove person", slog.Any("err", err))
	}
	return
}

func (pl *PersonList) GetList() (personList []models.Person, err error) {
	pl.log.Info("Attempt to get persons list")
	personList, err = pl.personProvider.GetList()
	if err != nil {
		pl.log.Error("Failed to get persons list", slog.Any("err", err))
	}
	return
}
