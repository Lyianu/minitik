package logic

import (
	"context"
	"net/http"

	"github.com/RaymondCode/simple-demo/api/user/internal/svc"
	"github.com/RaymondCode/simple-demo/api/user/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserRegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserRegisterLogic {
	return &UserRegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserRegisterLogic) UserRegister(req *types.UserRegisterRequest) (resp *types.UserRegisterResponse, err error) {
	resp = &types.UserRegisterResponse{
		StatusCode: http.StatusOK,
		UserId:     1,
		Token:      "abc",
	}

	return
}
