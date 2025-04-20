package config

import "fmt"

type Config struct {
	Port int
}

func NewConfiguration() (config Config, err error) {
	fmt.Println("hello world")
	return config, nil
}
