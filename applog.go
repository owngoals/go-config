package goconfig

import "github.com/spf13/viper"

var (
	DefaultApplogPath = "app.log"
)

type ApplogConfig struct {
	Enable bool
	Path   string
}

func applogConfig() map[string]ApplogConfig {
	items := viper.GetStringMap("applog")

	m := make(map[string]ApplogConfig)

	for k, v := range items {
		vv := v.(map[string]interface{})

		c := ApplogConfig{}

		if val, exist := vv["enable"]; exist && val != nil {
			c.Enable = val.(bool)
		} else {
			c.Enable = false
		}

		if val, exist := vv["path"]; exist && val != nil {
			c.Path = val.(string)
		} else {
			c.Path = DefaultApplogPath
		}

		m[k] = c
	}

	return m
}
