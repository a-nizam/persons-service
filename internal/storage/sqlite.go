package storage

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/a-nizam/persons-service/internal/domain/models"

	_ "github.com/mattn/go-sqlite3"
)

var (
	ErrUserExists   = errors.New("user already exists")
	ErrUserNotFound = errors.New("user id not found")
)

type Storage struct {
	db *sql.DB
}

func New(storagePath string) (*Storage, error) {
	db, err := sql.Open("sqlite3", storagePath)
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return &Storage{db: db}, nil
}

func (s *Storage) AddPerson(ctx context.Context, p *models.Person) (id int64, err error) {
	stmt, err := s.db.PrepareContext(ctx, "INSERT INTO info (name, birthdate) VALUES(?, ?)")
	if err != nil {
		return
	}
	res, err := stmt.ExecContext(ctx, p.Name, p.Birthdate.Format("2006-01-02"))
	if err != nil {
		return
	}
	id, err = res.LastInsertId()
	return
}

func (s *Storage) Stop() error {
	return s.db.Close()
}

func (s *Storage) GetPerson(ctx context.Context, id int64) (person *models.Person, err error) {
	row := s.db.QueryRowContext(ctx, "SELECT * FROM info WHERE id=?", id)
	person = new(models.Person)
	var birthdate string
	err = row.Scan(&person.ID, &person.Name, &birthdate)
	if err != nil {
		return
	}
	if person.Birthdate, err = time.Parse("2006-01-02", birthdate); err != nil {
		return
	}

	return
}

func (s *Storage) EditPerson(ctx context.Context, person *models.Person) (err error) {
	stmt, err := s.db.PrepareContext(ctx, "UPDATE info SET name=?, birthdate=? WHERE id=?")
	if err != nil {
		return
	}
	_, err = stmt.ExecContext(ctx, person.Name, person.Birthdate.Format("2006-01-02"), person.ID)
	return
}

func (s *Storage) RemovePerson(ctx context.Context, id int64) (err error) {
	stmt, err := s.db.PrepareContext(ctx, "DELETE FROM info WHERE id=?")
	if err != nil {
		return
	}
	_, err = stmt.ExecContext(ctx, id)
	return
}

func (s *Storage) GetList() (personList []models.Person, err error) {
	rows, err := s.db.Query("SELECT * FROM info")
	if err != nil {
		return
	}
	for rows.Next() {
		var person models.Person
		var birthdate string
		rows.Scan(&person.ID, &person.Name, &birthdate)
		person.Birthdate, err = time.Parse("2006-01-02", birthdate)
		if err != nil {
			return
		}
		personList = append(personList, person)
	}
	return
}
