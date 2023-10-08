package constants

const (
	Env = "env"
	// ServiceName = "SERVICE_NAME"
	ServiceName = "service_name"

	NetworkTCP = "tcp"

	Empty = ""

	DefaultConfigFileName = ConfigFileName + ConfigFileType

	PathWorkingDirectory       = "."
	PathWorkingDirectoryConfig = "./config"

	PathHome    = "$HOME"
	PathHomeApp = PathHome + "/." + ServiceName
	PathEtcApp  = "/etc/" + ServiceName

	//
	// ─────────────────────────────────────────────────── CONFIGURATION KEYWORDS ─────
	//

	Port        = "port"
	PortDefault = 9090

	Host             = "host"
	HostFrontRegular = Host + ".front_regular"
	HostFrontAdmin   = Host + ".front_admin"
	HostAPI          = Host + ".api"

	// ─── CONFIGURATION FILE ─────────────────────────────────────────────────────────

	// ConfigFileName is the name of configuration file.
	ConfigFileName = "config"

	// ConfigFileType is the type of configuration file.
	ConfigFileType = "yml"

	// ConfigFileFolder is a path inside the project may contain configuration file.
	ConfigFileFolder = "config/"

	// ConfigFilePathEtc is a path may contain configuration file.
	ConfigFilePathEtc = "/etc/krm-venture/ohs/"

	// ConfigFilePathHome is a path may contain configuration file.
	ConfigFilePathHome = "$HOME/.ohs/"

	// ─── REST ────────────────────────────────────────────────────────────────────

	// RestPrefix is the prefix for rest api configurations
	RestPrefix = "rest_api"

	// RestAddr is the key in configuration file for the address that rest client listens on
	RestPort = RestPrefix + ".port"

	// RestNetwork is the key in configuration file for the network protocol type for the rest listener.
	RestNetwork = RestPrefix + ".network"

	// ─── TLS ───────────────────────────────────────────────────────────────────

	// ─── Grpc Max Size Request Response ─────────────────────────────────────────────
	GrpcPigeonClientMaxRecvMsgSize = 1024 * 1024 * 50
	GrpcPigeonClientMaxSendMsgSize = 1024 * 1024 * 50

	// ─── AUTHENTICATION ───────────────────────────────────────────────────────────
	AuthenticationPrefix = "authentication"

	// AuthUsername is the username used for authentication of apis
	AuthUsername = AuthenticationPrefix + ".username"

	// AuthPassword is the username used for authentication of apis
	AuthPassword = AuthenticationPrefix + ".password"

	// ─── LOGGER ─────────────────────────────────────────────────────────────────────

	loggerPrefix = "logger"

	// LogLevel is the key in configuration file for the log level of the logger.
	LogLevel       = loggerPrefix + ".level"
	LogLevel_Info  = "info"
	LogLevel_Fatal = "fatal"
	LogLevel_Error = "error"
	LogLevel_Debug = "debug"
	// LogServiceName is the key in configuration file for the service name on which logger is working.
	LogServiceName = loggerPrefix + ".service_name"

	// LogFileName is the key in configuration file for the file's name on which we want to write the logs.
	LogFileName = loggerPrefix + ".logfile_name"

	TLS         = "tls"
	TLSCertFile = TLS + ".cert"
	TLSKeyFile  = TLS + ".key"

	// ─── INSTRUMENTATION ────────────────────────────────────────────────────────────

	instrumentationPrefix = "instrumentation"

	// InstrumentationClient is the key in configuration file for the client of instrumentation.
	// Its value in configuration file can be `prometheus` or `statsd`.
	InstrumentationClient = instrumentationPrefix + ".client"

	// InstrumentationPort is the key in configuration file for the port on which monitoring server runs in stand Alone mode.
	InstrumentationPort = instrumentationPrefix + ".port"

	// InstrumentationStandAlone
	InstrumentationStandAlone = instrumentationPrefix + ".standalone"

	// ─── DATABASE ───────────────────────────────────────────────────────────────────
	Dynamodb  = "dynamodb"
	Redis     = "redis"
	Addr      = "addr"
	RedisPass = "password"
	Region    = "region"
	TTL       = "ttl"

	Localhost     = "localhost"
	RedisAddr     = Redis + "." + Addr
	RedisTLS      = Redis + "." + "tls"
	RedisIndex    = Redis + "." + "index"
	RedisTLSCert  = Redis + "." + "cert_file"
	RedisTLSKey   = Redis + "." + "key_file"
	RedisPassword = Redis + "." + RedisPass

	// ─── CREDENTIAL ─────────────────────────────────────────────────────────────────

	HTTPXTokenHeader = "Authorization"

	// ─── LOGINSERVICE ───────────────────────────────────────────────────────────────

	LoginService          = "loginservice"
	LoginServiceAddr      = LoginService + "." + Addr // https://somelogin.tld
	LoginServicePass      = LoginService + ".password"
	LoginServiceURILogin  = LoginService + ".login"
	LoginServiceURILogout = LoginService + ".logout"

	// ─── ASSETS ─────────────────────────────────────────────────────────────────────
	Assets                                   = "assets"
	AssetsDefault                            = Assets + ".default"
	DefaultHTMLWebApp                        = AssetsDefault + ".html"
	DefaultHTMLWEBShareStructureAccessDenied = AssetsDefault + ".share_structure_access_denied"
	AssetsDefaultDir                         = AssetsDefault + ".dir"
	AssetsDefaultPath                        = AssetsDefault + ".path"

	// ─── HTTP ───────────────────────────────────────────────────────────────────────

	HttpContentType     = "Content-Type"
	HttpContentTypeHTML = "text/html; charset=UTF-8"
	HttpContentTypeJSON = "application/json"

	// ─── USER AGENT ─────────────────────────────────────────────────────────────────

	UserAgentDart = "Dart"

	// ─── PEGION ─────────────────────────────────────────────────────────────────────

	Pigeon          = "pigeon"
	PigeonPublisher = Pigeon + "." + publisher
	PigeonAddr      = Pigeon + "." + Addr
	PigeonPort      = Pigeon + "." + Port

	// ─── MINIO ─────────────────────────────────────────────────────────────────────
	MinIO           = "minio"
	MinIOAddr       = MinIO + "." + Addr
	MinIOSecretKey  = MinIO + ".secretkey"
	MinIOAccessKey  = MinIO + ".accesskey"
	MinIOBucketName = MinIO + ".bucketname"

	// ─── S3 ──────────────────────────────────────────────────────────────
	ObjectStorage = "object_storage"
	AWS           = "aws"
	AccessKey     = "access_key"
	SecretKey     = "secret_key"
	BucketName    = "bucket_name"
	AWSAccessKey  = ObjectStorage + "." + AWS + "." + AccessKey
	AWSRegion     = ObjectStorage + "." + AWS + "." + Region
	AWSSecretKey  = ObjectStorage + "." + AWS + "." + SecretKey
	AWSBucketName = ObjectStorage + "." + AWS + "." + BucketName

	// ─── Dynamodb ────────────────────────────────────────────────────────
	DynamodbAddr      = Dynamodb + "." + Addr
	DynamodbRegion    = Dynamodb + "." + Region
	DynamodbAccessKey = Dynamodb + "." + AccessKey
	DynamodbSecretKey = Dynamodb + "." + SecretKey

	// ─── InfluxDB ─────────────────────────────────────────────────────────────────────

	Influx        = "influx"
	InfluxAddress = Influx + "." + Addr
	Token         = "token"
	InfluxToken   = Influx + "." + Token
	Bucket        = "bucket"
	// Organization  = "organization"
	InfluxBucket = Influx + "." + Bucket
	InfluxORG    = Influx + "." + Organization

	// ─── REQUIRED REDIS KEYS FOR ANALYTIC ────────────────────────────────────────

	QuestionsKey     = "questions"
	TotalUsersKey    = "totalUsers"
	PaidUsersKey     = "paidUsers"
	OrganizationsKey = "organizations"
	CloudStorageKey  = "cloudStorage"
	TrialUsersKey    = "trialUsers"

	// ─── HERMES ─────────────────────────────────────────────────────────────────────
	HermesAddr = Hermes + "." + Addr
	// ─── Payment ────────────────────────────────────────────────────────────────────
	GRPCAddress        = "addrgrpc"
	HTTPAddress        = "addrhttp"
	PaymentGRPCAddress = Payment + "." + GRPCAddress
	PaymentHTTPAddress = Payment + "." + HTTPAddress

	IndividualProfileRecurringPriceID = "individual_profile_recurring_price_id"
	IndividualProfileOneTimePriceID   = "individual_profile_onetime_price_id"
	BusinessProfileRecurringPriceID   = "business_profile_recurring_price_id"

	PaymentIndividualProfileRecurringPriceID = Payment + "." + IndividualProfileRecurringPriceID
	PaymentIndividualProfileOneTimePriceID   = Payment + "." + IndividualProfileOneTimePriceID
	PaymentBusinessProfileRecurringPriceID   = Payment + "." + BusinessProfileRecurringPriceID

	// ─────────────────────────────────────────────────────────────────────────────

	Account                            = "account"
	TrialDuration                      = "trial_duration"
	InactiveOverTrialInterval          = "inactive_over_trial_interval"
	InactiveOverDueInterval            = "inactive_over_due_interval"
	OrganizationRenewalInterval        = "organization_renewal_interval"
	AccountTrialDuration               = Account + "." + TrialDuration
	AccountInactiveOverTrialInterval   = Account + "." + InactiveOverTrialInterval
	AccountInactiveOverDueInterval     = Account + "." + InactiveOverDueInterval
	AccountOrganizationRenewalInterval = Account + "." + OrganizationRenewalInterval

	ScriptPeriod         = "script_period"
	ScriptPeriodInMinute = ScriptPeriod + "." + "minute"
	ScriptPeriodInHour   = ScriptPeriod + "." + "hour"
	// ─── Flow type for upgrade user ─────────────────────────────────────────────────────
	UpgradeTo                   = "UpgradeTo"
	FlowTypeUpgradeToIndividual = UpgradeTo + "Individual"
	FlowTypeUpgradeToBusiness   = UpgradeTo + "Business"
	// ─── Standards ──────────────────────────────────────────────────────────────────
	Standards                      = "standards"
	TimeFormat                     = "time"
	StandardTimeFormat             = Standards + "." + TimeFormat
	LockoutDuration                = "lockout_duration"
	MaximumInvalidPasswordAttempts = "maximum_invalid_password_attempts"
	InvalidAttemptsTimeWindow      = "invalid_attempts_time_window"

	// ─────────────────────────────────────────────────────────────────────
	Maps             = "maps"
	Google           = "google"
	GoogleAPIKey     = "api_key"
	StripeAPIKey     = "stripe_api_key"
	MapsGoogleAPIKey = Maps + "." + Google + "." + GoogleAPIKey
	MapsTestMode     = Maps + "." + "test_mode"
	// ─────────────────────────────────────────────────────────────────────────────
	JWTSecret = "jwt_secret"

	// ─── WEBHOOK ─────────────────────────────────────────────────────────────────────
	WebHook            = "webhook"
	WebhookSecurityKey = WebHook + ".security"
	WebhookPublisher   = WebHook + "." + publisher
	WebhookConsumer    = WebHook + "." + Consumer

	// ─── SENTRY ────────────────────────────────────────────────────────────

	// Sentry holds the prefix for the fields defined under sentry in the config
	Sentry = "sentry"

	// SentryDsn holds the value of the dsn used for the initialization of
	// Sentry
	SentryDsn = Sentry + ".dsn"

	// Release holds the value of the release version
	SentryRelease = Sentry + ".release"

	// Environment holds the value of the environment
	SentryEnvironment = Sentry + ".environment"

	// AttachStacktrace holds a boolean value for sentry to attach the stack
	// trace or not
	SentryAttachStacktrace = Sentry + ".attach_stacktrace"
	SentryDebug            = Sentry + ".debug"

	// ─── Consumer Rmq ────────────────────────────────────────────────────
	QueueName            = "queue-name"
	Consumer             = "consumer"
	Name                 = "name"
	Key                  = "key"
	ExchangeName         = "exchange-name"
	ExchangeType         = "exchange-type"
	ConsumerHost         = Consumer + "." + Host
	ConsumerQueueName    = Consumer + "." + QueueName
	ConsumerKey          = Consumer + "." + Key
	ConsumerExchangeName = Consumer + "." + ExchangeName
	ConsumerName         = Consumer + "." + Name
	Firebase             = "firebase"
	Path                 = "path"
	FirebasePath         = Firebase + "." + Path
	LevelDBPath          = "level-db." + Path
	// ─── TLS ───────────────────────────────────────────────────────────────────────────────────

	// tlsPrefix is used as a prefix in tls constants
	tlsPrefix  = "tls"
	TLSEnabled = tlsPrefix + ".enable"
	// TLSServerSide holds the key in the configuration file that its true value indicates the type
	// of TLS is only server side. If mutual mode is true, server side value will not be checked.
	// If both of server side and mutual mode are false, then tls totally is false.
	TLSServerSide = tlsPrefix + ".server_side"

	// TLSMutual holds the key in configuration file that its true value indicates the type of TLS
	// is mutual. If mutual mode is true, client side value will not be checked. If both of client
	// side and mutual mode are false, then tls totally is false.
	TLSMutual = tlsPrefix + ".mutual"

	// TLSCaCertPath holds the key in configuration file that its value indicates the path of the
	// CA Cert file.
	TLSCaCertPath = tlsPrefix + ".ca_cert_path"

	// TLSClientCertPath holds the key in configuration file that its value indicates the path of
	// client Cert file.
	TLSClientCertPath = tlsPrefix + ".client_cert_path"

	// TLSClientKeyPath holds the key in configuration file that its value indicates the path of
	// client key file.
	TLSClientKeyPath = tlsPrefix + ".client_key_path"

	// Gateway config
	RateLimit          = "rate-limit"
	GatewayRateLimit   = Gateway + "." + RateLimit
	ExpirationDuration = "expiration_duration"

	Gmail         = "gmail"
	Username      = "username"
	Password      = "password"
	GmailUsername = Gmail + "." + Username
	GmailPassword = Gmail + "." + Password

	FileManager   = "file_manager"
	SchedulerNode = "scheduler_node"

	NextPageToken = "next_page_token"
	PageToken     = "page_token"

	Cache    = "cache"
	CacheTTL = Cache + "." + TTL
	// ─────────────────────────────────────────────────────────────────────

	Publisher = "publisher"
)

// Analytic Required Keys to reset information saved.
var RequiredKeys = []string{TotalUsersKey, PaidUsersKey, TrialUsersKey}
