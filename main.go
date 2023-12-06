package main

import (
	"example/application/account"
	"example/application/task"
	"example/lib/impl"
)

func main() {
	impl.Init()
	task.Detail()
	account.My()
}
