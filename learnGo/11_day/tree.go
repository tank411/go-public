package main

import (
	"encoding/json"
	"fmt"
	"time"
)

func main() {
	s := NewServices(
		SetName("peter"),
		SetTimeout(time.Second*5),
	)

	fmt.Println("name:", s.conf.Name)
	fmt.Println("time", s.conf.Timeout)

	{
		ret := struct {
			CollectInterval string `json:"collectInterval"`
			CollectRule     string `json:"collectRule"`
		}{
			CollectInterval: "1",
			CollectRule:     "123",
		}
		detail, _ := json.Marshal(&ret)
		fmt.Println(string(detail))
	}
}

type Option func(options *Config)

type Config struct {
	Name    string
	Timeout time.Duration
}

type Services struct {
	conf Config
}

func SetTimeout(t time.Duration) Option {
	return func(options *Config) {
		options.Timeout = t
	}
}

func SetName(name string) Option {
	return func(options *Config) {
		options.Name = name
	}
}

func NewServices(opts ...Option) Services {
	c := Config{}
	for _, op := range opts {
		op(&c)
	}
	s := Services{}
	s.conf = c
	return s
}
