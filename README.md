# GithubUserActivity2
# GitHub User Activity CLI

A simple command-line application written in Go that fetches and displays the recent public activity of any GitHub user using the GitHub REST API.

This project was built to practice:

- Working with REST APIs
- HTTP requests in Go
- JSON decoding
- Struct modeling
- Command-line flags
- Error handling
- Package organization

---

## Features

- Fetches a user's recent public GitHub activity
- Displays activity in a human-readable format
- Handles common API errors:
  - User not found
  - Rate limit exceeded
  - Network errors
- Uses only Go's standard library (no third-party packages)

---

## Project Structure

```
github-activity/
│
├── go.mod
│
├── cmd/
│   └── main.go
│
├── github/
│   ├── client.go
│   ├── event.go
│   └── formatter.go
│
└── README.md
```

---

## Installation

Clone the repository.

```bash
git clone https://github.com/<your-username>/GithubUserActivity2.git
```

Move into the project.

```bash
cd GithubUserActivity2
```

Run the application.

```bash
go run ./cmd -username=<github-username>
```

Example:

```bash
go run ./cmd -username=kamranahmedse
```

---

## Sample Output

```
created branch in joelyakubu269/GithubUserActivity2

opened a new issue in roadmapsh/project

starred octocat/Hello-World

pushed to joelyakubu269/webRecoding
```

---

## Supported Events

Currently the application formats:

- PushEvent
- CreateEvent
- IssuesEvent
- WatchEvent
- ForkEvent

Unknown events are displayed using a generic formatter.

---

## Technologies Used

- Go
- GitHub REST API
- JSON
- HTTP

---

## API Endpoint

```
GET https://api.github.com/users/{username}/events
```

---

## Error Handling

The application handles:

- Invalid usernames (404)
- GitHub API rate limiting (403)
- Network failures
- Empty activity

---

## Limitations

GitHub's current `/users/{username}/events` endpoint no longer returns commit counts for `PushEvent` payloads. As a result, this application displays push events without reporting the number of commits.

---

## What I Learned

While building this project I learned:

- How to model nested JSON using Go structs
- The difference between `json.Decoder` and `json.Unmarshal`
- Proper package organization in Go
- Building a CLI using the `flag` package
- Working with the `net/http` package
- Debugging API responses by comparing decoded structs with raw JSON
- Handling HTTP status codes gracefully

---

## Future Improvements

- Filter events by type
- Colorized terminal output
- GitHub authentication using Personal Access Tokens
- Better formatting using tables
- Fetch additional information for PushEvents using the Compare API