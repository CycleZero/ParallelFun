package rcon

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
	"net"
	"sync"
)

const (
	ServerAuth          = 3
	ServerAuthResponse  = 2
	ServerExecCommand   = 2
	ServerResponseValue = 0
)

// Server RCON服务端结构
type Server struct {
	listener   net.Listener
	addr       string
	password   string
	running    bool
	clients    map[net.Conn]*clientInfo
	mu         sync.RWMutex
	handleFunc func(string) string
}

// clientInfo 客户端连接信息
type clientInfo struct {
	authenticated bool
}

// NewServer 创建一个新的RCON服务端
func NewServer(addr, password string) *Server {
	return &Server{
		addr:     addr,
		password: password,
		clients:  make(map[net.Conn]*clientInfo),
	}
}

// SetCommandHandler 设置命令处理函数
func (s *Server) SetCommandHandler(handler func(string) string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.handleFunc = handler
}

// Start 启动RCON服务端
func (s *Server) Start() error {
	var err error
	s.listener, err = net.Listen("tcp", s.addr)
	if err != nil {
		return fmt.Errorf("failed to start RCON server: %w", err)
	}

	s.running = true

	// 处理连接
	go s.acceptConnections()

	return nil
}

// acceptConnections 接受客户端连接
func (s *Server) acceptConnections() {
	for s.running {
		conn, err := s.listener.Accept()
		if err != nil {
			if s.running {
				// TODO: 记录错误日志
			}
			return
		}

		// 处理客户端连接
		go s.handleConnection(conn)
	}
}

// handleConnection 处理客户端连接
func (s *Server) handleConnection(conn net.Conn) {
	defer conn.Close()

	// 添加客户端到连接映射
	s.mu.Lock()
	s.clients[conn] = &clientInfo{authenticated: false}
	s.mu.Unlock()

	defer func() {
		s.mu.Lock()
		delete(s.clients, conn)
		s.mu.Unlock()
	}()

	// 处理客户端请求
	for {
		id, typ, body, err := s.readPacket(conn)
		if err != nil {
			// 客户端断开连接或发生错误
			return
		}

		switch typ {
		case ServerAuth:
			s.handleAuth(conn, id, body)
		case ServerExecCommand:
			s.handleCommand(conn, id, body)
		default:
			// 未知类型，关闭连接
			return
		}
	}
}

// handleAuth 处理认证请求
func (s *Server) handleAuth(conn net.Conn, id int32, password string) {
	s.mu.RLock()
	client := s.clients[conn]
	s.mu.RUnlock()

	authenticated := (password == s.password)
	client.authenticated = authenticated

	// 发送认证响应
	s.sendPacket(conn, id, ServerAuthResponse, "")

	// 如果认证失败，关闭连接
	if !authenticated {
		conn.Close()
	}
}

// handleCommand 处理命令执行请求
func (s *Server) handleCommand(conn net.Conn, id int32, command string) {
	s.mu.RLock()
	client := s.clients[conn]
	handler := s.handleFunc
	s.mu.RUnlock()

	// 检查是否已认证
	if !client.authenticated {
		// 未认证的客户端发送-1作为请求ID
		s.sendPacket(conn, -1, ServerAuthResponse, "")
		conn.Close()
		return
	}

	// 执行命令
	var response string
	if handler != nil {
		response = handler(command)
	} else {
		response = "Unknown command"
	}

	// 发送响应
	s.sendPacket(conn, id, ServerResponseValue, response)

	// 发送空响应包表示结束
	s.sendPacket(conn, id, ServerResponseValue, "")
}

// sendPacket 发送数据包
func (s *Server) sendPacket(conn net.Conn, id int32, typ int32, body string) error {
	// 创建数据包
	buf := new(bytes.Buffer)
	binary.Write(buf, binary.LittleEndian, id)
	binary.Write(buf, binary.LittleEndian, typ)
	buf.Write([]byte(body))
	buf.Write([]byte{0, 0}) // 两个空字节结尾

	// 数据包长度
	length := int32(buf.Len())

	// 发送数据包
	err := binary.Write(conn, binary.LittleEndian, length)
	if err != nil {
		return err
	}

	_, err = conn.Write(buf.Bytes())
	return err
}

// readPacket 读取数据包
func (s *Server) readPacket(conn net.Conn) (id, typ int32, body string, err error) {
	// 读取数据包长度
	var length int32
	err = binary.Read(conn, binary.LittleEndian, &length)
	if err != nil {
		return 0, 0, "", err
	}

	// 验证长度
	if length < 10 || length > 4096 {
		return 0, 0, "", fmt.Errorf("invalid packet length: %d", length)
	}

	// 读取数据包内容
	data := make([]byte, length)
	_, err = io.ReadFull(conn, data)
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

// Stop 停止RCON服务端
func (s *Server) Stop() error {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.running = false
	if s.listener != nil {
		err := s.listener.Close()
		// 关闭所有客户端连接
		for conn := range s.clients {
			conn.Close()
		}
		return err
	}
	return nil
}

// IsRunning 检查服务端是否正在运行
func (s *Server) IsRunning() bool {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.running
}
