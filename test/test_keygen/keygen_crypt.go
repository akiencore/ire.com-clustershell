package main

import (
	"fmt"

	"ire.com/clustershell/crypting"
	"ire.com/clustershell/crypting/keypairs"
	"ire.com/clustershell/logger"

	cryptorand "crypto/rand"
)

// Test - pick up key pairs from keys.go
type Test struct {
	pvtkeystrScheduler string
	pubkeystrScheduler string
	pvtkeystrExecutor  string
	pubkeystrExecutor  string
}

var teststrings = `HelloWorld`

//rsa logic: length of passing is in crypting.GenerateKeyPair(asymmetricKeySize) from generateKeyPairs.go
func (t *Test) cryptingSimulationBoth(msg string) bool {
	fmt.Println("message length: ", len(msg))

	msgScheduler := msg
	//fmt.Println("Scheduler: ", cryptingSimulation(msgScheduler, t.pubkeystrScheduler, t.pvtkeystrScheduler))
	msgExecutor := msg
	//fmt.Println("Executor: ", cryptingSimulation(msgExecutor, t.pubkeystrExecutor, t.pvtkeystrExecutor), "\n")

	//return cryptingSimulation(msgScheduler, t.pubkeystrScheduler, t.pvtkeystrScheduler) && cryptingSimulation(msgExecutor, t.pubkeystrExecutor, t.pvtkeystrExecutor)
	return cryptingSimulationSymm(msgScheduler, t.pubkeystrScheduler, t.pvtkeystrScheduler) && cryptingSimulationSymm(msgExecutor, t.pubkeystrExecutor, t.pvtkeystrExecutor)
}

//crypting with only rsa keypairs
func cryptingSimulation(msg string, pubkeystr string, pvtkeystr string) bool {
	encrypted, err := crypting.EncryptWithPublicKeyStr([]byte(msg), pubkeystr)
	//fmt.Println(string(encrypted))
	if err != nil {
		logger.Error(err)
	}

	unencrypted, err := crypting.DecryptWithPrivateKeyStr(encrypted, pvtkeystr)
	//fmt.Println(string(unencrypted))
	if err != nil {
		logger.Error(err)
	}

	return string(unencrypted) == msg
}

//crypting with both rsa keypairs and symmetric keys
func cryptingSimulationSymm(msg string, pubkeystr string, pvtkeystr string) bool {
	encryptedMsg, encryptedKey := crypting.EncryptMsg([]byte(msg), pubkeystr)

	unencrypted, err := crypting.DecryptMsg(encryptedMsg, encryptedKey, pvtkeystr)
	if err != nil {
		logger.Error(err)
	}

	return string(unencrypted) == msg
}

func main() {
	fmt.Println("\n", "kengen_crypt", "\n")

	//fmt.Printf("Scheduler_private_key_byte_str\n%v", pvtkeystrScheduler)
	//fmt.Printf("Scheduler_public_key_byte_str\n%v", pubkeystrScheduler)

	//fmt.Printf("Executor_private_key_byte_str\n%v", pvtkeystrScheduler)
	//fmt.Printf("Executor_public_key_byte_str\n%v", pubkeystrScheduler)

	test := Test{keypairs.PvtkeyScheduler, keypairs.PubkeyScheduler, keypairs.PvtkeyExecutor, keypairs.PubkeyExecutor}

	str := make([]byte, 1000)
	cryptorand.Read(str)
	for { //this loop can show the maximum length of the message passes
		teststrings = teststrings + string(str)
		if !test.cryptingSimulationBoth(teststrings) {
			break
		}
	}

}
