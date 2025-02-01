package config

var (
	APP                      = "leaderboard-api"
	Environment              string
	Port                     string
	SentryDSN                string
	DatabaseConnectionString string

	AccessTokenSecretKey string
)

func Init() {
	Environment = getString("ENVIRONMENT", "development")
	Port = getString("PORT", "8080")

	DatabaseConnectionString = getString("DB_CONN_STRING", "")
	SentryDSN = getString("SENTRY_DSN", "")

	AccessTokenSecretKey = getString("ACCESS_TOKEN_SECRET_KEY", "")
}

func IsProductionEnvironment() bool {
	return Environment == "production"
}

func IsTestEnvironment() bool {
	return Environment == "test"
}
