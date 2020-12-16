package repository

import (
	"context"
	"testing"
	"time"

	assert "github.com/stretchr/testify/assert"
	"github.com/thedevsaddam/clean/domain"
	sqlmock "gopkg.in/DATA-DOG/go-sqlmock.v1"
)

func TestMysql_Store(t *testing.T) {
	now := time.Now()
	usr := &domain.User{
		Username:  "username",
		Password:  "password",
		Type:      "admin",
		CreatedAt: now,
		UpdatedAt: &now,
	}
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	query := `INSERT users SET type=\?, username=\?, password=\?, created_at=\?, updated_at=\?`
	prep := mock.ExpectPrepare(query)
	prep.ExpectExec().WithArgs(usr.Type, usr.Username, usr.Password, usr.CreatedAt, usr.UpdatedAt).WillReturnResult(sqlmock.NewResult(1, 1))

	u := NewMysqlUserRepository(db)
	id, err := u.Store(context.TODO(), usr)
	assert.NoError(t, err)
	assert.Equal(t, uint(1), id)
}

func TestMysql_Fetch(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	tm := time.Now()

	mockUsers := []domain.User{
		{ID: 1, Username: "username1", Password: "password", Type: "admin", CreatedAt: tm, UpdatedAt: &tm},
		{ID: 2, Username: "username2", Password: "password", Type: "manager", CreatedAt: tm, UpdatedAt: &tm},
	}

	rows := sqlmock.NewRows([]string{"id", "username", "type", "created_at", "updated_at"}).
		AddRow(mockUsers[0].ID, mockUsers[0].Username, mockUsers[0].Type, mockUsers[0].CreatedAt, mockUsers[0].UpdatedAt).
		AddRow(mockUsers[1].ID, mockUsers[1].Username, mockUsers[1].Type, mockUsers[1].CreatedAt, mockUsers[1].UpdatedAt)

	query := `SELECT id,username,type,created_at,updated_at FROM users WHERE ORDER BY id desc LIMIT ? ?`

	mock.ExpectQuery(query).WillReturnRows(rows)
	ur := NewMysqlUserRepository(db)
	ctr := &domain.UserCriteria{
		Limit:  20,
		Offset: 0,
	}
	list, err := ur.Fetch(context.TODO(), ctr)
	assert.NoError(t, err)
	assert.Len(t, list, 2)
}

func TestMysql_GetByID(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	tm := time.Now()
	rows := sqlmock.NewRows([]string{"id", "username", "type", "created_at", "updated_at"}).
		AddRow(1, "username", "admin", tm, &tm).
		AddRow(2, "username1", "admin1", tm, &tm)

	query := `SELECT id,username,type,created_at,updated_at FROM users WHERE id=?`

	mock.ExpectQuery(query).WithArgs(1).WillReturnRows(rows)
	u := NewMysqlUserRepository(db)
	id := uint(1)
	user, err := u.GetByID(context.TODO(), id)
	assert.NoError(t, err)
	assert.NotNil(t, user)
}

func TestMysql_Delete(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	query := `DELETE FROM users WHERE id=?`

	prep := mock.ExpectPrepare(query)
	prep.ExpectExec().WithArgs(1).WillReturnResult(sqlmock.NewResult(1, 1))

	a := NewMysqlUserRepository(db)
	id := uint(1)
	err = a.Delete(context.TODO(), id)
	assert.NoError(t, err)
}
