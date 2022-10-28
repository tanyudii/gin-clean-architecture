package entities

import "time"

type EmailLog struct {
	ID            int64         `json:"id"`
	CompanySerial string        `json:"companySerial"`
	ToEmail       string        `json:"toEmail"`
	ToName        *string       `json:"toName"`
	FromEmail     string        `json:"fromEmail"`
	FromName      *string       `json:"fromName"`
	Subject       string        `json:"subject"`
	Provider      string        `json:"provider"`
	Template      string        `json:"template"`
	TemplateCode  string        `json:"templateCode"`
	Data          *EmailLogData `json:"data"`
	AcceptedAt    *time.Time    `json:"acceptedAt"`
	DeliveredAt   *time.Time    `json:"deliveredAt"`
	Retry         int           `json:"retry"`
	MaxRetry      int           `json:"maxRetry"`
	LastError     string        `json:"lastError"`
}

func (e EmailLog) TableName() string {
	return "email_logs"
}
