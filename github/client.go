package github

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

func fetchUserActivity(username string) ([]Event, error) {
	url := fmt.Sprintf(
		"https://api.github.com/users/%s/events",
		username,
	)
	resp, err := http.Get(url)

	if err != nil {
		return nil, fmt.Errorf("github user activity %w", err)
	}
	resp.Body.Close()
	switch resp.StatusCode {
	case http.StatusOK:
	case http.StatusNotFound:
		return nil, errors.New("github user not found")
	case http.StatusForbidden:
		return nil, errors.New("Github Api rate limit exceeded")
	default:
		return nil, fmt.Errorf("GitHub returned %d", resp.StatusCode)
	}
	var events []Event
	err = json.NewDecoder(resp.Body).Decode(&events)
	if err != nil {
		return nil, err
	}
	return events, nil
}
