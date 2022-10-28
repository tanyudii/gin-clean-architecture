package entities

type PasswordReset struct {
	UserSerial string `json:"userSerial"`
	UserEmail  string `json:"userEmail"`
	Token      string `json:"token"`
}
