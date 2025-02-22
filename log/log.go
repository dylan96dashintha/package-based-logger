package log

import (
	"github.com/package-based-logger/config"
	logger "github.com/rs/zerolog"
)

type Logger interface {
	GetLogger(logConf config.LogConfig) (lg logger.Logger, err error)
	GetPackageBasedLogger(packageName string,
		logConf config.LogConfig) (lg logger.Logger, err error)
}
