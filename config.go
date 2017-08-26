package config

import (
  "github.com/hashicorp/consul/api"
)

func NewDefaultConfig() *api.Config {
  conf := api.DefaultConfig()
}
