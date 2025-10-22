package climanager

import (
	"context"
	"github.com/gorilla/websocket"
	"log"
	"sync"
	"sync/atomic"
	"time"
)

type ClientConn struct {
	conn            *websocket.Conn
	writeChan       chan []byte  // 发送消息的缓冲channel
	heartBeatTicker *time.Ticker // 心跳定时器（可选，建议全局批量检查替代）
	lastActive      time.Time    // 最后活动时间
	clientId        string

	closeSignal context.CancelFunc
	ctx         context.Context
	closeOnce   sync.Once
	closed      atomic.Bool
	// 其他元信息：如连接状态、第三方标识等
}

func NewClientConn(conn *websocket.Conn, clientId string) *ClientConn {
	c := &ClientConn{
		conn:      conn,
		writeChan: make(chan []byte, 100),
		clientId:  clientId,
	}
	c.ctx, c.closeSignal = context.WithCancel(context.Background())
	go c.startReadLoop()
	go c.startWriteLoop()
	return c
}

func (c *ClientConn) Close() {
	c.closeOnce.Do(func() {
		if c.closed.Load() == true {
			return
		}
		c.closed.Store(true)
		err := c.conn.Close()
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
			_, _, err := c.conn.ReadMessage()
			if err != nil {
				// 记录错误，退出循环
				return
			}
			c.lastActive = time.Now()
			// 发送到业务层（如业务channel）
			//businessChan <- msg
		}

	}
}

func (c *ClientConn) startWriteLoop() {
	defer c.Close()

	for msg := range c.writeChan {
		if err := c.conn.WriteMessage(websocket.TextMessage, msg); err != nil {
			// 发送失败，退出循环
			return
		}
	}
}
