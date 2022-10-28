package entities

type EmailTemplate struct {
}

func (e EmailTemplate) TableName() string {
	return "email_templates"
}
