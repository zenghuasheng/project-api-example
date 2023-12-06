package account

import "example/application/task/external"

func My() {
	project, _ := external.GetProject("projectUUID")
	println("account", project.Name)
}
