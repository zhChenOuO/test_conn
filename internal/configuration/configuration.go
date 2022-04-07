package configuration

import (
	"errors"

	"github.com/spf13/viper"
)

// Init 初始化 config & log & Global Setting
func Init(input interface{}) error {
	viper.AutomaticEnv()
	configPath := viper.GetString("CONFIG_PATH")
	if configPath == "" {
		configPath = viper.GetString("PROJ_DIR")

		if viper.GetString("PROJ_DIR") == "" {
			return errors.New("PROJ_DIR is required")
		}
	}

	configName := viper.GetString("CONFIG_NAME")
	if configName == "" {
		configName = "app"
	}

	viper.AddConfigPath(configPath)
	viper.SetConfigName(configName)
	viper.SetConfigType("yaml")

	if err := viper.ReadInConfig(); err != nil {
		return err
	}
	if err := viper.Unmarshal(input); err != nil {
		return err
	}

	return nil
}
