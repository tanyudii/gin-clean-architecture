package entities

import (
	"database/sql/driver"
	"encoding/json"
)

type NotificationData map[string]interface{}

func (n *NotificationData) Scan(src interface{}) error {
	return json.Unmarshal(src.([]byte), &n)
}

func (n NotificationData) Value() (driver.Value, error) {
	val, err := json.Marshal(n)
	return string(val), err
}
