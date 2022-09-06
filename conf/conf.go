package conf

import (
	"io/ioutil"

	"gopkg.in/yaml.v3"
)

func GetConf() (*Config, error) {
	c := &Config{}

	yamlFile, err := ioutil.ReadFile(ConfigFilePath)
	if err != nil {
		return c, err
	}
	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		return c, err
	}

	return c, nil
}

type Config struct {
	Telegram Telegram  `json:"telegram"`
	Mqtt     Mqtt      `json:"mqtt"`
	Devices  []Devices `json:"devices"`
}
type Telegram struct {
	Token  string `json:"token"`
	ChatID int64  `json:"chat_id"`
}
type Mqtt struct {
	Host     string      `json:"host"`
	Port     int         `json:"port"`
	Username interface{} `json:"username"`
	Password interface{} `json:"password"`
}
type Commands struct {
	Start string `json:"start"`
	Stop  string `json:"stop"`
	Stats string `json:"stats"`
}
type Devices struct {
	ID              string   `json:"_id"`
	Name            string   `json:"name"`
	Enabled         bool     `json:"enabled"`
	Type            string   `json:"type"`
	MqttStatsTopic  string   `json:"mqtt_stats_topic"`
	MqttConfigTopic string   `json:"mqtt_config_topic"`
	Commands        Commands `json:"commands,omitempty"`
}
