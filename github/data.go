package github

type RepoIssueResult SingleIssueResult

type SingleIssueResult struct {
	Id            int
	NodeId        string `json:"node_id"`
	Url           string
	RepositoryUrl string `json:"repository_url"`
	LabelsUrl     string `json:"labels_url"`
	HtmlUrl       string `json:"html_url"`
	Number        int
	State         string
	Title         string
	Body          string
	User          struct {
		Login     string
		Id        int
		NodeId    string `json:"node_id"`
		AvatarUrl string `json:"avatar_url"`
		Url       string
		HtmlUrl   string `json:"html_url"`
		Type      string
	}
	Labels []struct {
		Id          int
		NodeId      string `json:"node_id"`
		Url         string
		Name        string
		Description string
		Color       string
		Default     bool
	}
	Locked    bool
	Comments  int
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}
