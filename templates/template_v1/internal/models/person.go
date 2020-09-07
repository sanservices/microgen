package models

import "time"

type Person struct {
	Id          int64
	Name        string
	Age         int32
	DateCreated time.Time
}
