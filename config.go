package main

import (
	vscale "github.com/vozerov/go-vscale"
)

type Config struct {
	Token            string
	TerraformVersion string
}

func (c *Config) Client() *vscale.WebClient {
	return vscale.NewClient(c.Token)
}
