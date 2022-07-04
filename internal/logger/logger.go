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

package logger

import (
	"os"

	"github.com/sirupsen/logrus"
)

func NewLogger(logType string, logLevel string) *logrus.Logger {
	var logger = logrus.New()
	logger.SetFormatter(&logrus.TextFormatter{})

	switch logType {
	case TypeJSON:
		logger.SetFormatter(&logrus.JSONFormatter{})
	case TypeText:
		logger.SetFormatter(&logrus.TextFormatter{})
	default:
		logger.Warnf("failed to parse log type %q, using default type %q", logLevel, TypeText)
	}
	// default log output
	logger.Out = os.Stdout
	// Log level
	level, err := logrus.ParseLevel(logLevel)
	if err == nil {
		logger.SetLevel(level)
	} else {
		// Default info level
		logger.SetLevel(logrus.InfoLevel)
		logger.Warnf("failed to parse log level %q, using default level %q", logLevel, "info")
	}

	return logger
}
