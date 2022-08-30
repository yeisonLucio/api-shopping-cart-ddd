package config

var App *AppConfig

type AppConfig struct {
	MysqlHost     string `mapstructure:"MYSQL_HOST"`
	MysqlUser     string `mapstructure:"MYSQL_USER"`
	MysqlPassword string `mapstructure:"MYSQL_PASSWORD"`
	MysqlPort     string `mapstructure:"MYSQL_PORT"`
	MysqlDbName   string `mapstructure:"MYSQL_DB_NAME"`
	ServerPort    string `mapstructure:"SERVER_PORT"`
	JwtSecretKey  string `mapstructure:"JWT_SECRET_KEY"`
	EnableGormLog bool   `mapstructure:"ENABLE_GORM_LOG"`
}
