package client

import (
	"context"
	"crypto/tls"
	"github.com/osgochina/dmicro/drpc"
	"github.com/osgochina/dmicro/drpc/proto"
	"github.com/osgochina/dmicro/registry"
	"github.com/osgochina/dmicro/selector"
	"time"
)

type Options struct {
	Context           context.Context // 上下文
	Network           string          // 网络类型
	LocalIP           string          // 本地网络
	TlsCertFile       string
	TlsKeyFile        string
	TLSConfig         *tls.Config
	ProtoFunc         proto.ProtoFunc
	SessionAge        time.Duration
	ContextAge        time.Duration
	DialTimeout       time.Duration
	SlowCometDuration time.Duration
	BodyCodec         string
	PrintDetail       bool
	CountTime         bool
	HeartbeatTime     time.Duration
	Registry          registry.Registry
	Selector          selector.Selector
	RetryTimes        int
	GlobalPlugin      []drpc.Plugin
}

type Option func(*Options)

// NewOptions 初始化配置
func NewOptions(options ...Option) Options {
	opts := Options{
		Context:           context.Background(),
		Network:           "tcp",
		LocalIP:           "0.0.0.0",
		BodyCodec:         defaultBodyCodec,
		SessionAge:        defaultSessionAge,
		ContextAge:        defaultContextAge,
		DialTimeout:       defaultDialTimeout,
		SlowCometDuration: defaultSlowCometDuration,
		RetryTimes:        defaultRetryTimes,
		PrintDetail:       false,
		CountTime:         false,
		HeartbeatTime:     time.Duration(0),
		ProtoFunc:         drpc.DefaultProtoFunc(),
	}
	for _, o := range options {
		o(&opts)
	}

	return opts
}

func (that *Options) EndpointConfig() drpc.EndpointConfig {

	c := drpc.EndpointConfig{
		Network:           that.Network,
		LocalIP:           that.LocalIP,
		DefaultBodyCodec:  that.BodyCodec,
		DefaultSessionAge: that.SessionAge,
		DefaultContextAge: that.ContextAge,
		SlowCometDuration: that.SlowCometDuration,
		PrintDetail:       that.PrintDetail,
		CountTime:         that.CountTime,
		DialTimeout:       that.DialTimeout,
	}
	return c
}

// OptRegistry 设置服务注册中心
func OptRegistry(r registry.Registry) Option {
	return func(o *Options) {
		o.Registry = r
		// set in the selector
		_ = o.Selector.Init(selector.Registry(r))
	}
}

// OptSelector 设置选择器
func OptSelector(s selector.Selector) Option {
	return func(o *Options) {
		o.Selector = s
	}
}

// OptGlobalPlugin 设置插件
func OptGlobalPlugin(plugin ...drpc.Plugin) Option {
	return func(o *Options) {
		o.GlobalPlugin = append(o.GlobalPlugin, plugin...)
	}
}

// OptHeartbeatTime 设置心跳包时间
func OptHeartbeatTime(t time.Duration) Option {
	return func(o *Options) {
		o.HeartbeatTime = t
	}
}

// OptTlsFile 设置证书内容
func OptTlsFile(tlsCertFile string, tlsKeyFile string) Option {
	return func(o *Options) {
		o.TlsCertFile = tlsCertFile
		o.TlsKeyFile = tlsKeyFile
	}
}

// OptTlsConfig 设置证书对象
func OptTlsConfig(config *tls.Config) Option {
	return func(o *Options) {
		o.TLSConfig = config
	}
}

// OptProtoFunc 设置协议方法
func OptProtoFunc(pf proto.ProtoFunc) Option {
	return func(o *Options) {
		o.ProtoFunc = pf
	}
}

// OptRetryTimes 设置重试次数
func OptRetryTimes(n int) Option {
	return func(o *Options) {
		o.RetryTimes = n
	}
}

// OptSessionAge 设置会话生命周期
func OptSessionAge(n time.Duration) Option {
	return func(o *Options) {
		o.SessionAge = n
	}
}

// OptContextAge 设置单次请求生命周期
func OptContextAge(n time.Duration) Option {
	return func(o *Options) {
		o.ContextAge = n
	}
}

// OptSlowCometDuration 设置慢请求的定义时间
func OptSlowCometDuration(n time.Duration) Option {
	return func(o *Options) {
		o.SlowCometDuration = n
	}
}

// OptBodyCodec 设置消息内容编解码器
func OptBodyCodec(c string) Option {
	return func(o *Options) {
		o.BodyCodec = c
	}
}

// OptPrintDetail 是否打印消息详情
func OptPrintDetail(c bool) Option {
	return func(o *Options) {
		o.PrintDetail = c
	}
}

// OptCountTime 是否统计请求时间
func OptCountTime(c bool) Option {
	return func(o *Options) {
		o.CountTime = c
	}
}

// OptNetwork 设置网络类型
func OptNetwork(net string) Option {
	return func(o *Options) {
		o.Network = net
	}
}

// OptLocalIP 设置本地监听的地址
func OptLocalIP(addr string) Option {
	return func(o *Options) {
		o.LocalIP = addr
	}
}
