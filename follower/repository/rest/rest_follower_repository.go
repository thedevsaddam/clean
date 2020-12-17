package rest

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/thedevsaddam/clean/domain"
)

// cacheFollowers cache followers in-memoery
var cacheFollowers = make(map[string][]*domain.Follower, 0)

// HTTPClient interface
type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

type restFollowerRepository struct {
	url    string
	client HTTPClient
}

// NewRestFollowerRepository describe REST Followers repository; using gtihub as data source
func NewRestFollowerRepository(client HTTPClient) domain.FollowerRepository {
	return &restFollowerRepository{
		url:    "https://api.github.com",
		client: client,
	}
}

//GetByUserUsername fetch all followers by username from github
func (r *restFollowerRepository) GetByUserUsername(ctx context.Context, username string) ([]*domain.Follower, error) {
	if ff, exist := cacheFollowers[username]; exist {
		return ff, nil
	}

	endpoint := fmt.Sprintf("%s/users/%s/followers", r.url, username)
	req, err := http.NewRequest(http.MethodGet, endpoint, nil)
	if err != nil {
		return nil, fmt.Errorf("repository: %v", err)
	}
	req = req.WithContext(ctx)
	resp, err := r.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("repository: %v", err)
	}
	followers := make([]*domain.Follower, 0)
	if err := json.NewDecoder(resp.Body).Decode(&followers); err != nil {
		if err != nil {
			return nil, fmt.Errorf("repository: %v", err)
		}
	}
	cacheFollowers[username] = followers
	return followers, nil
}
