package config

import (
	"embed"
	"fmt"

	"github.com/sirijagadeesh/sendMailUsingGmailAPI/validate"
	"github.com/spf13/viper"
)

// Templates folder will embed into go binary.
//go:embed templates
var Templates embed.FS

// GmailAPIConfig configurations.
type GmailAPIConfig struct {
	GmailClientID        string `mapstructure:"GMAILAPIClientID" validate:"required"`
	GmailClientSecret    string `mapstructure:"GMAILAPIClientSecret" validate:"required"`
	GmailAPIRedirectURI  string `mapstructure:"GMAILAPIRedirectURI" validate:"required,url"`
	GmailAPIRefreshToken string `mapstructure:"GMAILAPIRefreshToken" validate:"required"`
}

var configApp GmailAPIConfig

// GmailAPI configurations.
func GmailAPI() GmailAPIConfig {
	return configApp
}

// Load configurations.
func Load() error {
	viper := viper.New()
	viper.SetConfigType("env")
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		return fmt.Errorf("unable to read config :: %w", err)
	}

	if err := viper.Unmarshal(&configApp); err != nil {
		return fmt.Errorf("unable to unmarshal the config :: %w", err)
	}

	if err := validate.Struct(&configApp, "mapstructure"); err != nil {
		return fmt.Errorf("validation errors %w", err)
	}

	return nil
}
