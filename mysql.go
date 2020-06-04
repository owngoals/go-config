package goconfig

import (
	"github.com/spf13/viper"
)

type MySQLConfig struct {
	Enable    bool
	Host      string
	Port      int
	User      string
	Password  string
	Database  string
	Engine    string
	Charset   string
	Collation string
}

func mysqlConfig() map[string]MySQLConfig {
	items := viper.GetStringMap("mysql")

	m := make(map[string]MySQLConfig)

	for k, v := range items {
		vv := v.(map[string]interface{})
		c := MySQLConfig{}

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

		if val, exist := vv["database"]; exist && val != nil {
			c.Database = val.(string)
		} else {
			c.Database = ""
		}

		if val, exist := vv["engine"]; exist && val != nil {
			c.Engine = val.(string)
		} else {
			c.Engine = ""
		}

		if val, exist := vv["charset"]; exist && val != nil {
			c.Charset = val.(string)
		} else {
			c.Charset = ""
		}

		if val, exist := vv["collation"]; exist && val != nil {
			c.Collation = val.(string)
		} else {
			c.Collation = ""
		}

		m[k] = c
	}

	return m
}
