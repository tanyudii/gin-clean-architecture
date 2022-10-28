package entities

import (
	"database/sql/driver"
	"encoding/json"
)

type EmailLogData map[string]interface{}

func (e *EmailLogData) Scan(src interface{}) error {
	return json.Unmarshal(src.([]byte), &e)
}

func (e EmailLogData) Value() (driver.Value, error) {
	val, err := json.Marshal(e)
	return string(val), err
}
