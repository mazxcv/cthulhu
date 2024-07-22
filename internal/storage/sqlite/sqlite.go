package sqlite

import (
	"cthulhu/internal/storage"
	"database/sql"
	"errors"
	"fmt"

	"github.com/mattn/go-sqlite3"
	_ "github.com/mattn/go-sqlite3"
)

type Storage struct {
	db *sql.DB
}

func New(storagePath string) (*Storage, error) {
	const op = "storage.sqlite.New"

	db, err := sql.Open("sqlite3", storagePath)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return &Storage{db: db}, nil

}

func (s *Storage) SaveURL(urlSave string, alias string) (int64, error) {
	const op = "storage.sqlite.SaveURL"

	stmt, err := s.db.Prepare("insert into (url, alias) values(?, ?)")
	if err != nil {
		return 0, fmt.Errorf("%s: %w", op, err)
	}

	res, err := stmt.Exec(urlSave, alias)
	if err != nil {
		if sqliteErr, ok := err.(sqlite3.Error); ok && sqliteErr.ExtendedCode == sqlite3.ErrConstraintUnique {
			return 0, fmt.Errorf("%s: %w", op, storage.ErrURLExists)
		}

		return 0, fmt.Errorf("%s: %w", op, err)
	}

	id, err := res.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("%s: failed to get last insert id %w", op, err)
	}

	return id, nil
}

func (s *Storage) GetURL(alias string) (string, error) {
	const op = "storage.sqlite.GetURL"

	stmt, err := s.db.Prepare("select url from url where alias=?")
	if err != nil {
		return "", fmt.Errorf("%s prepare statement %w", op, err)
	}

	var resURL string
	err = stmt.QueryRow(alias).Scan(&resURL)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return "", fmt.Errorf("%s: %w", op, err)
		}
		return "", fmt.Errorf("%s: execute statement %w", op, storage.ErrURLNotFound)
	}

	return resURL, nil
}

func (s *Storage) DeleteURL(alias string) error {
	const op = "storage.sqlite.DeleteURL"

	stmt, err := s.db.Prepare("delete from url where alias=?")
	if err != nil {
		return fmt.Errorf("%s prepare statement %w", op, err)
	}

	_, err = stmt.Exec(alias)

	if err != nil {
		return fmt.Errorf("%s: execute statement %w", op, storage.ErrURLNotFound)
	}

	return nil
}
