package main

import (
	"testing"
)

// GetAPIKey function retrieves the API key.
func GetAPIKey() string {
	// Simulate retrieving API key from somewhere
	return "example-api-key"
}

func TestGetAPIKey(t *testing.T) {
	// Call the function being tested
	apiKey := GetAPIKey()

	// Define the expected API key
	expectedKey := "example-api-key"

	// Check if the retrieved API key matches the expected key
	if apiKey != expectedKey {
		t.Errorf("Expected API key %s, but got %s", expectedKey, apiKey)
	}
}
