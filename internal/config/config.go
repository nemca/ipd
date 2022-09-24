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
	ListenAddress   string `mapstructure:"listen-address"`
	ListenPort      string `mapstructure:"listen-port"`
	ForwardedHeader string `mapstructure:"forwarded-header"`
}

type LogConfig struct {
	Type  string `mapstructure:"type"`
	Level string `mapstructure:"level"`
}

func Init(version, build string) (*Config, error) {
	viper.AddConfigPath("configs")
	viper.AddConfigPath("../../configs") // need to tests
	viper.AddConfigPath("/etc/ipd")
	viper.SetConfigName("ipd")
	viper.SetConfigType("yaml")

	// Defaults
	viper.SetDefault("http.listen-address", "127.0.0.1")
	viper.SetDefault("http.listen-port", "8080")
	viper.SetDefault("http.forwarded-header", "X-Forwarded-For")
	viper.SetDefault("log.type", "text")
	viper.SetDefault("log.level", "info")

	flagset := pflag.NewFlagSet("ipd", pflag.ExitOnError)
	flagset.BoolP("version", "v", false, "Show version and build info")
	flagset.StringP("http.listen-address", "l", viper.GetString("http.listen-address"), "Listen address for HTTP requests")
	flagset.StringP("http.listen-port", "p", viper.GetString("http.listen-port"), "Listen port for HTTP requests")
	flagset.StringP("http.forwarded-header", "H", viper.GetString("http.forwarded-header"), "HTTP header with real IP, settled by proxy")
	flagset.StringP("log.type", "T", viper.GetString("log.type"), "Format to write logs. \"json\" or \"text\"")
	flagset.StringP("log.level", "L", viper.GetString("log.level"), "Log level")

	flagset.Parse(os.Args[1:])
	viper.BindPFlags(flagset)

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

	return &cfg, nil
}
