package choreoconsul

import (
  "github.com/hashicorp/consul/api"
)

type ChoreoDiscoClient struct {
  client *api.Client
}

func New(config *api.Config) (*ChoreoDiscoClient, error) {
  c, err := api.NewClient(api.Config)
  if err != nil {
    return nil, err
  }
  return &c, nil
}
