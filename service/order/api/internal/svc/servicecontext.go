package svc

import (
	"mall/service/order/api/internal/config"
	"mall/service/order/rpc/order"

	"github.com/zeromicro/go-zero/zrpc"
)

// ServiceContext 非常核心的服务上下文（ServiceContext） 定义
// 这是一个资源容器 它充当依赖注入容器的角色，统一管理 API 服务所需的全部资源和配置
type ServiceContext struct {
	// Config: 保存服务的所有配置信息
	Config config.Config
	// OrderRpc: 是 order RPC 服务的客户端实例。通过它，当前的 API 服务才能调用远程的 order.rpc 微服务的方法，比如 Create, GetOrder等
	OrderRpc order.Order
}

// NewServiceContext 创建一个 ServiceContext 实例 这是一个初始化函数（构造函数） 它接收配置参数 c，并构造一个完整的 ServiceContext实例返回
func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		// zrpc.MustNewClient(c.OrderRpc) 根据配置创建了一个到 RPC 服务的连接客户端
		// rder.NewOrder(...): 用这个连接客户端生成具体的 RPC 服务调用实例
		OrderRpc: order.NewOrder(zrpc.MustNewClient(c.OrderRpc)),
	}

	/**
	Name: order-api
	Host: 0.0.0.0
	Port: 8888
	# OrderRpc 客户端的配置
	OrderRpc:
	  Etcd:
	    Hosts:
	      - 127.0.0.1:2379  # etcd 注册中心地址
	    Key: order.rpc       # 要调用的目标 RPC 服务在 etcd 中的键名

	这种方式称为 基于服务发现的服务调用。工作流程如下：
	1.	order.rpc服务启动时，会将自己的服务名（order.rpc）和网络地址注册到 etcd 中。
	2.	order.api服务启动时，zrpc客户端会根据配置中的 Key（order.rpc）去 etcd 查询该服务当前可用的真实地址列表。
	3.	zrpc客户端获取地址后，会通过负载均衡策略选择一个地址建立 gRPC 连接，后续的 RPC 调用都通过这个连接进行
	*/
}
