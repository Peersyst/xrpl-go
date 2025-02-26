package integration

import (
	"os"
	"testing"
)

const (
	IntegrationEnvVar = "INTEGRATION"
)

// GetWebsocketEnv returns the integration environment.
// If the environment is not set, it skips the tests.
// This function is intended to be used in tests that need to run against a specific environment.
// Run it before creating the runner to retrieve the environment host and faucet provider.
func GetWebsocketEnv(t *testing.T) Env {
	if _, ok := IntegrationWebsocketEnvs[EnvKey(os.Getenv(IntegrationEnvVar))]; !ok {
		t.Skip("skipping integration tests")
	}

	return IntegrationWebsocketEnvs[EnvKey(os.Getenv(IntegrationEnvVar))]
}

// GetRPCEnv returns the integration environment.
// If the environment is not set, it skips the tests.
// This function is intended to be used in tests that need to run against a specific environment.
// Run it before creating the runner to retrieve the environment host and faucet provider.
func GetRPCEnv(t *testing.T) Env {
	if _, ok := IntegrationRPCEnvs[EnvKey(os.Getenv(IntegrationEnvVar))]; !ok {
		t.Skip("skipping integration tests")
	}

	return IntegrationRPCEnvs[EnvKey(os.Getenv(IntegrationEnvVar))]
}
