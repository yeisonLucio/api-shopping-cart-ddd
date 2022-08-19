package helpers

import (
	"github.com/spf13/viper"
)

type App struct {
	Host string `mapstructure:"MYSQL_HOST"`
}

func LoadConfig(
	fileName string,
	path string,
	fileType string,
	config interface{},
) (err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName(fileName)
	viper.SetConfigType(fileType)

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(config)

	return
}
