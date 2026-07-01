package github

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func fetchUserActivity(username string) ([]Event, error) {
	url := fmt.Sprintf(
		"https://api.github.com/users/%s/events",
		username,
	)
	resp, err := http.Get(url)
	resp.Body.Close()
	if err != nil {
		return nil, fmt.Errorf("github user activity %w", err)
	}
	switch resp.StatusCode {
	case http.StatusOK:
	}
	var events []Event
	err := json.NewDecoder(resp.Body).Decode(&events)
	if err != nil {
		return err, nil
	}
	return events, nil
}
