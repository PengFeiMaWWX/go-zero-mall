package config

import (
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

// Config 配置结构体 微服务配置文件，它定义了服务的配置结构
type Config struct {

	// rest.RestConf是一个内嵌结构体，它定义了 HTTP/RESTful 服务运行所需的所有基础配置
	rest.RestConf

	// 这个匿名结构体用于配置 JWT（JSON Web Token）认证 所需的参数
	Auth struct {
		// 一个高度保密的字符串，用于签名和验证 JWT Token。这是最重要的安全密钥，在生产环境中必须严格保密。
		AccessSecret string
		// Token 的有效期，通常以秒为单位。例如，设置为 7200表示 Token 在 2 小时后过期
		AccessExpire int64
	}
	// 此字段的类型为 zrpc.RpcClientConf，它定义了如何连接和调用另一个名为 "order" 的微服务（即订单服务）
	OrderRpc zrpc.RpcClientConf

	//1、 基于服务发现（如 Etcd）：
	//	这是更常见的生产环境用法。配置中会指定 Etcd 服务器的地址 (Hosts) 和服务在注册中心唯一的键名 (Key)。
	//	客户端会根据服务名自动发现其所有可用实例
	// OrderRpc:
	//  Etcd:
	//    Hosts:
	//      - 127.0.0.1:2379
	//    Key: order.rpc # 订单服务在Etcd中的注册键

	//2、直连（Direct）：通常用于开发或测试。直接在配置中指定订单服务的一个或多个具体网络地址。
	// OrderRpc:
	//  Endpoints:
	//    - "127.0.0.1:8080" # 直接指定订单服务的地址和端口
	//  # 或者使用 Target 字段
	//  # Target: "127.0.0.1:8080"
}
