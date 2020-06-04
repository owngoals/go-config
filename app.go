package goconfig

import "github.com/spf13/viper"

type AppConfig struct {
	Name  string
	Port  int
	Debug bool
	Node  int
}

func appConfig() AppConfig {
	c := AppConfig{}

	c.Name = viper.GetString("app.name")
	c.Port = viper.GetInt("app.name")
	c.Debug = viper.GetBool("app.debug")
	c.Node = viper.GetInt("app.node")

	return c
}
