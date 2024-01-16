package configservice

import (
	"errors"

	log "github.com/asfarsharief/money_management_backend/common/logingservice"
	"github.com/asfarsharief/money_management_backend/lib/config"
	"github.com/mitchellh/mapstructure"
	"github.com/spf13/viper"
)

type ConfigProvider struct {
}

// DEFAULT - default env
const DEFAULT = "default"
const LOG_LEVEL = "debug"

var ConfigFileMap map[string]string = make(map[string]string)
var ConfigPathMap map[string]string = make(map[string]string)

var Configuration *config.Configuration

func init() {
	ConfigFileMap[DEFAULT] = "config"
	ConfigPathMap[DEFAULT] = "./lib/config"
}

// LoadConfiguration - load configuration
func LoadConfiguration(runEnvMap map[string]string, env string) (*config.Configuration, error) {
	var cfg *config.Configuration
	cp := ConfigProvider{}

	configFile := getConfigFile(env)
	if len(configFile) == 0 {
		return nil, errors.New("configFile not available")
	}
	cfgMap := cp.GetConfiguration(configFile, getConfigPath(env), runEnvMap)

	//returned value is map[string]interface{}. To convert it to object, use the following or define mapstructure in struct.
	err := mapstructure.Decode(cfgMap, &cfg)
	if err != nil {
		log.Errorf("Error in mapstructure.Decode: %+v", err)
		return nil, errors.New("Error while decoding server configuration")
	}

	Configuration = cfg
	return cfg, nil

}

func getConfigFile(env string) string {
	configFile := ConfigFileMap[env]
	if len(configFile) > 0 {
		return configFile
	}
	return ConfigFileMap[DEFAULT]

}

func getConfigPath(env string) string {
	configPath := ConfigPathMap[env]
	if len(configPath) > 0 {
		return configPath
	}
	return ConfigPathMap[DEFAULT]

}

// GetConfiguration - pass on the configfile and its path and default config map, it will return interface after unmarshal. Convert and use it in the calling API.
func (cp ConfigProvider) GetConfiguration(configFile string, configPath string, runEnv map[string]string) interface{} {

	v := viper.New()
	v.SetConfigName(configFile)

	log.Debugf("Config file:      %s, Config file Path:      %s", configFile, configPath)

	// Set the path to look for the configurations file
	v.AddConfigPath(configPath)

	// Enable VIPER to read Environment Variables
	v.AutomaticEnv()

	v.SetConfigType("yml")

	err := v.ReadInConfig()
	if err != nil {
		log.Errorf("Error reading config file, %s", err)
	}

	// Set undefined variables
	for key, element := range runEnv {
		log.Debugf("Key: %s, Value: %s", key, element)
		if element != "" {
			v.Set(key, element)
		}

	}
	var config interface{}

	uerr := v.Unmarshal(&config)

	if uerr != nil {
		log.Error("Error unmarshalling the configuration, %s", uerr)
	}
	return config
}
