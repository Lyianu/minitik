// Code generated by goctl. DO NOT EDIT.
package types

type UserRegisterRequest struct {
	Username string `form:"username"`
	Password string `form:"password"`
}

type UserRegisterResponse struct {
	StatusCode int32  `json:"status_code"`
	StatusMsg  string `json:"status_msg"`
	UserId     int64  `json:"user_id"`
	Token      string `json:"token"`
}

type UserLoginRequest struct {
	Username string `form:"username"`
	Password string `form:"password"`
}

type UserLoginResponse struct {
	StatusCode int32  `json:"status_code"`
	StatusMsg  string `json:"status_msg"`
	UserId     int64  `json:"user_id"`
	Token      string `json:"token"`
}
