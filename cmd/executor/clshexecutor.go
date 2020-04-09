package main

import (
	"fmt"
	"sync"

	"ire.com/clustershell/communicate"
	"ire.com/clustershell/logger"
	"ire.com/clustershell/nodesvcs"
)

func main() {
	var executorsvc communicate.CommNode

	wg := new(sync.WaitGroup)
	executorsvc = &nodesvcs.ExecutorSVC{}
	executorsvc.Init(wg)

	executorsvc.HandleListenOnUDP()
	logger.Info("Executor is ready to get command.")

	fmt.Println(executorsvc)
	wg.Wait()

}
