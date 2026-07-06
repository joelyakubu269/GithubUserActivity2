package github

type Event struct {
	Type      string  `json:"type"`
	Repo      Repo    `json:"repo"`
	Payload   Payload `json:"payload"`
	CreatedAt string  `json:"created_at"`
}
type Repo struct {
	Name string `json:"name"`
}
type Payload struct {
	Action    string `json:"action"`
	Ref       string `json:"ref"`
	RefType   string `json:"ref_type"`
	Size      int    `json:"size"`
	CreatedAt string `json:"created_at"`
}
