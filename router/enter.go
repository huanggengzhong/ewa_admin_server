package router

import "github.com/huanggengzhong/ewa_admin_server/router/system"

type RouterGroup struct {
	System system.RouterGroup
}

var RouterGroupApp = new(RouterGroup)
