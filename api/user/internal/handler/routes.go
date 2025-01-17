// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	"github.com/RaymondCode/simple-demo/api/user/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/register",
				Handler: UserRegisterHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/login",
				Handler: UserLoginHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/",
				Handler: UserInfoHandler(serverCtx),
			},
		},
		rest.WithPrefix("/douyin/user"),
	)
}
