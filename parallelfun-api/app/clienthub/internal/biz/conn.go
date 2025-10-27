package biz

import (
	"context"
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"strconv"
	"sync"
	"sync/atomic"
	"time"
)

type ClientConn struct {
	Conn      *websocket.Conn
	WriteChan chan []byte // 发送消息的缓冲channel
	ReadChan  chan []byte

	//respMap map[string]chan []byte

	HeartBeatTicker *time.Ticker // 心跳定时器（可选，建议全局批量检查替代）
	LastActive      time.Time    // 最后活动时间
	ClientId        string

	msgIdOnce sync.Once
	MsgId     uint64

	closeFunc context.CancelFunc
	ctx       context.Context
	closeOnce sync.Once
	closed    atomic.Bool
	// 其他元信息：如连接状态、第三方标识等
}

func NewClientConn(ctx context.Context, target string, secret string, clientId string, recvChan chan []byte) (*ClientConn, error) {
	headers := http.Header{}
	headers.Add("Authorization", "Bearer "+secret)
	conn, _, err := websocket.DefaultDialer.Dial(target, headers)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to server: %v", err)
	}
	mctx, cancel := context.WithCancel(ctx)
	c := &ClientConn{
		Conn:      conn,
		WriteChan: make(chan []byte, 128),
		ReadChan:  recvChan,
		ClientId:  clientId,
		ctx:       mctx,
		closeFunc: cancel,
		MsgId:     uint64(0),
	}
	c.ctx, c.closeFunc = context.WithCancel(context.Background())
	go c.startReadLoop()
	go c.startWriteLoop()
	return c, nil
}

func (c *ClientConn) Close() {
	c.closeOnce.Do(func() {
		if c.closed.Load() == true {
			return
		}
		c.closed.Store(true)
		err := c.Conn.Close()
		if err != nil {
			log.Println(err)
			return
		}
	})
}

func (c *ClientConn) startReadLoop() {
	defer c.Close() // 退出时清理连接
	for {
		select {
		case <-c.ctx.Done():
			return
		default:
			_, msg, err := c.Conn.ReadMessage()
			if err != nil {
				// 记录错误，退出循环
				return
			}
			c.LastActive = time.Now()
			// 发送到业务层（如业务channel）
			c.ReadChan <- msg
		}

	}
}

func (c *ClientConn) startWriteLoop() {
	defer c.Close()

	for {
		select {
		case <-c.ctx.Done():
			return
		case msg := <-c.WriteChan:
			log.Println("write msg:", string(msg))
			if err := c.Conn.WriteMessage(websocket.TextMessage, msg); err != nil {
				// 发送失败，退出循环
				return
			}
		}

	}

}

func (c *ClientConn) NextId() string {
	id := atomic.AddUint64(&c.MsgId, 1)
	return c.ClientId + "-" + strconv.FormatUint(id, 10)
}

func (c *ClientConn) IsAlive() bool {
	return time.Now().Sub(c.LastActive) < 10*time.Second
}
