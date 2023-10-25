package system

import (
	"github.com/huanggengzhong/ewa_admin_server/utils"
)

type Register struct {
	Name   string `form:"name" json:"name" binding:"required"`
	Mobile string `form:"mobile" json:"mobile" binding:"required,mobile"`
}

func (register Register) GetMessages() utils.ValidatorMessages {
	return utils.ValidatorMessages{
		"name.required":   "用户名称不能为空",
		"mobile.required": "手机号不能为空",
		"mobile.mobile":   "手机号格式不正确",
	}
}
