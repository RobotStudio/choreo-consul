package agent

import (
  "github.com/hashicorp/consul/api"
)

type ChoreoDiscoAgent struct {
  node *api.Agent
}

type ChoreoDiscoAgents struct {
  agents []ChoreoDiscoAgent
}

func New(client *api.Client) ChoreoDiscoAgent {
  a := ChoreoDiscoAgent{ node: client.Agent() }
  return a
}
