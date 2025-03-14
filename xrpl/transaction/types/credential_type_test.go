package types

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCredentialType_String(t *testing.T) {
	tests := []struct {
		name     string
		input    CredentialType
		expected string
	}{
		{
			name:     "pass - valid CredentialType",
			input:    CredentialType("my_credential"),
			expected: "my_credential",
		},
		{
			name:     "pass - empty URI",
			input:    CredentialType(""),
			expected: "",
		},
		{
			name:     "pass - CredentialType with special characters",
			input:    CredentialType("https://example.com/nft/1?query=param&another=param"),
			expected: "https://example.com/nft/1?query=param&another=param",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.input.String(); got != tt.expected {
				t.Errorf("CredentialType.String(), got: %v but we want: %v", got, tt.expected)
			}
		})
	}
}

func TestCredentialType_IsValid(t *testing.T) {
	tests := []struct {
		name     string
		credType CredentialType
		expected bool
	}{
		{
			name:     "fail - empty string",
			credType: "",
			expected: false,
		},
		{
			name:     "pass - valid hex string",
			credType: "6D795F63726564656E7469616C",
			expected: true,
		},
		{
			name:     "fail - invalid hex string",
			credType: "invalid",
			expected: false,
		},
		{
			name:     "fail - short hex string",
			credType: CredentialType(strings.Repeat("0", MinCredentialTypeLength-1)),
			expected: false,
		},
		{
			name:     "fail - long hex string",
			credType: CredentialType(strings.Repeat("0", MaxCredentialTypeLength+1)),
			expected: false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := test.credType.IsValid()
			assert.Equal(t, test.expected, result)
		})
	}
}
