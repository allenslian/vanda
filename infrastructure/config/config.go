package config

type (
	// Configuration describes all the sections in vanda configuration file.
	Configuration struct {
		Application applicationSection `mapstructure:"application"`
		Network     networkSection     `mapstructure:"network"`
		Database    databaseSection    `mapstructure:"database"`
		Cache       cacheSection       `mapstructure:"cache"`
		Broker      brokerSection      `mapstructure:"broker"`
		Security    securitySection    `mapstructure:"security"`
		Log         logSection         `mapstructure:"log"`
	}

	// ApplicationSection describes application section in the configuration file.
	applicationSection struct {
		ServiceName string `mapstructure:"service_name"`
		APIURI      string `mapstructure:"api_uri"`
		StaticURI   string `mapstructure:"static_uri"`
		StaticDir   string `mapstructure:"static_dir"`
		TemplateDir string `mapstructure:"template_dir"`
		UploadDir   string `mapstructure:"upload_dir"`
		PageSize    int    `mapstructure:"page_size"`
	}

	// NetworkSection describes network section in the configuration file.
	networkSection struct {
		Host            string `mapstructure:"host"`
		Listen          string `mapstructure:"listen"`
		AutoDiscovery   bool   `mapstructure:"auto_discovery"`
		RegistryAddress string `mapstructure:"registry_address"`
	}

	// DatabaseSection describes database section in the configuration file.
	databaseSection struct {
		DefaultURI  string `mapstructure:"default_uri"`
		ReadonlyURI string `mapstructure:"readonly_uri"`
		MaxOpen     int    `mapstructure:"sql_max_open"`
		MaxWait     int    `mapstructure:"sql_max_wait"`
	}

	// CacheSection describes cache section in the configuration file.
	cacheSection struct {
		KVURI      string `mapstructure:"kv_uri"`
		KVPassword string `mapstructure:"kv_password"`
		CookieName string `mapstructure:"cookie_name"`
	}

	// BrokerSection describes broker section in the configuration file.
	brokerSection struct {
		DefaultURI string `mapstructure:"default_uri"`
	}

	// SecuritySection describes security section in the configuration file.
	securitySection struct {
		SSLCertificate    string `mapstructure:"ssl_certificate"`
		SSLCertificateKey string `mapstructure:"ssl_certificate_key"`
		SecretKey         string `mapstructure:"secret_key"`
		EncryptKey        string `mapstructure:"encrypt_key"`
		TokenExpiry       int    `mapstructure:"token_expiry"`
	}

	// LogSection describes log section in the configuration file.
	logSection struct {
		FileName string `mapstructure:"filename"`
		Level    string `mapstructure:"level"`
		Rotation bool   `mapstructure:"rotation"`
		MaxSize  int    `mapstructure:"maxsize"`
	}
)
