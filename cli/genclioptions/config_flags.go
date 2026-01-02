package genclioptions

const (
	FlagSymbolConfig  = "symbolconfig"
	FlagBearerToken   = "user.token"
	FlagUsername      = "user.username"
	FlagPassword      = "user.password"
	FlagSecretID      = "user.secret-id"
	FlagSecretKey     = "user.secret-key"
	FlagCertFile      = "user.client-certificate"
	FlagKeyFile       = "user.client-key"
	FlagTLSServerName = "server.tls-server-name"
	FlagInsecure      = "server.insecure-skip-tls-verify"
	FlagCAFile        = "server.certificate-authority"
	FlagAPIServer     = "server.address"
	FlagTimeout       = "server.timeout"
	FlagMaxRetries    = "server.max-retries"
	FlagRetryInterval = "server.retry-interval"
)

type ConfigFlags struct {
	SymbolConfig *string

	BearerToken *string
	Username    *string
	Password    *string
	SecretID    *string
	SecretKey   *string

	Insecure      *bool
	TLSServerName *string
	CertFile      *string
	KeyFile       *string
	CAFile        *string
}
