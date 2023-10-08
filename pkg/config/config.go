package config

import (
	"log"

	"github.com/mymhimself/simple-csv-reader/pkg/constants"
	"github.com/spf13/viper"
)

func InitConfig(rootProjectPath string) error {
	// viper.SetEnvPrefix(constants.EnvironmentVariablePrefix)

	// ─── CONFIG THE PATH AND FILE NAME ──────────────────────────────────────────────
	viper.SetConfigName(constants.ConfigFileName)
	viper.SetConfigType(constants.ConfigFileType)

	viper.AddConfigPath(rootProjectPath)
	viper.AddConfigPath("/config")
	viper.AddConfigPath(".")
	viper.AddConfigPath(constants.ConfigFilePathEtc)
	viper.AddConfigPath(constants.ConfigFilePathHome)
	viper.AddConfigPath("./config")
	viper.AutomaticEnv()
	err := viper.ReadInConfig()
	if err != nil {
		log.Printf("Fatal error config file: %v \n", err)
		return err
	}

	viper.WatchConfig()

	// ─── SAFE WRITE CONFIG IF IT DOES NOT EXIST ────────────────────────────────────────────
	// viper.SafeWriteConfigAs(rootProjectPath + constants.ConfigFileFolder + constants.ConfigFileName + ".yml")

	// ─── INIT THE DEFAULT VALUES ────────────────────────────────────────────────────
	initDefault()

	return nil
}

// initDefault - Sets default values for the viper.
// Priority in viper: default < config.yml < osEnvironment < flags
func initDefault() {

	// ─── REST ───────────────────────────────────────────────────────────────────────
	viper.SetDefault(constants.RestPort, 8076)
	viper.SetDefault(constants.RestNetwork, "tcp")

	// ─── TLS ───────────────────────────────────────────────────────────────────
	viper.SetDefault(constants.TLSServerSide, false)
	viper.SetDefault(constants.TLSMutual, false)
	viper.SetDefault(constants.TLSCaCertPath, "")
	viper.SetDefault(constants.TLSClientCertPath, "")
	viper.SetDefault(constants.TLSClientKeyPath, "")

	// ─── LOGGER ─────────────────────────────────────────────────────────────────────
	// viper.SetDefault(constants.LogLevel, logger.LevelDebug)
	viper.SetDefault(constants.LogFileName, "")

	// ─── INSTRUMENTATION ────────────────────────────────────────────────────────────
	viper.SetDefault(constants.InstrumentationClient, "prometheus")
	viper.SetDefault(constants.InstrumentationPort, 8283)
	viper.SetDefault(constants.InstrumentationStandAlone, false)
}
