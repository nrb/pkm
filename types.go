package main

import "time"

// PKMConfig holds configuration details for pkm
type PKMConfig struct {
	// WorkDir holds the directory for scripts and data generated from them.
	WorkDir string `json:"workDir"`
	// GitRoot holds the directory where local git clones are kept.
	GitRoot string `json:"gitRoot"`
}

// GitHubReview represents a GitHub pull request review request
type GitHubReview struct {
	Number     int              `json:"number"`
	Repository GitHubRepository `json:"repository"`
	Title      string           `json:"title"`
	UpdatedAt  time.Time        `json:"updatedAt"`
}

// GitHubRepository represents repository information
type GitHubRepository struct {
	Name          string `json:"name"`
	NameWithOwner string `json:"nameWithOwner"`
}

// JiraIssue represents a Jira issue
type JiraIssue struct {
	Key    string     `json:"key"`
	Fields JiraFields `json:"fields"`
}

// JiraFields contains all fields of a Jira issue
type JiraFields struct {
	Summary     string          `json:"summary"`
	Description string          `json:"description"`
	Labels      []string        `json:"labels"`
	Resolution  JiraNamedObject `json:"resolution"`
	IssueType   JiraIssueType   `json:"issueType"`
	Assignee    JiraUser        `json:"assignee"`
	Priority    JiraNamedObject `json:"priority"`
	Reporter    JiraUser        `json:"reporter"`
	Watches     JiraWatches     `json:"watches"`
	Status      JiraNamedObject `json:"status"`
	Components  []JiraComponent `json:"components"`
	FixVersions []JiraVersion   `json:"fixVersions"`
	Versions    []JiraVersion   `json:"versions"`
	Comment     JiraComments    `json:"comment"`
	Subtasks    []any           `json:"Subtasks"`
	IssueLinks  []JiraIssueLink `json:"issueLinks"`
	Created     string          `json:"created"`
	Updated     string          `json:"updated"`
}

// JiraNamedObject represents simple Jira objects with just a name
type JiraNamedObject struct {
	Name string `json:"name"`
}

// JiraIssueType represents Jira issue type information
type JiraIssueType struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Subtask bool   `json:"subtask"`
}

// JiraUser represents a Jira user
type JiraUser struct {
	DisplayName string `json:"displayName"`
}

// JiraWatches represents watch information for a Jira issue
type JiraWatches struct {
	IsWatching bool `json:"isWatching"`
	WatchCount int  `json:"watchCount"`
}

// JiraComponent represents a Jira component
type JiraComponent struct {
	Name string `json:"name"`
}

// JiraVersion represents a Jira version
type JiraVersion struct {
	Name string `json:"name"`
}

// JiraComments represents comment information
type JiraComments struct {
	Comments any `json:"comments"`
	Total    int `json:"total"`
}

// JiraIssueLink represents a link between Jira issues
type JiraIssueLink struct {
	ID           string            `json:"id"`
	Type         JiraIssueLinkType `json:"type"`
	InwardIssue  *JiraLinkedIssue  `json:"inwardIssue,omitempty"`
	OutwardIssue *JiraLinkedIssue  `json:"outwardIssue,omitempty"`
}

// JiraIssueLinkType represents the type of link between issues
type JiraIssueLinkType struct {
	Name    string `json:"name"`
	Inward  string `json:"inward"`
	Outward string `json:"outward"`
}

// JiraLinkedIssue represents a linked Jira issue (simplified)
type JiraLinkedIssue struct {
	Key    string     `json:"key"`
	Fields JiraFields `json:"fields"`
}

// GitCommit represents a git commit made today
type GitCommit struct {
	AbbreviatedCommit string    `json:"abbreviated_commit"`
	Branch            string    `json:"branch"`
	Author            GitPerson `json:"author"`
	Committer         GitPerson `json:"committer"`
	Subject           string    `json:"subject"`
}

// GitPerson represents a person (author or committer) in git
type GitPerson struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Date  string `json:"date"`
}
