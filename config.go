package goconfig

import (
	"github.com/spf13/viper"
)

var (
	DefaultConfigPath = "."
	DefaultConfigName = "config"
	DefaultConfigType = "yaml"
)

type Section string

const (
	App       Section = "app"
	MySQL     Section = "mysql"
	Redis     Section = "redis"
	Mongo     Section = "mongo"
	RabbitMQ  Section = "rabbitmq"
	Applog    Section = "applog"
	AliyunOSS Section = "aliyun.oss"
	AwsSES    Section = "aws.ses"
)

type Options struct {
	Path, Name, Type string
}

type Option func(o *Options)

func Path(s string) Option {
	return func(o *Options) {
		o.Path = s
	}
}

func Name(s string) Option {
	return func(o *Options) {
		o.Name = s
	}
}

func Type(s string) Option {
	return func(o *Options) {
		o.Type = s
	}
}

func newOptions(options ...Option) Options {
	o := Options{
		Path: DefaultConfigPath,
		Name: DefaultConfigName,
		Type: DefaultConfigType,
	}

	for _, v := range options {
		v(&o)
	}

	return o
}

func readInConfig(o Options) error {
	viper.SetConfigName(o.Name)
	viper.SetConfigType(o.Type)
	viper.AddConfigPath(o.Path)
	return viper.ReadInConfig()
}

func Config(sections []Section, options ...Option) (map[Section]interface{}, error) {
	o := newOptions(options...)
	if err := readInConfig(o); err != nil {
		return nil, err
	}

	m := make(map[Section]interface{})
	for _, s := range sections {
		switch s {
		case App:
			m[s] = appConfig()
		case MySQL:
			m[s] = mysqlConfig()
		case Redis:
			m[s] = redisConfig()
		case Mongo:
			m[s] = mongoConfig()
		case RabbitMQ:
			m[s] = rabbitmqConfig()
		case Applog:
			m[s] = applogConfig()
		case AliyunOSS:
			m[s] = aliyunOSSConfig()
		case AwsSES:
			m[s] = awsSESConfig()
		}
	}

	return m, nil
}
