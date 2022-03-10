package settings

import (
	"context"
	"testing"
)

func TestGet(t *testing.T) {
	testCases := []struct {
		name          string
		filePath      string
		expectedError error
	}{
		{name: "correct file load", filePath: "../../settings.yml", expectedError: nil},
		{name: "file doesn't exists", filePath: "invalid.file", expectedError: ErrNoFile},
		{name: "incorrect format", filePath: "../../README.md", expectedError: ErrParsingFile},
	}

	ctx := context.Background()

	for i := range testCases {
		tc := testCases[i]

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			_, err := Get(ctx, tc.filePath)

			if err != tc.expectedError {
				t.Errorf("\nExpected: %v\nReceived: %v", tc.expectedError, err)
			}
		})
	}
}
