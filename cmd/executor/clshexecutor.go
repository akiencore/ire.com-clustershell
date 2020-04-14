package main

import (
	"context"
	"fmt"
	"log"
	"sync"

	"ire.com/clustershell/communicate"
	"ire.com/clustershell/logger"
	"ire.com/clustershell/nodesvcs"
)

func main() {
	var executorsvc communicate.CommNode

	log.SetFlags(log.LstdFlags | log.Lshortfile)
	ctx := context.Background()
	wg := new(sync.WaitGroup)
	defer wg.Wait()
	
	returnLineChan := make(chan nodesvcs.ReturnLineChan, 1)
	executorsvc = &nodesvcs.ExecutorSVC{ReturnLineChan: returnLineChan}

	err := executorsvc.Init(ctx, wg)
	if err != nil {
		logger.Error("Init:", err)
		return
	}

	logger.Info("Executor is ready to get command.")

	fmt.Println(executorsvc)


}
