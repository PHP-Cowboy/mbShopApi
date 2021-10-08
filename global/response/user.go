package response

import (
	"mbShopApi/utils/timeutil"
)

type UserList struct {
	Total uint32      `json:"total"` //总条数
	Data  []*UserInfo `json:"data"`  //用户信息
}

type UserInfo struct {
	Id uint64 `json:"id"`
	//Password string `json:"password"` //密码
	Mobile   string            `json:"mobile"`   //手机号
	NickName string            `json:"nickName"` //昵称
	Birthday timeutil.JsonTime `json:"birthday"` //生日
	Gender   uint32            `json:"gender"`   //0男；1女
	Role     uint32            `json:"role"`     // 1普通用户；2管理员用户
}
