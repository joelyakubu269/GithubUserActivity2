package github

type Event struct {
	Type    string  `json:"type"`
	Repo    Repo    `json:"repo"`
	Payload Payload `json:"payload"`
}
type Repo struct {
	Name string `json:"name"`
}
type Payload struct {
	Action  string   `json:"action"`
	Ref     string   `json:"ref"`
	RefType string   `json:"ref_type"`
	Commits []Commit `json:"commits"`
}
type Commit struct {
	SHA string `json:"sha"`
}
