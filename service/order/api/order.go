package main

import (
	"flag" // 用于解析命令行参数的包
	"fmt"

	"mall/service/order/api/internal/config"
	"mall/service/order/api/internal/handler"
	"mall/service/order/api/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

// 在启动程序时，能够通过命令行指定一个非默认的配置文件路径。
var configFile = flag.String("f", "etc/order.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	// 配置加载工具
	// 目的：集中管理服务的各项参数，如服务端口、数据库连接、日志配置、超时设置等。
	// 工作机制：conf.MustLoad是 go-zero 提供的配置加载工具，支持 JSON、YAML、TOML 等格式。
	// 如果加载失败，此方法会直接导致程序终止，确保了服务启动时配置的正确性
	conf.MustLoad(*configFile, &c)

	// 初始化了服务的上下文（Service Context）
	// - 目的：服务上下文是一个容器，用于托管整个服务生命周期中所需的共享资源，例如数据库连接、缓存客户端、RPC 客户端或其他自定义依赖项
	// - 好处：通过依赖注入的方式，将这些资源统一管理并传递给各个请求处理器（Handler），避免了在每个处理器中重复初始化，保证了资源的一致性和可测试性。
	ctx := svc.NewServiceContext(c)

	// 创建了一个 HTTP 服务器实例
	// 1、 核心组件：在 go-zero 内部，这会创建一个 engine结构体实例，它封装了路由、中间件链、熔断器、限流器等核心功能
	// 2、 配置集成：传入的 c.RestConf包含了 HTTP 服务所需的特定配置，服务器会根据这些配置（如端口号）进行初始化。
	//		MustNewServer同样意味着初始化失败会直接终止程序
	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	// 路由注册 (Route Registration)
	// 作用：此函数将具体的 URL 路径（Path）和 HTTP 方法（Method）与对应的处理函数（Handler）绑定起来，
	// 并通常会将前面创建的服务上下文 ctx注入到每个处理器中，以便处理器能访问数据库等资源
	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	// 最后，通过 server.Start()启动服务。
	// 流程：这个方法会阻塞当前 goroutine，开始监听指定的网络端口，等待并处理传入的 HTTP 请求
	server.Start()
}
