package xworkflow

import "context"

type Operator interface {
	Execute(ctx context.Context,input Inputs) interface{}
}

// BaseOperator 基础执行操作
type BaseOperator struct {

}

// GoOperator 执行的方法
type GoOperator struct {

}

// BashOperator bash执行
type BashOperator struct {

}

// MySqlOperator mysql执行器
type MySqlOperator struct {

}