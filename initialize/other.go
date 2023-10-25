package initialize

import (
	"fmt"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"github.com/huanggengzhong/ewa_admin_server/utils"
	"reflect"
	"strings"
)

func OtherInit() {
	initializeValidator()
	fmt.Println("初始化自定义校验器")
}

func initializeValidator() {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		// 注册自定义验证器,例如下面mobile
		_ = v.RegisterValidation("mobile", utils.ValidateMobile)

		// 注册自定义 json tag 函数
		v.RegisterTagNameFunc(func(fld reflect.StructField) string {
			name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
			if name == "-" {
				return ""
			} else {
				return name
			}
		})

	}
}
