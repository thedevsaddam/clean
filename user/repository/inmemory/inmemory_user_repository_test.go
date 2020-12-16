package repository

import (
	"context"
	"testing"

	"github.com/thedevsaddam/clean/domain"
	"gotest.tools/assert"
	is "gotest.tools/assert/cmp"
)

func TestInMemory_Store(t *testing.T) {
	ctx := context.TODO()
	im := NewInMemoryUserRepository()
	testCases := []struct {
		expectedID uint
		tag        string
		user       *domain.User
	}{
		{
			expectedID: 1,
			user: &domain.User{
				Username: "tom",
				Password: "pass",
			},
		},
		{
			expectedID: 2,
			user: &domain.User{
				Username: "jerry",
				Password: "pass",
			},
		},

		{
			expectedID: 3,
			user: &domain.User{
				Username: "mike",
				Password: "pass",
			},
		},
	}

	for _, tc := range testCases {
		id, err := im.Store(ctx, tc.user)
		assert.Equal(t, id, tc.expectedID)
		assert.NilError(t, err)
	}
}

func TestInMemory_Fetch(t *testing.T) {
	im := NewInMemoryUserRepository()
	ctx := context.TODO()
	ctr := &domain.UserCriteria{}
	users, err := im.Fetch(ctx, ctr)
	assert.NilError(t, err)
	assert.Assert(t, is.Len(users, 3))
}

func TestInMemory_GetByID(t *testing.T) {
	im := NewInMemoryUserRepository()
	ctx := context.TODO()
	user, err := im.GetByID(ctx, 3)
	expectedUser := &domain.User{
		ID:       3,
		Username: "mike",
		Password: "pass",
	}
	assert.NilError(t, err)
	assert.DeepEqual(t, user, expectedUser)
}

func TestInMemory_Delete(t *testing.T) {
	im := NewInMemoryUserRepository()
	ctx := context.TODO()
	err := im.Delete(ctx, 3)
	assert.NilError(t, err)

	user, err := im.GetByID(ctx, 3)
	assert.NilError(t, err)
	assert.Assert(t, is.Nil(user))
}
