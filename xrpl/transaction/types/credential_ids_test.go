package types

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCredentialIDs_IsValid(t *testing.T) {
	tests := []struct {
		name          string
		credentialIDs CredentialIDs
		expected      bool
	}{
		{
			name:          "fail - empty",
			credentialIDs: CredentialIDs{},
			expected:      false,
		},
		{
			name:          "pass -valid",
			credentialIDs: CredentialIDs{"0000000000000000000000000000000000000000000000000000000000000000"},
			expected:      true,
		},
		{
			name: "pass - valid with two ids",
			credentialIDs: CredentialIDs{
				"0000000000000000000000000000000000000000000000000000000000000000",
				"6D795F63726564656E7469616C",
			},
			expected: true,
		},
		{
			name: "fail - invalid id, not hex",
			credentialIDs: CredentialIDs{
				"0000000000000000000000000000000000000000000000000000000000000000",
				"invalid",
			},
			expected: false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := test.credentialIDs.IsValid()
			require.Equal(t, test.expected, result)
		})
	}
}

func TestCredentialIDs_Flatten(t *testing.T) {
	tests := []struct {
		name          string
		credentialIDs CredentialIDs
		expected      []string
	}{
		{
			name:          "pass -empty",
			credentialIDs: CredentialIDs{},
			expected:      []string{},
		},
		{
			name:          "pass - valid",
			credentialIDs: CredentialIDs{"0000000000000000000000000000000000000000000000000000000000000000"},
			expected:      []string{"0000000000000000000000000000000000000000000000000000000000000000"},
		},
		{
			name:          "pass - valid with two ids",
			credentialIDs: CredentialIDs{"0000000000000000000000000000000000000000000000000000000000000000", "6D795F63726564656E7469616C"},
			expected:      []string{"0000000000000000000000000000000000000000000000000000000000000000", "6D795F63726564656E7469616C"},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := test.credentialIDs.Flatten()
			require.Equal(t, test.expected, result)
		})
	}
}
