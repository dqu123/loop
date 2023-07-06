package main

import (
	"github.com/dqu123/loop/controller"
	"github.com/dqu123/loop/logger"
)

func main() {
	err := controller.NewServer()
	if err != nil {
		logger.LogError("ERROR in controller.NewServer(): ", err)
	}
}
