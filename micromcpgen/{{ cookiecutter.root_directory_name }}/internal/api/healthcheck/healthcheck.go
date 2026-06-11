package healthcheck

import (
	"time"
)

type Healthcheck struct {
	Host     string    `json:"host"`
	Datetime time.Time `json:"datetime"`
}

func New() *Healthcheck {
	return &Healthcheck{}
}
