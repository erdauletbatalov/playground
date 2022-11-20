package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	Token TokenConfig
}

type TokenConfig struct {
	Token string
}

func GetConfig(path, name string) (*Config, error) {
	// read config file
	viper.AddConfigPath(path)
	viper.SetConfigName(name)
	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	// read config
	conf := &Config{}

	tokenSub := viper.Sub("messaging")
	if tokenSub != nil {
		readTokenConfig(tokenSub, &conf.Token)
	}

	return conf, nil
}

func readTokenConfig(tree *viper.Viper, c *TokenConfig) {

	c.Token = tree.GetString("token")
}
