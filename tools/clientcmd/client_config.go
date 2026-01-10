package clientcmd

import "time"

type Server struct {
	LocationOfOrigin         string
	Timeout                  time.Duration `yaml:"timeout,omitempty" mapstructure:"timeout,omitempty"`
	MaxRetries               int           `yaml:"max-retries,omitempty" mapstructure:"max-retries,omitempty"`
	RetryInterval            time.Duration `yaml:"retry-interval,omitempty" mapstructure:"retry-interval,omitempty"`
	Address                  string        `yaml:"address,omitempty,omitempty" mapstructure:"address,omitempty"`
	TLSServerName            string        `yaml:"tls-server-name,omitempty" mapstructure:"tls-server-name,omitempty"`
	InsecureSkipTLSVerify    bool          `yaml:"insecure-skip-tls-verify,omitempty" mapstructure:"insecure-skip-tls-verify,omitempty"`
	CertificateAuthority     string        `yaml:"certificate-authority,omitempty" mapstructure:"certificate-authority,omitempty"`
	CertificateAuthorityData []byte        `yaml:"certificate-authority-data,omitempty" mapstructure:"certificate-authority-data,omitempty"`
}

type AuthInfo struct {
	LocationOfOrigin      string
	ClientCertificate     string `yaml:"client-certificate,omitempty" mapstructure:"client-certificate,omitempty"`
	ClientCertificateData string `yaml:"client-certificate-data,omitempty" mapstructure:"client-certificate-data,omitempty"`
	ClientKey             string `yaml:"client-key,omitempty" mapstructure:"client-key,omitempty"`
	ClientKeyData         string `yaml:"client-key-data,omitempty" mapstructure:"client-key-data,omitempty"`
	Token                 string `yaml:"token,omitempty" mapstructure:"token,omitempty"`
	Username              string `yaml:"username,omitempty" mapstructure:"username,omitempty"`
	Password              string `yaml:"password,omitempty" mapstructure:"password,omitempty"`
	SecretID              string `yaml:"secret-id,omitempty" mapstructure:"secret-id,omitempty"`
	SecretKey             string `yaml:"secret-key,omitempty" mapstructure:"secret-key,omitempty"`
}

type Config struct {
	AuthInfo *AuthInfo `json:"user,omitempty" mapstructure:"user"`
	Server   *Server   `json:"server,omitempty" mapstructure:"server"`
}

func NewConfig() *Config {
	return &Config{
		AuthInfo: &AuthInfo{},
		Server:   &Server{},
	}
}
