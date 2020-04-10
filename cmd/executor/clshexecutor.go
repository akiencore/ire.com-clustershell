package main

import (
	"context"
	"fmt"
	"sync"

	"ire.com/clustershell/communicate"
	"ire.com/clustershell/logger"
	"ire.com/clustershell/nodesvcs"
)

func main() {
	var executorsvc communicate.CommNode

	ctx := context.Background()

	wg := new(sync.WaitGroup)
	lineChan := make(chan nodesvcs.LineChan, 1)
	executorsvc = &nodesvcs.ExecutorSVC{LineChan: lineChan}

	executorsvc.Init(wg)

	executorsvc.HandleListenOnUDP(ctx)
	logger.Info("Executor is ready to get command.")

	fmt.Println(executorsvc)
	wg.Wait()

}
