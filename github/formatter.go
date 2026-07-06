package github

import (
	"fmt"
)

func Formatter(events []Event) (string, error) {
	if len(events) == 0 {
		return "", fmt.Errorf("no activity found")
	}
	data := ""
	for _, event := range events {
		switch event.Type {
		case "PushEvent":

			data += fmt.Sprintf("pushed commits to %s", event.Repo.Name)
			data += "\n"
		case "IssuesEvent":
			data += fmt.Sprintf("%s a new issue in %s", event.Payload.Action, event.Repo.Name)
			data += "\n"
		case "WatchEvent":
			data += fmt.Sprintf("starred %s ", event.Repo.Name)
			data += "\n"
		case "ForkEvent":
			data += fmt.Sprintf("starred %s ", event.Repo.Name)
			data += "\n"
		case "CreateEvent":
			data += fmt.Sprintf("created %s in %s  at %s", event.Payload.RefType, event.Repo.Name, event.CreatedAt)
			data += "\n"

		default:
			data = fmt.Sprintf("%s occured in % s", event.Type, event.Repo.Name)
		}
	}
	return data, nil
}
