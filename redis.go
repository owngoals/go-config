package goconfig

import "github.com/spf13/viper"

type RedisConfig struct {
	Enable   bool
	Host     string
	Port     int
	Password string
	Database int
	LockDB   int
}

func redisConfig() map[string]RedisConfig {
	items := viper.GetStringMap("redis")

	m := make(map[string]RedisConfig)

	for k, v := range items {
		vv := v.(map[string]interface{})
		c := RedisConfig{}

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

		if val, exist := vv["password"]; exist && val != nil {
			c.Password = val.(string)
		} else {
			c.Password = ""
		}

		if val, exist := vv["database"]; exist && val != nil {
			c.Database = val.(int)
		} else {
			c.Database = 0
		}

		if val, exist := vv["lock_db"]; exist && val != nil {
			c.LockDB = val.(int)
		} else {
			c.LockDB = 0
		}

		m[k] = c
	}

	return m
}
