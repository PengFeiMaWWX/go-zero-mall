package handler

import (
	"net/http"

	"mall/service/order/api/internal/logic"
	"mall/service/order/api/internal/svc"
	"mall/service/order/api/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

/*
*

	参数：服务上下文容器
	集中管理了数据库连接、缓存客户端、配置信息等所有业务逻辑可能需要的资源
	通过这种方式，框架实现了依赖注入，使得资源在处理器和逻辑层之间共享，提高了代码的可维护性和可测试性
*/
func CreateHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// 首先声明一个 types.CreateRequest类型的变量 req，用于存储解析后的请求数据
		var req types.CreateRequest
		// 然后使用 httpx.Parse(r, &req)来解析 HTTP 请求 r中的数据并填充到 req结构体中
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}
		// 创建逻辑实例：通过 logic.NewCreateLogic创建业务逻辑层的实例 l。
		// 这里传入了两个重要参数：r.Context()和 svcCtx。r.Context()包含了当前请求的上下文信息（如超时控制），而 svcCtx则提供了业务逻辑所需的资源
		l := logic.NewCreateLogic(r.Context(), svcCtx)
		// 调用 l.Create(&req)方法，将解析好的请求参数传入，
		// 执行具体的业务逻辑（例如，验证订单数据、计算金额、写入数据库等）。这一步是真正处理“创建订单”这个业务需求的地方
		resp, err := l.Create(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
