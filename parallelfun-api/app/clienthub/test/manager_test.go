package test

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"net/http"
	_ "net/http/pprof"
	"parallelfun-api/app/clienthub/internal/biz"
	"strconv"
	"sync"
	"testing"
	"time"
)

func Test_msmp(t *testing.T) {

	go func() {
		http.ListenAndServe("0.0.0.0:6060", nil)
	}()
	m := biz.NewConnManager(context.Background(), 512)
	defer m.Close()
	url := "ws://msmp.server.poyuan233.cn:8088"
	secret := "MjHrY9yN3WTUKXsgtB1bMxTtvWlnJwVAVEbLFT2z"

	clientNum := 16
	reqNum := 1024
	for i := 0; i < clientNum; i++ {
		_, err := m.NewConn_Test(url, secret, strconv.Itoa(i))
		if err != nil {
			log.Error(err)
			return
		}
	}
	p := &sync.Pool{
		New: func() any {
			return biz.RpcRequest{
				JsonRpc: "2.0",
				Method:  "server/status",
			}
		},
	}
	wg := sync.WaitGroup{}
	resMap := make(map[string]int, clientNum)
	resMapLock := sync.RWMutex{}
	for i := 0; i < clientNum; i++ {
		resMap[strconv.Itoa(i)] = 0
	}

	for i := 0; i < clientNum; i++ {
		wg.Add(1)
		go func(clientId string) {
			curReqNum := 0
			defer wg.Done()
			for {
				ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
				req := p.Get().(biz.RpcRequest)
				res, err := m.SendRpcMsg(context.Background(), clientId, &req)
				p.Put(req)
				if err != nil {
					log.Info("clientId:", clientId, "err", err)
					return
				}
				select {
				case <-ctx.Done():
					log.Info("clientId:", clientId, "timeout")
				case msg := <-res:
					log.Info("clientId:", clientId, "res", string(msg))
					resMapLock.Lock()
					resMap[clientId]++
					resMapLock.Unlock()
				}
				curReqNum++
				if curReqNum >= reqNum {
					return
				}
				//time.Sleep(time.Millisecond * 1000)

			}
		}(strconv.Itoa(i))
	}
	wg.Wait()

	log.Info("all done")
	log.Info(resMap)

}
