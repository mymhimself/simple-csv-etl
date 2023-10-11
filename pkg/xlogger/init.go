package xlogger

import (
	"strings"

	"github.com/mymhimself/logger"
	"github.com/mymhimself/simple-csv-reader/pkg/constants"
	"github.com/spf13/viper"
)

// initLogger initializes the logger
func InitLogger() error {
	loglevel := logger.LevelInfo
	switch strings.ToLower(viper.GetString(constants.LogLevel)) {
	case constants.LogLevel_Info:
		loglevel = logger.LevelInfo
	case constants.LogLevel_Fatal:
		loglevel = logger.LevelFatal
	case constants.LogLevel_Error:
		loglevel = logger.LevelError
	case constants.LogLevel_Debug:
		loglevel = logger.LevelDebug
	default:
	}

	ops := []logger.Option{
		logger.OptionSetLevel(loglevel),
		logger.OptionSetFormatter(logger.FormatterText),
		logger.OptionServiceName(constants.ServiceName),
		logger.OptionReportCaller(3),
	}

	err := logger.Init(ops...)
	return err
}
