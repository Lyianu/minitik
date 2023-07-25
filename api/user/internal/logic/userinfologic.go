package logic

import (
	"context"
	"github.com/RaymondCode/simple-demo/api/user/internal/svc"
	"github.com/RaymondCode/simple-demo/api/user/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserInfoLogic {
	return &UserInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserInfoLogic) UserInfo(req *types.UserInfoRequest) (resp *types.UserInfoResponse, err error) {
	resp = &types.UserInfoResponse{
		StatusCode:      0,
		StatusMsg:       "",
		Id:              1,
		Name:            "zhanglei",
		FollowCount:     9999,
		FollowerCount:   99999,
		IsFollow:        false,
		Avatar:          "",
		BackgroundImage: "",
		Signature:       "test_signature",
		TotalFavorited:  0,
		WorkCount:       0,
		FavoriteCount:   0,
	}

	return
}
