package config

import (
	"backend/internal/core/utils"
	"bytes"
	"strings"

	"fmt"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"

	//	_ "github.com/spf13/viper/remote"
	"os"
	"strconv"
)

// New -> example c := config.New("config/.env", "config/config_api.yml", false)
func New(defaultEnvFile string, defaultConfigFile string, defaultEnableRemoteConfig bool) *ConfigFlag {
	useDefaultEnvFileEnv := os.Getenv("USE_DEFAULT_ENV_FILE")
	configFileEnv := os.Getenv("CONFIG_FILE")
	enableRemoteConfigEnv := os.Getenv("ENABLE_REMOTE_CONFIG")
	remoteConfigTypeEnv := os.Getenv("REMOTE_CONFIG_TYPE")

	var configFile string
	if configFileEnv == "" {
		configFile = defaultConfigFile
	} else {
		configFile = configFileEnv
	}

	var envFile string
	if useDefaultEnvFileEnv != "" {
		useDefaultEnvFile, err := strconv.ParseBool(useDefaultEnvFileEnv)
		if err != nil {
			envFile = defaultEnvFile
		} else {
			if useDefaultEnvFile {
				envFile = defaultEnvFile
			}
		}
	} else {
		envFile = defaultEnvFile
	}

	enableRemoteConfig := defaultEnableRemoteConfig
	if enableRemoteConfigEnv != "" {
		enableRemoteConf, err := strconv.ParseBool(enableRemoteConfigEnv)
		if err != nil {
			enableRemoteConfig = defaultEnableRemoteConfig
		} else {
			enableRemoteConfig = enableRemoteConf
		}
	}

	return &ConfigFlag{
		EnvFile:            envFile,
		ConfigFile:         configFile,
		EnableRemoteConfig: enableRemoteConfig,
		RemoteConfigType:   remoteConfigTypeEnv,
	}
}

type ConfigFlag struct {
	EnvFile            string
	ConfigFile         string
	EnableRemoteConfig bool
	RemoteConfigType   string
}

func mapperEnv(placeholderName string) string {
	split := strings.Split(placeholderName, ":")
	var defValue string
	if len(split) == 2 {
		placeholderName = split[0]
		defValue = split[1]
	}

	val, ok := os.LookupEnv(placeholderName)
	if !ok {
		return defValue
	}

	return val
}

func LoadConfig(c *ConfigFlag) (*Configuration, error) {
	fmt.Println(utils.ToJson(c))

	var configuration = &Configuration{}

	if utils.FileExists(c.EnvFile) {
		err := godotenv.Load(c.EnvFile)
		if err != nil {
			return nil, err
		}
	}

	fmt.Println("Load configuration file : " + c.ConfigFile)
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	ymlData, err := utils.ReadFile(c.ConfigFile)
	if err != nil {
		return nil, err
	}

	value := utils.ParseEnvVar(string(ymlData), os.Getenv)
	err = viper.ReadConfig(bytes.NewReader([]byte(value)))
	if err != nil {
		return nil, err
	}

	err = viper.Unmarshal(configuration)
	if err != nil {
		return nil, fmt.Errorf("unable to decode into struct, %v", err)
	}

	return configuration, nil
}
