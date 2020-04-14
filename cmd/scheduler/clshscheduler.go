package main

import (
	"context"
	"log"
	"sync"

	"ire.com/clustershell/logger"

	"ire.com/clustershell/communicate"
	"ire.com/clustershell/nodesvcs"
)

func main() {
	var schsvc communicate.CommNode

	log.SetFlags(log.LstdFlags | log.Lshortfile)
	ctx := context.Background()
	wg := new(sync.WaitGroup)
	defer wg.Wait()

	sendMsgChan := make(chan nodesvcs.SendMsgChan, 1)
	schsvc = &nodesvcs.SchedulerSVC{
		SendMsgChan: sendMsgChan,
	}
	err := schsvc.Init(ctx, wg)
	if err != nil {
		logger.Error("Init:", err)
		return
	}

	/*

		err = schsvc.SendMsg([]byte("Hello world!"), "127.0.0.1"+nodesvcs.XCTUDPPORT)
		if err != nil {
			logger.Error(err)
		}

			err := schsvc.ListenOnUnixSocket()
			if err != nil {
				logger.Error(err)
				return
			}
			logger.Info("Listening on unix socket...")

			err = schsvc.ListenOnUDP()
			if err != nil {
				logger.Error(err)
				return
			}
			logger.Info("Listening on UDP Port...")
			logger.Info("Init completes.")
	*/
	//	fmt.Println(schsvc)

}
