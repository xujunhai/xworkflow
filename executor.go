package xworkflow

import (
	"context"
	"sync"
)

// Executor 并行执行，串行执行
type Executor interface {
}

// ConcurrentExecutor Dag层执行器
type ConcurrentExecutor struct {
}

func (c ConcurrentExecutor) Execute(ctx context.Context, tasks []*Task) {
	var wg sync.WaitGroup
	rsChan := make(chan chan interface{}, len(tasks))
	for _, v := range tasks {
		//resolve input
		NewG(func(ctx context.Context, i interface{}) interface{} {
			return v.Operator.Execute(ctx, InputParam{})
		}, ctx, WithWg(&wg), WithChannel2(rsChan))

	}

	for {
		select {
		case <-ctx.Done():
		case rc, ok := <-rsChan:
			if !ok {
				return
			}
			r := <-rc
			switch r.(type) {

			}
		}
	}
}

// SeqExecutor 串行执行
type SeqExecutor struct {
}

// DistributedExecutor 分布式执行
type DistributedExecutor struct {

}
