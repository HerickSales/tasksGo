package main

import (
	"github.com/elipe1/tasksGo-tarefa2/config"
	"github.com/elipe1/tasksGo-tarefa2/router"
)

func main() {
	config.Init()
	router.Init()
}
