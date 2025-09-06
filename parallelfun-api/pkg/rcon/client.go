package rcon

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
	"net"
	"sync"
	"time"
)

const (
	AUTH           = 3
	AUTH_RESPONSE  = 2
	EXEC_COMMAND   = 2
	RESPONSE_VALUE = 0
)

// Client RCON客户端结构
type Client struct {
	conn      net.Conn
	addr      string
	password  string
	mu        sync.Mutex
	requestID int32
}

// NewClient 创建一个新的RCON客户端
func NewClient(addr, password string) *Client {
	return &Client{
		addr:      addr,
		password:  password,
		requestID: 0,
	}
}

// Connect 连接到RCON服务器并进行身份验证
func (c *Client) Connect() error {
	c.mu.Lock()
	defer c.mu.Unlock()

	conn, err := net.Dial("tcp", c.addr)
	if err != nil {
		return fmt.Errorf("failed to connect to RCON server: %w", err)
	}

	c.conn = conn

	// 进行身份验证
	err = c.authenticate()
	if err != nil {
		c.conn.Close()
		c.conn = nil
		return fmt.Errorf("authentication failed: %w", err)
	}

	return nil
}

// Close 关闭RCON连接
func (c *Client) Close() error {
	c.mu.Lock()
	defer c.mu.Unlock()

	if c.conn != nil {
		err := c.conn.Close()
		c.conn = nil
		return err
	}
	return nil
}

// authenticate 进行身份验证
func (c *Client) authenticate() error {
	// 发送认证请求
	err := c.sendPacket(AUTH, c.password)
	if err != nil {
		return fmt.Errorf("failed to send auth packet: %w", err)
	}

	// 读取响应
	id, typ, _, err := c.readPacket()
	if err != nil {
		return fmt.Errorf("failed to read auth response: %w", err)
	}

	// 检查认证结果
	if id == -1 {
		return fmt.Errorf("authentication failed: incorrect password")
	}

	if typ != AUTH_RESPONSE {
		return fmt.Errorf("unexpected response type during authentication: %d", typ)
	}

	return nil
}

// Execute 执行RCON命令
func (c *Client) Execute(command string) (string, error) {
	c.mu.Lock()
	defer c.mu.Unlock()

	if c.conn == nil {
		return "", fmt.Errorf("not connected to RCON server")
	}

	// 发送命令
	err := c.sendPacket(EXEC_COMMAND, command)
	if err != nil {
		return "", fmt.Errorf("failed to send command packet: %w", err)
	}

	// 读取响应
	var response string
	for {
		id, typ, body, err := c.readPacket()
		if err != nil {
			return "", fmt.Errorf("failed to read command response: %w", err)
		}

		// 检查请求ID是否匹配
		if id != c.requestID-1 {
			return "", fmt.Errorf("unexpected request ID: expected %d, got %d", c.requestID-1, id)
		}

		// 检查响应类型
		if typ != RESPONSE_VALUE {
			return "", fmt.Errorf("unexpected response type: %d", typ)
		}

		// 如果body为空，表示响应结束
		if len(body) == 0 {
			break
		}

		response += body
	}

	return response, nil
}

// sendPacket 发送数据包
func (c *Client) sendPacket(typ int32, body string) error {
	// 生成请求ID
	c.requestID++
	if c.requestID < 0 {
		c.requestID = 1
	}

	// 创建数据包
	buf := new(bytes.Buffer)
	binary.Write(buf, binary.LittleEndian, c.requestID)
	binary.Write(buf, binary.LittleEndian, typ)
	buf.Write([]byte(body))
	buf.Write([]byte{0, 0}) // 两个空字节结尾

	// 数据包长度
	length := int32(buf.Len())

	// 发送数据包
	err := binary.Write(c.conn, binary.LittleEndian, length)
	if err != nil {
		return err
	}

	_, err = c.conn.Write(buf.Bytes())
	return err
}

// readPacket 读取数据包
func (c *Client) readPacket() (id, typ int32, body string, err error) {
	// 读取数据包长度
	var length int32
	err = binary.Read(c.conn, binary.LittleEndian, &length)
	if err != nil {
		return 0, 0, "", err
	}

	// 验证长度
	if length < 10 || length > 4096 {
		return 0, 0, "", fmt.Errorf("invalid packet length: %d", length)
	}

	// 读取数据包内容
	data := make([]byte, length)
	_, err = io.ReadFull(c.conn, data)
	if err != nil {
		return 0, 0, "", err
	}

	// 解析数据包
	buf := bytes.NewBuffer(data)
	err = binary.Read(buf, binary.LittleEndian, &id)
	if err != nil {
		return 0, 0, "", err
	}

	err = binary.Read(buf, binary.LittleEndian, &typ)
	if err != nil {
		return 0, 0, "", err
	}

	// 读取body部分
	bodyBytes := make([]byte, len(data)-8) // 减去id和typ的长度
	_, err = buf.Read(bodyBytes)
	if err != nil {
		return 0, 0, "", err
	}

	// 移除末尾的两个空字节
	if len(bodyBytes) >= 2 {
		bodyBytes = bodyBytes[:len(bodyBytes)-2]
	}

	body = string(bodyBytes)
	return id, typ, body, nil
}

// SetTimeout 设置读写超时
func (c *Client) SetTimeout(timeout time.Duration) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	if c.conn != nil {
		return c.conn.SetDeadline(time.Now().Add(timeout))
	}
	return nil
}
