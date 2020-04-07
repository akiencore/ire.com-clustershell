package main

import (
	"fmt"

	"ire.com/clustershell/communicate"
	"ire.com/clustershell/nodesvcs"
)

func main() {
	var schedulersvc communicate.CommNode

	schedulersvc = &nodesvcs.SchedulerSVC{}
	schedulersvc.Init()

	fmt.Println(schedulersvc)

}
