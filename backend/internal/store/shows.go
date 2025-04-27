package store

import (
	"errors"

	"github.com/Gumilho/gumi-fetch/internal/types"
	"github.com/jmoiron/sqlx"
)

var (
	ErrDuplicateShow = errors.New("a show with that mal_id already exists")
)

type ShowStore struct {
	db *sqlx.DB
}

func NewShowStore(db *sqlx.DB) *ShowStore {
	return &ShowStore{db: db}
}

func (ss *ShowStore) List() ([]types.Show, error) {
	shows := []types.Show{}
	err := ss.db.Select(&shows, "SELECT * FROM shows")
	if err != nil {
		return nil, err
	}
	return shows, nil
}
func (ss *ShowStore) Create(show types.Show) error {
	query := `
		INSERT INTO shows (mal_id, title, source, source_id)
		VALUES (:mal_id, :title, :source, :source_id)
	`
	_, err := ss.db.NamedExec(query, show)

	if err != nil {
		switch {
		case err.Error() == `pq: duplicate key value violates unique constraint "show_mal_id_key"`:
			return ErrDuplicateShow
		default:
			return err
		}
	}
	return nil
}
func (ss *ShowStore) Delete(id int) error {
	_, err := ss.db.Exec("DELETE FROM shows WHERE mal_id=$1", id)
	if err != nil {
		return err
	}
	return nil
}
