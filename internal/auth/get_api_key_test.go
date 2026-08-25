package auth

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestTetAPIKey(t *testing.T) {
	// Test: Good header
	headers := http.Header{}
	headers.Set("Authorization", "ApiKey 12345")

	key, err := GetAPIKey(headers)
	require.NoError(t, err)
	require.NotNil(t, key)
	assert.Equal(t, "12345", key)

	// Test: No Authorization Header
	headers = http.Header{}
	headers.Set("Content-Type", "application/json")

	key, err = GetAPIKey(headers)
	require.Error(t, err)
	assert.Equal(t, "", key)

	// Test: Malformed header
	headers = http.Header{}
	headers.Set("Authorization", "12345")

	key, err = GetAPIKey(headers)
	require.Error(t, err)
	assert.Equal(t, "", key)
}
