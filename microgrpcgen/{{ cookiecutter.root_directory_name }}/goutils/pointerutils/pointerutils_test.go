package pointerutils

import (
	"strings"
	"testing"
	"time"
)

func TestToStringPointer(t *testing.T) {

	cc := []struct {
		Name  string
		Value string
	}{
		{
			Name:  "Convert to pointer",
			Value: "value",
		},
		{
			Name:  "Convert to nil with empty value",
			Value: "",
		},
	}

	for i := range cc {

		tc := cc[i]
		t.Run(tc.Name, func(t *testing.T) {
			t.Parallel()

			val := ToStringPointer(tc.Value)

			if val == nil && tc.Value != "" {
				t.Errorf("Unexpected nil result with %v", tc.Value)
			}

			if val != nil && strings.Compare(*val, tc.Value) != 0 {
				t.Errorf("Incorrect result %v with value %v", *val, tc.Value)
			}
		})
	}
}

func TestToTimePointer(t *testing.T) {
	testCases := []struct {
		Name  string
		Value time.Time
	}{
		{
			Name:  "Convert to pointer",
			Value: time.Now(),
		},
		{
			Name:  "zero value",
			Value: time.Time{},
		},
	}

	for i := range testCases {
		tc := testCases[i]

		t.Run(tc.Name, func(t *testing.T) {
			t.Parallel()

			val := ToTimePointer(tc.Value)

			if val == nil && !tc.Value.IsZero() {
				t.Errorf("Unexpected nil result with %v", tc.Value)
			}

			if val != nil && *val != tc.Value {
				t.Errorf("Incorrect result %v with value %v", *val, tc.Value)
			}
		})
	}
}
