package rest

import (
	"bytes"
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/thedevsaddam/clean/domain"
	"gotest.tools/assert"
)

type ClientMock struct {
}

var followers = make([]*domain.Follower, 0, 2)

func init() {
	followers = append(followers, &domain.Follower{ID: 1, Login: "tom"})
	followers = append(followers, &domain.Follower{ID: 2, Login: "jerry"})
}

func (c *ClientMock) Do(r *http.Request) (*http.Response, error) {
	bb, err := json.Marshal(followers)
	if err != nil {
		return nil, err
	}
	b := ioutil.NopCloser(bytes.NewReader(bb))
	return &http.Response{
		StatusCode: http.StatusOK,
		Body:       b,
	}, nil
}

func TestRestFollower_GetByUserUsername(t *testing.T) {
	fr := &restFollowerRepository{
		url:    "http://localhost",
		client: &ClientMock{},
	}
	ctx := context.Background()
	username := "thedevsaddam"
	got, err := fr.GetByUserUsername(ctx, username)
	assert.NilError(t, err)
	assert.DeepEqual(t, got, followers)
}
