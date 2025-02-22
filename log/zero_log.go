package log

import (
	"github.com/package-based-logger/config"
	logger "github.com/rs/zerolog"
	"os"
)

type loggerObj struct {
}

func NewLogger() Logger {
	return &loggerObj{}
}

func (l *loggerObj) GetLogger(logConf config.LogConfig) (lg logger.Logger, err error) {
	level, err := logger.ParseLevel(logConf.Level)
	if err != nil {
		return lg, err
	}
	log := logger.New(os.Stdout).Level(level).With().Timestamp().Logger()
	return log, nil
}

// GetPackageBasedLogger set a package specific log level
// if the package related log levels are not defined, global log level is keeps as the level.
func (l *loggerObj) GetPackageBasedLogger(packageName string,
	logConf config.LogConfig) (lg logger.Logger, err error) {
	for _, pkg := range logConf.Package {
		if pkg.Name == packageName {
			level, err := logger.ParseLevel(pkg.Level)
			if err != nil {
				return lg, err
			}
			log := logger.New(os.Stdout).Level(level).With().Timestamp().Logger()
			return log, nil

		}
	}

	level, err := logger.ParseLevel(logConf.Level)
	if err != nil {
		return lg, err
	}
	lg = logger.New(os.Stdout).Level(level).With().Timestamp().Logger()
	return lg, nil
}
