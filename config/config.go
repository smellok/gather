package config

import (
	"fmt"
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

var _config Config

type Config struct {
	Mysql    MysqlConfig    `yaml:"Mysql"`
	RabbitMQ RabbitMQConfig `yaml:"RabbitMq"`
}

type MysqlConfig struct {
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
	User string `yaml:"user"`
	Pass string `yaml:"pass"`
	Db   string `yaml:"db"`
}

type RabbitMQConfig struct {
	Host              string `yaml:"host"`
	Port              int    `yaml:"port"`
	User              string `yaml:"user"`
	Pass              string `yaml:"pass"`
	PublisherConfirms bool   `yaml:"publisherConfirms"`
	VirtualHost       string `yaml:"virtualHost"`
	AcknowledgeMode   string `yaml:"acknowledgeMode"`
	Concurrency       int    `yaml:"concurrency"`
	MaxConcurrency    int    `yaml:"maxConcurrency"`
	Prefetch          int    `yaml:"prefetch"`
	CacheChannelSize  int    `yaml:"cacheChannelSize"`
}

func init() {
	fmt.Println("开始读取配置文件！")
	_config = Config{}
	buffer, err := ioutil.ReadFile("./app.yaml")
	err = yaml.Unmarshal(buffer, &_config)
	if err != nil {
		log.Fatalf(err.Error())
	}
	fmt.Printf("%v\n", _config)
	fmt.Println("读取配置文件完成！")
}

func Get() *Config {
	return &_config
}

func Mysql() *MysqlConfig {
	return &_config.Mysql
}

func RabbitMQ() *RabbitMQConfig {
	return &_config.RabbitMQ
}
