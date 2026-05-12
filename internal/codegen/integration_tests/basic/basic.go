// DO NOT EDIT! GENERATED CODE!
package basic

import (
	"context"
	"database/sql"
)

type Basic struct {
	Id        int64  `json:"id"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type BasicStore struct {
	db *sql.DB
}

func NewBasicStore(db *sql.DB) *BasicStore {
	return &BasicStore{db: db}
}

func (s *BasicStore) GetById(ctx context.Context, id int64) (*Basic, error) {
	query := `
		select	id,
				created_at,
				updated_at
		from	basic
		where	id = ?;
	`

	var item Basic
	err := s.db.QueryRowContext(ctx, query, id).Scan(
		&item.Id,
		&item.CreatedAt,
		&item.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}

	return &item, nil
}

func (s *BasicStore) UpdateById(ctx context.Context, item *Basic) error {
	query := `
		update	basic
		set		updated_at = datetime()
		where	id = ?;
	`

	_, err := s.db.ExecContext(
		ctx,
		query,
		item.Id,
	)

	if err != nil {
		return err
	}

	return nil
}

func (s *BasicStore) Insert(ctx context.Context, item *Basic) (int64, error) {
	query := `
		insert into basic(
			created_at,
			updated_at
		)
		values (datetime(),
			datetime()
		);
	`

	result, err := s.db.ExecContext(
		ctx,
		query,
	)

	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (s *BasicStore) DeleteById(ctx context.Context, id int64) error {
	query := `
		delete from basic
		where id = ?;
	`

	_, err := s.db.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}

	return nil
}

func (s *BasicStore) GetByIdTx(ctx context.Context, tx *sql.Tx, id int64) (*Basic, error) {
	query := `
		select	id,
				created_at,
				updated_at
		from	basic
		where	id = ?;
	`

	var item Basic
	err := tx.QueryRowContext(ctx, query, id).Scan(
		&item.Id,
		&item.CreatedAt,
		&item.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}

	return &item, nil
}

func (s *BasicStore) UpdateByIdTx(ctx context.Context, tx *sql.Tx, item *Basic) error {
	query := `
		update	basic
		set		updated_at = datetime()
		where	id = ?;
	`

	_, err := tx.ExecContext(
		ctx,
		query,
		item.Id,
	)

	if err != nil {
		return err
	}

	return nil
}

func (s *BasicStore) InsertTx(ctx context.Context, tx *sql.Tx, item *Basic) (int64, error) {
	query := `
		insert into basic(
			created_at,
			updated_at
		)
		values (datetime(),
			datetime()
		);
	`

	result, err := tx.ExecContext(
		ctx,
		query,
	)

	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (s *BasicStore) DeleteByIdTx(ctx context.Context, tx *sql.Tx, id int64) error {
	query := `
		delete from basic
		where id = ?;
	`

	_, err := tx.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}

	return nil
}

func (s *BasicStore) GetMany(ctx context.Context, ids []int64) ([]*Basic, error) {
	query := `
		select	id,
				created_at,
				updated_at
		from	basic
		where	id in ?;
	`

	stmt, err := s.db.PrepareContext(ctx, query)
	if err != nil {
		return nil, err
	}

	var results []*Basic
	for _, id := range ids {
		var result Basic
		_, err = stmt.ExecContext(ctx, id)
		if err != nil {
			return nil, err
		}

		results = append(results, &result)
	}

	return results, nil
}

func (s *BasicStore) GetManyTx(ctx context.Context, tx *sql.Tx, ids []int64) ([]*Basic, error) {
	query := `
		select	id,
				created_at,
				updated_at
		from	basic
		where	id in ?;
	`

	stmt, err := tx.PrepareContext(ctx, query)
	if err != nil {
		return nil, err
	}

	var results []*Basic
	for _, id := range ids {
		var result Basic
		_, err = stmt.ExecContext(ctx, id)
		if err != nil {
			return nil, err
		}

		results = append(results, &result)
	}

	return results, nil
}

func (s *BasicStore) DeleteManyTx(ctx context.Context, tx *sql.Tx, ids []int64) error {
	query := `
		delete from basic
		where id = ?;
	`

	stmt, err := tx.PrepareContext(ctx, query)
	if err != nil {
		return err
	}

	for _, id := range ids {
		_, err = stmt.ExecContext(ctx, id)
		if err != nil {
			return err
		}
	}

	return nil
}

func (s *BasicStore) UpdateManyTx(ctx context.Context, tx *sql.Tx, items []*Basic) ([]int64, error) {
	query := `
		update	basic
		set		updated_at = datetime()
		where id = ?;
	`

	stmt, err := tx.PrepareContext(ctx, query)
	if err != nil {
		return nil, err
	}

	var results []int64
	for range items {
		result, err := stmt.ExecContext(
			ctx,
		)
		if err != nil {
			return nil, err
		}

		id, err := result.LastInsertId()
		if err != nil {
			return nil, err
		}

		results = append(results, id)
	}

	return results, nil
}

func (s *BasicStore) InsertManyTx(ctx context.Context, tx *sql.Tx, items []*Basic) ([]int64, error) {
	query := `
		insert into basic (
			created_at,
			updated_at
		) values (datetime(),
			datetime()
		);
	`

	stmt, err := tx.PrepareContext(ctx, query)
	if err != nil {
		return nil, err
	}

	var results []int64
	for range items {
		result, err := stmt.ExecContext(
			ctx,
		)
		if err != nil {
			return nil, err
		}

		id, err := result.LastInsertId()
		if err != nil {
			return nil, err
		}

		results = append(results, id)
	}

	return results, nil
}
