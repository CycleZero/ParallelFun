package discovery

import (
	"context"
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/go-redis/redis/v8"
	"parallelfun-api/app/server/internal/biz"
)

type contextKey string

const ClientIDKey contextKey = "clientId"
const IsApplyConsistentHashKey contextKey = "isApplyConsistentHash"

// 使用时：ctx.Value(ClientIDKey)

type ConsistentDiscovery struct {
	discovery registry.Discovery // Kratos原生服务发现（Nacos）
	redis     *redis.Client      // Redis缓存
	service   string             // 目标服务名
	cHash     *biz.ConsistentHash
}

func (d *ConsistentDiscovery) GetService(ctx context.Context, serviceName string) ([]*registry.ServiceInstance, error) {
	clientId, ok := (ctx.Value(ClientIDKey)).(string)
	if !ok {
		return d.discovery.GetService(ctx, serviceName)
	}
	isApplyConsistentHash, ok := (ctx.Value(IsApplyConsistentHashKey)).(bool)
	if !ok || !isApplyConsistentHash {
		return d.discovery.GetService(ctx, serviceName)
	}
	instances, err := d.discovery.GetService(ctx, serviceName)
	if err != nil {
		return instances, err
	}
	d.cHash.Update(instances)
	ins, success := d.cHash.Get(clientId)
	if !success {
		return instances, nil
	}
	return []*registry.ServiceInstance{ins}, nil
}

func (d *ConsistentDiscovery) Watch(ctx context.Context, serviceName string) (registry.Watcher, error) {
	return d.discovery.Watch(ctx, serviceName)
}

func NewConsistentDiscovery(d registry.Discovery, redis *redis.Client, serviceName string) *ConsistentDiscovery {
	return &ConsistentDiscovery{
		discovery: d,
		redis:     redis,
		service:   serviceName,
		cHash:     biz.NewConsistentHash(128),
	}
}
