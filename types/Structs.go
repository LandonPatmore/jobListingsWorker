package types

type GithubJob struct {
	ID         string `json:"id"`
	Type       string `json:"type"`
	CreatedAt  string `json:"created_at"`
	Company    string `json:"company"`
	Location   string `json:"location"`
	Title      string `json:"title"`
	HowToApply string `json:"how_to_apply"`
}
