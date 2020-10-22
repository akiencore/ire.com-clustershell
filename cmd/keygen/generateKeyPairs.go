package main

import (
	"fmt"

	"ire.com/clustershell/crypting"
)

const (
	// variable length of generated asymmetric keys(128, 256, 512, 2048...)
	asymmetricKeySize = 1024
)

func main() {
	fmt.Println("package keypairs\n")

	pvtkeyScheduler, pubkeyScheduler := crypting.GenerateKeyPair(asymmetricKeySize)
	pvtkeyExecutor, pubkeyExecutor := crypting.GenerateKeyPair(asymmetricKeySize)

	pvtbyteScheduler := crypting.PrivateKeyToBytes(pvtkeyScheduler)
	pubbyteScheduler := crypting.PublicKeyToBytes(pubkeyScheduler)

	pvtbyteExecutor := crypting.PrivateKeyToBytes(pvtkeyExecutor)
	pubbyteExecutor := crypting.PublicKeyToBytes(pubkeyExecutor)

	fmt.Printf("var PvtkeyScheduler = `%s`\n", string(pvtbyteScheduler))
	fmt.Printf("var PubkeyScheduler = `%s`\n", string(pubbyteScheduler))

	fmt.Printf("var PvtkeyExecutor = `%s`\n", string(pvtbyteExecutor))
	fmt.Printf("var PubkeyExecutor = `%s`\n", string(pubbyteExecutor))
}
