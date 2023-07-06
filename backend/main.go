package main

import (
	"fmt"

	"github.com/dqu123/loop/controller"
)

func main() {
	err := controller.NewServer()
	if err != nil {
		fmt.Println("ERROR in controller.NewServer(): ", err)
	}
}
