package config

import (
	"strings"

	"github.com/spf13/viper"
)

type LoaderOption func(v *viper.Viper)

func WithConfigPath(p string) LoaderOption {
	return func(v *viper.Viper) {
		v.SetConfigFile(p)
	}
}

func PrepareLoader(options ...LoaderOption) *viper.Viper {
	v := viper.New()
	v.AutomaticEnv()
	v.SetEnvKeyReplacer(strings.NewReplacer("-", "_", ".", "_"))
	for _, opt := range options {
		opt(v)
	}
	return v
}
