package external

import "example/lib/expose_api/project"

// 对外部的调用统一放在这里

// GetProject 获取项目
func GetProject(projectID string) (*project.Project, error) {
	return project.GetService().GetProject(projectID)
}
