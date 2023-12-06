package impl

import (
	"example/lib/expose_api/project"
	projectImpl "example/lib/impl/project"
)

func Init() {
	// 注册 project 模块接口的实现
	project.ServiceImpl = &projectImpl.Impl{}
}
