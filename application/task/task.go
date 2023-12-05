package task

import "example/application/task/external"

func Detail() {
	project, _ := external.GetProject("projectUUID")
	println(project.Name)
}
