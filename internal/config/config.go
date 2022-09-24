/*
Copyright © 2022 Michael Bruskov <mixanemca@yandex.ru>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package config

import (
	"os"

	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

type Config struct {
	HTTP    HTTPConfig `mapstructure:"http"`
	Log     LogConfig  `mapstructure:"log"`
	Version string
	Build   string
}

type HTTPConfig struct {
	ListenAddress   string `mapstructure:"listen_address"`
	ListenPort      string `mapstructure:"listen_port"`
	ForwardedHeader string `mapstructure:"forwarded_header"`
}

type LogConfig struct {
	Type  string `mapstructure:"type"`
	Level string `mapstructure:"level"`
}

func Init(version, build string) (*Config, error) {
	viper.AddConfigPath("configs")
	viper.AddConfigPath("../../configs")
	viper.AddConfigPath("/etc/ipd")
	viper.SetConfigName("ipd")
	viper.SetConfigType("yaml")

	// Defaults
	viper.SetDefault("http.forwarded_header", "X-Forwarded-For")

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	var cfg Config

	err := viper.Unmarshal(&cfg)
	if err != nil {
		return nil, err
	}

	cfg.Version = version
	cfg.Build = build

	// Flags
	var flagVersion bool

	flagset := pflag.NewFlagSet("ipdFlags", pflag.ExitOnError)
	flagset.BoolVarP(&flagVersion, "version", "v", false, "show version and build info")

	flagset.Parse(os.Args[1:])
	viper.BindPFlags(flagset)

	return &cfg, nil
}
