package goconfig

import "testing"

func TestConfig(t *testing.T) {
	sections := []Section{App, MySQL, Redis, Mongo, RabbitMQ, Applog, AliyunOSS, AwsSES}

	m, err := Config(sections)
	if err != nil {
		t.Log(err)
		t.FailNow()
	}

	t.Log(m[App])
	t.Log(m[MySQL])
	t.Log(m[Redis])
	t.Log(m[Mongo])
	t.Log(m[RabbitMQ])
	t.Log(m[Applog])
	t.Log(m[AliyunOSS])
	t.Log(m[AwsSES])
}
