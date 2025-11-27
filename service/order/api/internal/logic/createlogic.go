package logic

import (
	"context"

	"mall/service/order/api/internal/svc"
	"mall/service/order/api/internal/types"
	"mall/service/order/rpc/types/order"

	"github.com/zeromicro/go-zero/core/logx"
)

// CreateLogic
/**
订单创建功能的业务逻辑核心
*/

type CreateLogic struct {
	logx.Logger
	ctx context.Context // ctx context.Context：用于传递请求的上下文信息，如设置超时、取消操作，以及获取在网关层注入的 JWT 身份信息等
	// 这是依赖注入的核心。它充当一个容器，包含了服务所需的所有共享资源，比如数据库连接、缓存客户端，以及这里非常重要的 RPC 客户端 OrderRpc
	// 通过 svcCtx，业务逻辑可以方便地使用这些资源，而无需关心它们是如何创建的。
	svcCtx *svc.ServiceContext
}

// NewCreateLogic函数是创建逻辑实例的构造函数，它接收来自上层（如 handler）的 context和 svcCtx，并正确地初始化 CreateLogic结构体。
// 这种模式保证了逻辑层与 HTTP 请求细节的解耦，使其更易于测试
func NewCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateLogic {
	return &CreateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// Create Create方法是真正的业务逻辑执行者
func (l *CreateLogic) Create(req *types.CreateRequest) (resp *types.CreateResponse, err error) {
	// 调用 RPC 服务：通过 l.svcCtx.OrderRpc.Create方法，向订单的 RPC 服务发起调用。
	// 这里会将控制权交给另一个专门的订单微服务（例如 order.rpc）来执行真正的创建订单、操作数据库等持久化操作
	res, err := l.svcCtx.OrderRpc.Create(l.ctx, &order.CreateRequest{
		// 该方法接收到来自 API 层的 types.CreateRequest请求参数
		// 将这些 API 层的参数转换为 RPC 服务所需的格式，即构造 order.CreateRequest对象
		Uid:    req.Uid,
		Pid:    req.Pid,
		Amount: req.Amount,
		Status: req.Status,
	})
	if err != nil {
		return nil, err
	}

	return &types.CreateResponse{
		Id: res.Id,
	}, nil
}
