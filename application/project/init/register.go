package init

import (
	"example/application/project/impl"
	"example/lib/expose_api/project"
)

func init() {
	// 注册 project 模块接口的实现
	project.ServiceImpl = &impl.Impl{}
}
