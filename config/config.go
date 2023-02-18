package config

type Config struct {
	MySQL  MySQL
	Hertz  Hertz
	JWT    JWT
	Upload Upload
	Redis  Redis
	Zap    Zap
}
