package main

import (
	"fmt"

	"ire.com/clustershell/communicate"
	"ire.com/clustershell/nodesvcs"
)

func main() {
	var executorsvc communicate.CommNode

	executorsvc = &nodesvcs.ExecutorSVC{}
	executorsvc.Init()

	fmt.Println(executorsvc)

}
