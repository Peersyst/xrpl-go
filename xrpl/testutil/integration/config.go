package integration

const (
	DefaultMaxRetries  = 3
	DefaultWalletCount = 1
)

// RunnerConfig is the configuration for the integration test runner.
// It contains the configuration for the websocket client and the number of wallets to create.
type RunnerConfig struct {
	WalletCount int
	Client      Client
	MaxRetries  int
}

// Option is a function that modifies the RunnerConfig.
type Option func(*RunnerConfig)

// WithWallets sets the number of wallets to create.
func WithWallets(count int) Option {
	return func(c *RunnerConfig) {
		c.WalletCount = count
	}
}

// WithMaxRetries sets the maximum number of retries for a transaction.
func WithMaxRetries(maxRetries int) Option {
	return func(c *RunnerConfig) {
		c.MaxRetries = maxRetries
	}
}

// NewRunnerConfig creates a new RunnerConfig with the given websocket configuration and options.
func NewRunnerConfig(opts ...Option) *RunnerConfig {
	config := &RunnerConfig{
		WalletCount: DefaultWalletCount,
		MaxRetries:  DefaultMaxRetries,
	}

	for _, opt := range opts {
		opt(config)
	}

	return config
}
