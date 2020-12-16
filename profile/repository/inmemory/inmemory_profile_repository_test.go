package repository

import (
	"context"
	"testing"

	"github.com/thedevsaddam/clean/domain"
	"gotest.tools/assert"
	is "gotest.tools/assert/cmp"
)

func TestInMemoryProfileRepository_Store(t *testing.T) {
	testCases := []struct {
		expectedID uint
		profile    *domain.Profile
	}{
		{
			expectedID: 1,
			profile: &domain.Profile{
				UserID: 1,
				Name:   "Tom",
				Age:    33,
				Bio:    "Tom is a good cat",
			},
		},
		{
			expectedID: 2,
			profile: &domain.Profile{
				UserID: 2,
				Name:   "Jerry",
				Age:    30,
				Bio:    "Jerry is a good mouse",
			},
		},
		{
			expectedID: 3,
			profile: &domain.Profile{
				UserID: 3,
				Name:   "Mike",
				Age:    40,
				Bio:    "Mike is a good software engineer",
			},
		},
	}

	im := NewInMemoryProfileRepository()
	ctx := context.TODO()
	for _, tc := range testCases {
		id, err := im.Store(ctx, tc.profile)
		assert.NilError(t, err)
		assert.Equal(t, id, tc.expectedID)
	}
}

func TestInMemoryProfileRepository_GetByUserID(t *testing.T) {
	im := NewInMemoryProfileRepository()
	ctx := context.TODO()
	expectedProfile := domain.Profile{
		ID:     3,
		UserID: 3,
		Name:   "Mike",
		Age:    40,
		Bio:    "Mike is a good software engineer",
	}
	profile, err := im.GetByUserID(ctx, 3)
	assert.NilError(t, err)
	assert.DeepEqual(t, profile, &expectedProfile)
}

func TestInMemoryProfileRepository_Delete(t *testing.T) {
	im := NewInMemoryProfileRepository()
	ctx := context.TODO()
	err := im.Delete(ctx, 3)
	assert.NilError(t, err)
	user, err := im.GetByUserID(ctx, 3)
	assert.NilError(t, err)
	assert.Assert(t, is.Nil(user))
}
