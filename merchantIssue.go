package domain

import (
	"encoding/json"
	"time"
)

type MerchantIssue struct {
	MerchantId  string    `json:"merchant_id"`
	Project     string    `json:"project"`
	Time        time.Time `json:"time"`
	JsonBodyStr string    `json:"json_body"`
	JsonBody
}

type JsonBody struct {
	Message      string `json:"message"`
	Subject      string `json:"subject"`
	Category     string `json:"category"`
	MerchantName string `json:"merchant_name"`
	Status       string `json:"status,omitempty"`
	JiraKey      string `json:"jira_key,omitempty"`
	Answer       string `json:"answer,omitempty"`
}

func (i *MerchantIssue) UnmarshalJsonBodyStrToJsonBody() error {

	err := json.Unmarshal([]byte(i.JsonBodyStr), &i.JsonBody)
	if err != nil {
		return err
	}

	return nil
}
