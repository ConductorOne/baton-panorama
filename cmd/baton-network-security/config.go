package main

import (
	"context"
	"fmt"

	"github.com/conductorone/baton-sdk/pkg/cli"
	"github.com/spf13/cobra"
)

// config defines the external configuration required for the connector to run.
type config struct {
	cli.BaseConfig `mapstructure:",squash"` // Puts the base config options in the same place as the connector options

	PanoramaUrl string `mapstructure:"panorama-url"`
	Username    string `mapstructure:"username"`
	Password    string `mapstructure:"password"`
}

// validateConfig is run after the configuration is loaded, and should return an error if it isn't valid.
func validateConfig(ctx context.Context, cfg *config) error {
	if cfg.PanoramaUrl == "" {
		return fmt.Errorf("panorama-url is required")
	}
	if cfg.Username == "" {
		return fmt.Errorf("username is required")
	}
	if cfg.Password == "" {
		return fmt.Errorf("password is required")
	}
	return nil
}

func cmdFlags(cmd *cobra.Command) {
	cmd.PersistentFlags().String("panorama-url", "", "Url of Panorama instance")
	cmd.PersistentFlags().String("username", "", "Username")
	cmd.PersistentFlags().String("password", "", "Password")
}
