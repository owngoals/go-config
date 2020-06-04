package goconfig

import "github.com/spf13/viper"

type AliyunOSSConfig struct {
	Enable          bool
	AccessKeyId     string
	AccessKeySecret string
	Endpoint        string
	Bucket          string
	Domain          string
}

func aliyunOSSConfig() map[string]AliyunOSSConfig {
	items := viper.GetStringMap("aliyun.oss")

	m := make(map[string]AliyunOSSConfig)

	for k, v := range items {
		vv := v.(map[string]interface{})

		c := AliyunOSSConfig{}

		if val, exist := vv["enable"]; exist && val != nil {
			c.Enable = val.(bool)
		} else {
			c.Enable = false
		}

		if val, exist := vv["access-key-id"]; exist && val != nil {
			c.AccessKeyId = val.(string)
		} else {
			c.AccessKeyId = ""
		}

		if val, exist := vv["access-key-secret"]; exist && val != nil {
			c.AccessKeySecret = val.(string)
		} else {
			c.AccessKeySecret = ""
		}

		if val, exist := vv["endpoint"]; exist && val != nil {
			c.Endpoint = val.(string)
		} else {
			c.Endpoint = ""
		}

		if val, exist := vv["bucket"]; exist && val != nil {
			c.Bucket = val.(string)
		} else {
			c.Bucket = ""
		}

		if val, exist := vv["domain"]; exist && val != nil {
			c.Domain = val.(string)
		} else {
			c.Domain = ""
		}

		m[k] = c
	}

	return m
}
