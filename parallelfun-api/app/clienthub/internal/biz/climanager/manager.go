package climanager

import (
	"sync"
)

type ConnManager struct {
	Shards     []*Shard
	ShardCount int
}
type Shard struct {
	Mu    sync.RWMutex
	Conns map[string]*ClientConn // key: clientID
}

func (m *ConnManager) NewConn(clientId string, target string) {

}
