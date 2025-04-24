package config

import (
	"strconv"
	"strings"
)

type Config struct {
	Name  string
	Addr  string
	Host  string
	Port  int
	Debug bool
}

func NewConfiguration(params Config) (config Config, err error) {
	config.Addr = params.Addr

	addrParts := strings.Split(config.Addr, ":")
	config.Host = addrParts[0]

	port, err := strconv.Atoi(addrParts[1])
	if err != nil {
		return config, err
	}
	config.Port = port

	return config, nil
}
