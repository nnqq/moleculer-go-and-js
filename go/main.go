package main

import (
	"fmt"
	"github.com/moleculer-go/moleculer"
	"github.com/moleculer-go/moleculer/broker"
	"sync"
	"time"
)

type MathService struct{}

func (s *MathService) Name() string {
	return "goMath"
}

func (s *MathService) Add(params moleculer.Payload) int {
	return params.Get("a").Int() + params.Get("b").Int()
}

func main() {
	var bkr = broker.New(&moleculer.Config{
		LogLevel:    "info",
		Transporter: "nats://localhost:4222",
	})
	bkr.Publish(&MathService{})
	bkr.Start()
	time.Sleep(5 * time.Second)

	jsResponse := <-bkr.Call("jsMath.add", map[string]int{
		"a": 3,
		"b": 4,
	})
	fmt.Println("Go=>JS. Response from JS service jsMath.add =>", jsResponse)

	var wg sync.WaitGroup
	wg.Add(1)
	wg.Wait()
}
