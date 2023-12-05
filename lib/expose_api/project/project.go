package project

type Project struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

var ServiceImpl Service

// Service 对外暴露的接口
type Service interface {
	// GetProject 获取项目信息
	GetProject(projectID string) (*Project, error)
}

func GetService() Service {
	if ServiceImpl == nil {
		panic("project service not init")
	}
	return ServiceImpl
}
