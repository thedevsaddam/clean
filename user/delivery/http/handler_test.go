package http

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/bxcodec/faker"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/thedevsaddam/clean/domain"
	"github.com/thedevsaddam/clean/domain/mocks"
)

func TestFetchUsers(t *testing.T) {
	var mockUser domain.User
	err := faker.FakeData(&mockUser)
	assert.NoError(t, err)
	mockUCase := new(mocks.UserUsecase)
	mockListUser := make([]*domain.User, 0)
	mockListUser = append(mockListUser, &mockUser)

	mockUCase.On("Fetch", context.TODO(), &domain.UserCriteria{}).Return(mockListUser, nil)

	req := httptest.NewRequest(http.MethodGet, "/users", strings.NewReader(""))

	rec := httptest.NewRecorder()
	handler := UserHandler{
		UserUsecase: mockUCase,
	}

	handler.FetchUsers(rec, req)
	assert.Equal(t, http.StatusOK, rec.Code)
	mockUCase.AssertExpectations(t)
}

func TestStore(t *testing.T) {
	var mockUser domain.User
	err := faker.FakeData(&mockUser)
	assert.NoError(t, err)
	mockUCase := new(mocks.UserUsecase)

	mockUCase.On("Store", mock.Anything, mock.AnythingOfType("*domain.User")).Return(uint(1), nil)

	jUser, err := json.Marshal(mockUser)
	assert.NoError(t, err)

	req := httptest.NewRequest(http.MethodPost, "/users", strings.NewReader(string(jUser)))

	rec := httptest.NewRecorder()
	handler := UserHandler{
		UserUsecase: mockUCase,
	}

	handler.Store(rec, req)

	assert.Equal(t, http.StatusCreated, rec.Code)
	mockUCase.AssertExpectations(t)
}
