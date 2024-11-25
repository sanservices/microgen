package pointerutils

import "time"

func ToStringPointer(value string) *string {
	if value == "" {
		return nil
	}

	return &value
}

func ToTimePointer(date time.Time) *time.Time {
	if date.IsZero() {
		return nil
	}

	return &date
}
