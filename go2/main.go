package main

import (
	"fmt"
	"github.com/moleculer-go/moleculer"
	"github.com/moleculer-go/moleculer/broker"
	"sync"
	"time"
)

func main() {
	var bkr = broker.New(&moleculer.Config{
		LogLevel:    "info",
		Transporter: "nats://localhost:4222",
	})
	bkr.Start()
	time.Sleep(15 * time.Second)

	goResponse := <-bkr.Call("goMath.add", map[string]int{
		"a": 5,
		"b": 6,
	})
	fmt.Println("Go=>Go. Response from Go service goMath.add =>", goResponse)

	var wg sync.WaitGroup
	wg.Add(1)
	wg.Wait()
}
