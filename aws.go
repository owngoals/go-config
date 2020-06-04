package goconfig

import "github.com/spf13/viper"

type AwsSESConfig struct {
	Enable         bool
	AccessIdKey    string
	AccessIdSecret string
	Region         string
	FromAddress    string
}

func awsSESConfig() map[string]AwsSESConfig {
	items := viper.GetStringMap("aws.ses")

	m := make(map[string]AwsSESConfig)

	for k, v := range items {
		vv := v.(map[string]interface{})

		c := AwsSESConfig{}

		if val, exist := vv["enable"]; exist && val != nil {
			c.Enable = val.(bool)
		} else {
			c.Enable = false
		}

		if val, exist := vv["access-id-key"]; exist && val != nil {
			c.AccessIdKey = val.(string)
		} else {
			c.AccessIdSecret = ""
		}

		if val, exist := vv["access-id-secret"]; exist && val != nil {
			c.AccessIdSecret = val.(string)
		} else {
			c.AccessIdSecret = ""
		}

		if val, exist := vv["region"]; exist && val != nil {
			c.Region = val.(string)
		} else {
			c.Region = ""
		}

		if val, exist := vv["from-address"]; exist && val != nil {
			c.FromAddress = val.(string)
		} else {
			c.FromAddress = ""
		}

		m[k] = c
	}

	return m
}
