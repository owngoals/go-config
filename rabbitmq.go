package goconfig

import "github.com/spf13/viper"

type RabbitmqConfig struct {
	Enable   bool
	Host     string
	Port     int
	User     string
	Password string
}

func rabbitmqConfig() map[string]RabbitmqConfig {
	items := viper.GetStringMap("rabbitmq")

	m := make(map[string]RabbitmqConfig)

	for k, v := range items {
		vv := v.(map[string]interface{})

		c := RabbitmqConfig{}

		if val, exist := vv["enable"]; exist && val != nil {
			c.Enable = val.(bool)
		} else {
			c.Enable = false
		}

		if val, exist := vv["host"]; exist && val != nil {
			c.Host = val.(string)
		} else {
			c.Host = ""
		}

		if val, exist := vv["port"]; exist && val != nil {
			c.Port = val.(int)
		} else {
			c.Port = 0
		}

		if val, exist := vv["user"]; exist && val != nil {
			c.User = val.(string)
		} else {
			c.User = ""
		}

		if val, exist := vv["password"]; exist && val != nil {
			c.Password = val.(string)
		} else {
			c.Password = ""
		}

		m[k] = c
	}

	return m
}
