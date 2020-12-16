package repository

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/lytics/logrus"
	"github.com/thedevsaddam/clean/domain"
)

// NewMysqlUserRepository ...
func NewMysqlUserRepository(db *sql.DB) domain.UserRepository {
	return &mysqlUserRepository{
		db: db,
	}
}

// mysqlUserRepository ...
type mysqlUserRepository struct {
	db *sql.DB
}

// Store ...
func (m *mysqlUserRepository) Store(ctx context.Context, user *domain.User) (uint, error) {
	query := `INSERT users SET type=?, username=?, password=?, created_at=?, updated_at=?`
	stmt, err := m.db.PrepareContext(ctx, query)
	if err != nil {
		return 0, err
	}

	t := time.Now()
	if user.CreatedAt.Unix() == 0 {
		user.CreatedAt = t
	}

	if user.UpdatedAt == nil {
		user.UpdatedAt = &t
	}

	// user.Password = // perform hash

	res, err := stmt.ExecContext(ctx, user.Type, user.Username, user.Password, user.CreatedAt, user.UpdatedAt)
	if err != nil {
		return 0, err
	}

	lastID, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	user.ID = uint(lastID)
	return user.ID, nil
}

// Fetch ...
func (m *mysqlUserRepository) Fetch(ctx context.Context, ctr *domain.UserCriteria) ([]*domain.User, error) {
	query := `SELECT id,username,type,created_at,updated_at FROM users WHERE ORDER BY id desc LIMIT ? ?`
	rows, err := m.db.QueryContext(ctx, query, ctr.Limit, ctr.Offset)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	users := make([]*domain.User, 0)
	for rows.Next() {
		u := new(domain.User)
		err = rows.Scan(
			&u.ID,
			&u.Username,
			&u.Type,
			&u.CreatedAt,
			&u.UpdatedAt,
		)
		if err != nil {
			logrus.Error(err)
			return nil, err
		}
		users = append(users, u)
	}
	return users, nil
}

// GetByID ...
func (m *mysqlUserRepository) GetByID(ctx context.Context, id uint) (*domain.User, error) {
	query := `SELECT id,username,type,created_at,updated_at FROM users WHERE id = ?`
	user := new(domain.User)
	err := m.db.QueryRowContext(ctx, query, id).Scan(
		&user.ID,
		&user.Username,
		&user.Type,
		&user.CreatedAt,
		&user.UpdatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return user, nil
}

// Delete ...
func (m *mysqlUserRepository) Delete(ctx context.Context, id uint) error {
	query := `DELETE FROM users WHERE id = ?`

	stmt, err := m.db.PrepareContext(ctx, query)
	if err != nil {
		return err
	}

	res, err := stmt.ExecContext(ctx, id)
	if err != nil {
		return err
	}

	rowsAfected, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAfected != 1 {
		err = fmt.Errorf("Weird  Behaviour. Total Affected: %d", rowsAfected)
		return err
	}

	return nil
}
