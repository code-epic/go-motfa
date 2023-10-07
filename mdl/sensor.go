package mdl

import (
	"context"
	"runtime"
	"time"
)

// Sensor de decisiones
type Sensor struct {
	Code    int
	Task    int
	Memory  string
	Count   int
	Limit   int
	Space   int
	MaxCore int
	Min     int
}

func init() {
	var ld Load
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, time.Second*4)
	defer cancel()

	go ld.Analyzed(ctx)
	ld.ListApi()

	for {
		select {
		case <-ctx.Done():
			//  ld.TestinFnx()
			//	ld.SendData()
			print("Finalizando ")
			return
		}
	}
}

func (S *Sensor) Evaluar() {
	S.MaxCore = runtime.GOMAXPROCS(1)
	S.Count = 1000
}
