package github

import (
	"fmt"
)

func Formatter(events []Event) error {
	if len(events) == 0 {
		return fmt.Errorf("no activity found")
	}
	data := ""
	for _, event := range events {
		switch event.Type {
		case "PushEvent":
			No_Count := len(event.Payload.Commits)
			data = fmt.Sprintf("pushed %d commits to %s", No_Count, event.Repo.Name)
		case "IssuesEvent":
			data = fmt.Sprintf("%s a new issue in %s", event.Payload.Action, event.Repo.Name)
		case "WatchEvent":
			data = fmt.Sprintf("starred %s ", event.Repo.Name)
		case "ForkEvent":
			data = fmt.Sprintf("starred %s ", event.Repo.Name)
		case "CreateEvent":
			data = fmt.Sprintf("created %s in %s ", event.Payload.RefType, event.Repo.Name)

		default:
			data = fmt.Sprintf("%s ")
		}
	}
}
