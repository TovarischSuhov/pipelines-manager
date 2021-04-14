package models

import "time"

type Pipeline struct {
	ID          int64          `json:"id"`
	ProjectID   int64          `json:"project_id"`
	Status      string         `json:"status"`
	Ref         string         `json:"ref"`
	SHA         string         `json:"sha"`
	BeforeSHA   string         `json:"before_sha"`
	Tag         bool           `json:"tag"`
	YamlErrors  *string        `json:"yaml_errors"`
	User        User           `json:"user"`
	CreatedAt   *time.Time     `json:"created_at"`
	UpdatedAt   *time.Time     `json:"updated_at"`
	StartedAt   *time.Time     `json:"started_at"`
	FinishedAt  *time.Time     `json:"finished_at"`
	CommittedAt *time.Time     `json:"committed_at"`
	Duration    *time.Duration `json:"duration"`
	Coverage    *string        `json:"coverage"`
	WebURL      string         `json:"web_url"`
}
