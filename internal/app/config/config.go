package config

import "github.com/spf13/viper"

type Config struct {
	Server struct {
		Host string `mapstructure:"host" json:"host" yaml:"host"`
		Port string `mapstructure:"port" json:"port" yaml:"port"`
	} `mapstructure:"server" json:"server" yaml:"server"`
}

func LoadConfig(filepath string) (*Config, error) {
	var config Config

	viper.SetConfigFile(filepath)
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	if err := viper.Unmarshal(&config); err != nil {
		return nil, err
	}

	return &config, nil
}
