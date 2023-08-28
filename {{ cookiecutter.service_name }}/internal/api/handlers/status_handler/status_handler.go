package statushandler

import (
	"time"
)

type StatusHandler struct{}

type Healthcheck struct {
	Host     string    `json:"host"`
	Datetime time.Time `json:"datetime"`
}

func New() *StatusHandler {
	return &StatusHandler{}
}
