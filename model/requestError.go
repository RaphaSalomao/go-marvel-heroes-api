package model

import (
	"fmt"
	"time"
)

type ErrorResponse struct {
	Timestamp time.Time `json:"timestamp"`
	Message   string    `json:"message"`
}

func (e ErrorResponse) Error() string {
	return fmt.Sprintf("timestamp: %d\nmessage: %s", e.Timestamp.Unix(), e.Message)
}
