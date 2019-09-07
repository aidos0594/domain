package domain

import "time"

type MerchantIssueJira struct {
	JiraKey    string    `json:"jira_key"`
	MerchantId string    `json:"merchant_id"`
	Project    string    `json:"project"`
	Time       time.Time `json:"time"`
}
