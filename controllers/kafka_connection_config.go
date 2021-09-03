package controllers

import (
	"errors"
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v2"
	ctrl "sigs.k8s.io/controller-runtime"
)

var (
	kc             = make(map[string]KafkaConnection)
	configSetupLog = ctrl.Log.WithName("config load")
)

// KafkaConnection ...
type KafkaConnection struct {
	Brokers          []string `yaml:"brokers"`
	SecurityProtocol string   `yaml:"security-protocol"`
	Username         string   `yaml:"username"`
	Password         string   `yaml:"password"`
}

// GetKafkaConnectionConfig eliminates process of reading yaml each time on a call
// instead it reads yaml only when 'kc' instance is nil
func getKafkaConnectionConfig() map[string]KafkaConnection {
	if kc == nil {
		err := LoadConfig()
		if err != nil {
			configSetupLog.Error(err, "error while reloading kafka connection config")
			return nil
		}
		return kc
	}

	return kc
}

// LoadConfig returns a new decoded ConnectionConfig struct
func LoadConfig() error {
	configPath, ok := os.LookupEnv("KAFKA_CONNECTION_CONFIG_PATH")
	if !ok {
		return errors.New("KAFKA_CONNECTION_CONFIG_PATH not set")
	}

	configFile, err := ioutil.ReadFile(configPath)
	if err != nil {
		return err
	}

	err = yaml.Unmarshal(configFile, &kc)
	if err != nil {
		return err
	}

	return nil
}
