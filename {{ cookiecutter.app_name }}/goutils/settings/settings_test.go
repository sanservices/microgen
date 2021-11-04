package settings

import (
	"context"
	"os"
	"testing"
)

func TestGetEnv(t *testing.T) {

	envValue := "myTestValue"
	err := os.Setenv("test_env", envValue)
	if err != nil {
		t.Error(err)
	}

	testCases := []struct {
		name     string
		key      string
		value    string
		fallback string
	}{
		{name: "correct value", key: "test_env", value: envValue},
		{name: "fallback value", key: "fallback", fallback: "myFallback"},
	}

	for i := range testCases {
		tc := testCases[i]

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			val := GetEnv(tc.key, tc.fallback)

			if val != tc.value && val != tc.fallback {
				t.Errorf("Didn't get the correct value from %s", tc.key)
			}
		})
	}
}

func TestGet(t *testing.T) {
	testCases := []struct {
		name          string
		filePath      string
		expectedError error
	}{
		{name: "correct file load", filePath: "settings.yml", expectedError: nil},
		{name: "file doesn't exists", filePath: "invalid.file", expectedError: ErrNoFile},
		{name: "incorrect format", filePath: "test_file.yml", expectedError: ErrParsingFile},
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

func TestNew(t *testing.T) {

	ctx := context.Background()
	_, err := New(ctx)
	if err != nil {
		t.Error(err)
	}
}
