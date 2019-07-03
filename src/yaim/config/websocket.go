package config

import (
	"github.com/kataras/iris/context"
	"net/http"
	"time"
)
type wsConfig struct {
	// IDGenerator用于创建（以及稍后设置）
	//每个传入的websocket连接（客户端）的ID。
	//请求是一个参数，您可以使用它来生成ID（例如，来自标题）。
	//如果为空，则由DefaultIDGenerator生成ID：randomString（64）
	IDGenerator func(ctx context.Context) string
	Error       func(w http.ResponseWriter, r *http.Request, status int, reason error)
	CheckOrigin func(r *http.Request) bool
	// HandshakeTimeout指定握手完成的持续时间。
	HandshakeTimeout time.Duration
	//允许WriteTimeout时间向连接写入消息。
	// 0表示没有超时。
	//默认值为0
	WriteTimeout time.Duration
	//允许ReadTimeout时间从连接中读取消息。
	// 0表示没有超时。
	//默认值为0
	ReadTimeout time.Duration
	// PongTimeout允许从连接中读取下一个pong消息。
	//默认值为60 * time.Second
	PongTimeout time.Duration
	// PingPeriod将ping消息发送到此期间的连接。必须小于PongTimeout。
	//默认值为60 * time.Second
	PingPeriod time.Duration
	// MaxMessageSize连接允许的最大消息大小。
	//默认值为1024
	MaxMessageSize int64
	// BinaryMessages将其设置为true，以表示二进制数据消息而不是utf-8文本
	//兼容，如果您想使用Connection的EmitMessage将自定义二进制数据发送到客户端，就像本机服务器 - 客户端通信一样。
	//默认为false
	BinaryMessages bool
	// ReadBufferSize是下划线阅读器的缓冲区大小
	//默认值为4096
	ReadBufferSize int
	// WriteBufferSize是下划线编写器的缓冲区大小
	//默认值为4096
	WriteBufferSize int
	// EnableCompression指定服务器是否应尝试协商每个
	//消息压缩（RFC 7692）。将此值设置为true则不会
	//保证支持压缩。目前只有“没有背景
	//支持“接管”模式。
	EnableCompression bool
	//子协议按顺序指定服务器支持的协议
	//偏好。如果设置了此字段，则Upgrade方法通过使用协议选择此列表中的第一个匹配来协商子协议
	//客户要求。
	Subprotocols []string
}
