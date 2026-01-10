package rest

import (
	"fmt"
	"time"
)

type Config struct {
	Host    string
	APIPath string
	ContentConfig

	Username string
	Password string

	SecretID  string
	SecretKey string

	BearerToken     string
	BearerTokenFile string

	TLSClientConfig

	// UserAgent is optional field to set user agent on requests.
	UserAgent     string
	Timeout       time.Duration
	MaxRetries    int
	RetryInterval time.Duration
}

type ContentConfig struct {
	ServiceName        string
	AcceptContentTypes string
	ContentType        string
}

type sanitizedConfig *Config

func (c *Config) GoString() string {
	return c.String()
}

func (c *Config) String() string {
	if c == nil {
		return "<nil>"
	}
	cc := sanitizedConfig(CopyConfig(c))
	if cc.Password != "" {
		cc.Password = "--- REDACTED ---"
	}
	if cc.SecretKey != "" {
		cc.SecretKey = "--- REDACTED ---"
	}
	if cc.BearerToken != "" {
		cc.BearerToken = "--- REDACTED ---"
	}
	return fmt.Sprintf("%#v", *cc)
}

// TLSClientConfig contains settings to enable transport layer security
// for the clientcmd.
type TLSClientConfig struct {
	Insecure bool
	// ServerName is used to verify the hostname on the returned
	// certificates unless Insecure is true.
	ServerName string
	CertFile   string
	KeyFile    string
	CAFile     string
	CertData   []byte
	KeyData    []byte
	CAData     []byte
	NextProtos []string
}

// AddUserAgent add a http user-agent to the config.
func AddUserAgent(config *Config, userAgent string) *Config {
	fullUserAgent := userAgent
	config.UserAgent = fullUserAgent
	return config
}

// CopyConfig returns a copy of the given config.
func CopyConfig(config *Config) *Config {
	return &Config{
		Host:            config.Host,
		APIPath:         config.APIPath,
		ContentConfig:   config.ContentConfig,
		Username:        config.Username,
		Password:        config.Password,
		SecretID:        config.SecretID,
		SecretKey:       config.SecretKey,
		BearerToken:     config.BearerToken,
		BearerTokenFile: config.BearerTokenFile,
		TLSClientConfig: TLSClientConfig{
			Insecure:   config.TLSClientConfig.Insecure,
			ServerName: config.TLSClientConfig.ServerName,
			CertFile:   config.TLSClientConfig.CertFile,
			KeyFile:    config.TLSClientConfig.KeyFile,
			CAFile:     config.TLSClientConfig.CAFile,
			CertData:   config.TLSClientConfig.CertData,
			KeyData:    config.TLSClientConfig.KeyData,
			CAData:     config.TLSClientConfig.CAData,
			NextProtos: config.TLSClientConfig.NextProtos,
		},
		UserAgent: config.UserAgent,
		Timeout:   config.Timeout,
	}
}
