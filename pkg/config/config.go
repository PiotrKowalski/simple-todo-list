package config

import (
	"github.com/spf13/viper"
)

func ReadEnvInt(key string) (int, error) {
	ok := viper.IsSet(key)
	if !ok {
		return 0, EnvVarNotSetError{key: key}
	}

	value := viper.GetInt(key)
	if value == 0 {
		return 0, EnvVarNotSetError{key: key}
	}

	return value, nil
}

func ReadEnvString(key string) (string, error) {
	ok := viper.IsSet(key)
	if !ok {
		return "", EnvVarNotSetError{key: key}
	}

	value := viper.GetString(key)
	if value == "" {
		return "", EnvVarNotSetError{key: key}
	}

	return value, nil
}

func init() {
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()
	err := viper.ReadInConfig()
	if err != nil {
		return
	}
}
