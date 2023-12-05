package impl

import "example/lib/expose_api/project"

// Impl 实现 Service 接口
type Impl struct {
}

func (s *Impl) GetProject(projectID string) (*project.Project, error) {
	return &project.Project{
		ID:   projectID,
		Name: "project name",
	}, nil
}
