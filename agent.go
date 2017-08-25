package agent

import (
  "github.com/hashicorp/consul/api"
)

type ChoreoAgent struct {
  source *api.Agent
}

type ChoreoAgents struct {
  agents []ChoreoAgent
}

func New(client *api.Client) ChoreoAgent {
  a := ChoreoAgent{ source: client.Agent() }
  return a
}
