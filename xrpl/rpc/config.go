package rpc

import (
	"errors"
	"net/http"
	"strings"
	"time"

	"github.com/Peersyst/xrpl-go/xrpl/common"
)

var ErrEmptyURL = errors.New("empty port and IP provided")

type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

type Config struct {
	HTTPClient HTTPClient
	URL        string
	Headers    map[string][]string

	// Retry config
	maxRetries int
	retryDelay time.Duration

	// Fee config
	maxFeeXRP  float32
	feeCushion float32

	// Faucet config
	faucetProvider common.FaucetProvider

	timeOut time.Duration
}

type ConfigOpt func(c *Config)

func WithHTTPClient(cl HTTPClient) ConfigOpt {
	return func(c *Config) {
		c.HTTPClient = cl
	}
}

func WithMaxFeeXRP(maxFeeXRP float32) ConfigOpt {
	return func(c *Config) {
		c.maxFeeXRP = maxFeeXRP
	}
}

func WithFeeCushion(feeCushion float32) ConfigOpt {
	return func(c *Config) {
		c.feeCushion = feeCushion
	}
}

func WithFaucetProvider(fp common.FaucetProvider) ConfigOpt {
	return func(c *Config) {
		c.faucetProvider = fp
	}
}

func WithTimeOut(timeOut time.Duration) ConfigOpt {
	return func(c *Config) {
		c.timeOut = timeOut
	}
}

func NewClientConfig(url string, opts ...ConfigOpt) (*Config, error) {

	// validate a url has been passed in
	if len(url) == 0 {
		return nil, ErrEmptyURL
	}
	// add slash if doesn't already end with one
	if !strings.HasSuffix(url, "/") {
		url += "/"
	}

	cfg := &Config{
		HTTPClient: &http.Client{Timeout: common.DefaultTimeOut},
		URL:        url,
		Headers: map[string][]string{
			"Content-Type": {"application/json"},
		},

		maxRetries: common.DefaultMaxRetries,
		retryDelay: common.DefaultRetryDelay,

		maxFeeXRP:  common.DefaultMaxFeeXRP,
		feeCushion: common.DefaultFeeCushion,
	}

	for _, opt := range opts {
		opt(cfg)
	}
	return cfg, nil
}
