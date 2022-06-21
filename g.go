package xworkflow

import (
	"context"
	"fmt"
	"go.uber.org/zap"
	"sync"
)

var recoverFunc = func(err *error, msg string) {
	if p := recover(); p != nil {
		*err = fmt.Errorf("%v", p)
		zap.Stack("stack")
	}
}

// GoFunc 可以是闭包引用
type GoFunc func(context.Context, interface{}) interface{}

type gOption func(*gOptions)
type gOptions struct {
	wg       *sync.WaitGroup
	c        chan interface{} //结果输出的通道
	m        chan chan interface{} //merge channel和c互斥
	param    interface{} //GoFunc 参数
	panicMsg string //panic错误信息
}

func WithWg(wg *sync.WaitGroup) gOption {
	return func(g *gOptions) {
		g.wg = wg
	}
}

func WithChannel2(m chan  chan interface{}) gOption  {
	return func(g *gOptions) {
		g.m = m
	}
}

func WithChannel(c chan interface{}) gOption {
	return func(g *gOptions) {
		g.c = c
	}
}

func WithParam(param interface{}) gOption {
	return func(g *gOptions) {
		g.param = param
	}
}

func WithMsg(msg string) gOption {
	return func(g *gOptions) {
		g.panicMsg = msg
	}
}

func newGOptions(opts ...gOption) *gOptions {
	gCtx := &gOptions{
		wg:       nil,
		c:        nil,
		param:    nil,
		panicMsg: "",
	}
	for _, v := range opts {
		v(gCtx)
	}
	return gCtx
}

// NewG 创建goroutine
func NewG(f GoFunc, ctx context.Context,opts ...gOption) {
	gOpts := newGOptions(opts...)
	if gOpts.wg != nil {
		gOpts.wg.Add(1)
	}
	go func(p interface{}) { //创建goroutine时copy到stack上
		var err error
		defer recoverFunc(&err,gOpts.panicMsg)

		if gOpts.wg != nil {
			defer gOpts.wg.Done()
		}
		rt := f(ctx, p)
		if gOpts.m != nil {
			c := make(chan interface{},1)
			defer close(c)
			c <- rt
			gOpts.m <- c
		}

		if gOpts.c != nil {
			gOpts.c <- rt
		}

		return
	}(gOpts.param)
}