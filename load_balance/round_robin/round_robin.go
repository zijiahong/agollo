package round_robin

import (
	"github.com/zouyx/agollo/v2/env/config"
	"github.com/zouyx/agollo/v2/load_balance"
	"sync"
)

func init() {
	load_balance.SetLoadBalance(&RoundRobin{})
}

type RoundRobin struct {
}

func (r *RoundRobin) Load(servers *sync.Map) *config.ServerInfo {
	var returnServer *config.ServerInfo
	servers.Range(func(k, v interface{}) bool {
		server := v.(*config.ServerInfo)
		// if some node has down then select next node
		if server.IsDown {
			return true
		}
		returnServer = server
		return false
	})
	return returnServer
}
